package repository

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/commonentity"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/dbconnection"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type userRepoImpl struct {
	pool dbconnection.Clients
}

func NewUserRepo(pool dbconnection.Clients) *userRepoImpl {
	return &userRepoImpl{
		pool: pool,
	}
}

func (u *userRepoImpl) GetBy(ctx context.Context, userQuery entity.UserQuery) (entity.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqb := psql.Select("*").From("users")
	sqb = u.generateUpdateWhere(sqb, userQuery)
	query, args, err := sqb.ToSql()
	if err != nil {
		return entity.User{}, err
	}
	var user UserData
	err = u.pool.QueryRow(ctx, query, args...).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName,
		&user.Email, &user.Age, &user.IsMarried, &user.Password, &user.Status, &user.CreatedAt)
	if err != nil {
		return entity.User{}, err
	}
	return user.MapToEntity()
}

func (u *userRepoImpl) GetsBy(ctx context.Context, userQuery entity.UserQuery) ([]entity.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqb := psql.Select("id", "user_name", "first_name", "last_name", "email", "age", "is_married", "password",
		"status", "created_at").From("users")
	sqb = u.generateUpdateWhere(sqb, userQuery)
	query, args, err := sqb.ToSql()
	if err != nil {
		return nil, commonentity.NewDatabaseError(
			errors.Wrapf(err, "userRepoImpl.GetsBy.ToSql, error while getting users"),
			commonentity.ErrUnknownStatus,
			"error while getting users",
		)
	}
	rows, err := u.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, commonentity.NewDatabaseError(
			errors.Wrapf(err, "userRepoImpl.GetsBy.Query, error while getting users"),
			commonentity.ErrUnknownStatus,
			"error while getting users",
		)
	}
	defer rows.Close()

	var users []UserData
	for rows.Next() {
		var user UserData
		err = rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Age,
			&user.IsMarried, &user.Password, &user.Status, &user.CreatedAt)
		if err != nil {
			return nil, commonentity.NewDatabaseError(
				errors.Wrapf(err, "userRepoImpl.GetsBy.Scan, error while getting users"),
				commonentity.ErrUnknownStatus,
				"error while getting users",
			)
		}
		users = append(users, user)
	}
	var resultUser UserData
	return resultUser.MapToEntityList(users)
}

func (u *userRepoImpl) Create(ctx context.Context, usersIn ...entity.User) ([]entity.User, error) {
	var usersData UserData
	usersDataList := usersData.MapToRepoNewList(usersIn)
	query := `INSERT INTO users (user_name, first_name, last_name, email, age, is_married, password, status, 
                   created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	batch := &pgx.Batch{}

	for _, user := range usersDataList {
		batch.Queue(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Age, user.IsMarried,
			user.Password, user.Status, user.CreatedAt)

	}
	br := u.pool.SendBatch(ctx, batch)
	defer func() {
		err := br.Close()
		if err != nil {
			slog.Error("userRepoImpl.Create.br.Close, error while closing batch results", slog.Any("error", err))
		}
	}()
	var createdUsers []UserData
	for _, user := range usersDataList {
		var id int
		err := br.QueryRow().Scan(&id)
		if err != nil {
			return nil, err
		}
		user.ID = id
		createdUsers = append(createdUsers, user)
	}
	var resultUser UserData
	return resultUser.MapToEntityList(createdUsers)
}

func (u *userRepoImpl) Update(ctx context.Context, usersIn ...entity.User) error {
	var usersData UserData
	usersDataList := usersData.MapToRepoList(usersIn)
	_, err := u.pool.Exec(ctx, `CREATE TEMP TABLE users_tmp ON COMMIT DROP AS SELECT * FROM users WITH NO DATA`)
	if err != nil {
		return commonentity.NewDatabaseError(
			errors.Wrapf(err, "userRepoImpl.Update.Exec, error while updating users"),
			commonentity.ErrUnknownStatus,
			"error while updating users",
		)
	}
	defer func() {
		_, err := u.pool.Exec(ctx, `DROP TABLE users_tmp`)
		if err != nil {
			slog.Error("userRepoImpl.Update.Exec, error while dropping table users_tmp", slog.Any("error", err))
		}
	}()

	for _, user := range usersDataList {
		_, err := u.pool.Exec(ctx, `INSERT INTO users_tmp (id, user_name, first_name, last_name, email, age, 
			is_married, password, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			user.ID, user.UserName, user.FirstName, user.LastName, user.Email, user.Age, user.IsMarried,
			user.Password, user.Status, user.CreatedAt)
		if err != nil {
			return commonentity.NewDatabaseError(
				errors.Wrapf(err, "userRepoImpl.Update.Exec, error while inserting into users_tmp"),
				commonentity.ErrUnknownStatus,
				"error while updating users",
			)
		}
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqb := psql.Update("users")

	for _, user := range usersDataList {
		sqb = u.generateUpdateSet(sqb, user)
	}
	sqb = sqb.Suffix("FROM users_tmp WHERE users.id = users_tmp.id")

	query, args, err := sqb.ToSql()
	if err != nil {
		return commonentity.NewDatabaseError(
			errors.Wrapf(err, "userRepoImpl.Update.ToSql, error while generating a query"),
			commonentity.ErrUnknownStatus,
			"error while updating users",
		)
	}
	_, err = u.pool.Exec(ctx, query, args...)
	if err != nil {
		return commonentity.NewDatabaseError(
			errors.Wrapf(err, "userRepoImpl.Update.Query, error while updating users"),
			commonentity.ErrUnknownStatus,
			"error while updating users",
		)
	}
	return nil
}

func (u *userRepoImpl) Delete(ctx context.Context, userID int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := u.pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepoImpl) generateUpdateSet(sqb sq.UpdateBuilder, user UserData) sq.UpdateBuilder {
	if user.UserName.Valid {
		sqb = sqb.Set("user_name", user.UserName)
	}
	if user.FirstName.Valid {
		sqb = sqb.Set("first_name", user.FirstName)
	}
	if user.LastName.Valid {
		sqb = sqb.Set("last_name", user.LastName)
	}
	if user.Email.Valid {
		sqb = sqb.Set("email", user.Email)
	}
	if user.Age.Valid {
		sqb = sqb.Set("age", user.Age)
	}
	if user.IsMarried.Valid {
		sqb = sqb.Set("is_married", user.IsMarried)
	}
	if user.Password.Valid {
		sqb = sqb.Set("password", user.Password)
	}
	if user.Status.Valid {
		sqb = sqb.Set("status", user.Status)
	}
	return sqb
}

func (u *userRepoImpl) generateUpdateWhere(sqb sq.SelectBuilder, userQuery entity.UserQuery) sq.SelectBuilder {
	if len(userQuery.IDs()) > 0 {
		sqb = sqb.Where(sq.Eq{"id": userQuery.IDs})
	}
	if len(userQuery.UserName()) > 0 {
		sqb = sqb.Where(sq.Eq{"user_name": userQuery.UserName})
	}
	if len(userQuery.FirstName()) > 0 {
		sqb = sqb.Where(sq.Eq{"first_name": userQuery.FirstName})
	}
	if len(userQuery.LastName()) > 0 {
		sqb = sqb.Where(sq.Eq{"last_name": userQuery.LastName})
	}
	if len(userQuery.Email()) > 0 {
		sqb = sqb.Where(sq.Eq{"email": userQuery.Email})
	}
	if len(userQuery.Age()) > 0 {
		sqb = sqb.Where(sq.Eq{"age": userQuery.Age})
	}
	if len(userQuery.IsMarried()) > 0 {
		sqb = sqb.Where(sq.Eq{"is_married": userQuery.IsMarried})
	}
	if len(userQuery.Status()) > 0 {
		sqb = sqb.Where(sq.Eq{"status": userQuery.Status})
	}
	if len(userQuery.CreatedAt()) > 0 {
		sqb = sqb.Where(sq.Eq{"created_at": userQuery.CreatedAt})
	}
	return sqb
}
