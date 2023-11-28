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

type productRepo struct {
	pool dbconnection.Clients
}

func NewProductRepo(pool dbconnection.Clients) *productRepo {
	return &productRepo{
		pool: pool,
	}
}

func (p productRepo) Create(ctx context.Context, productIn ...entity.ProductItems) ([]entity.ProductItems, error) {
	var productData ProductData
	productDataList := productData.MapToRepoNewList(productIn)
	query := `INSERT INTO products (name, description, tag, price, quantity, created_at) VALUES 
                                    ($1, $2, $3, $4, $5, $6) RETURNING id`

	batch := &pgx.Batch{}

	for _, product := range productDataList {
		batch.Queue(query, product.Name, product.Description, product.Tag, product.Price, product.Quantity,
			product.CreatedAt)
	}
	br := p.pool.SendBatch(ctx, batch)
	defer func() {
		err := br.Close()
		if err != nil {
			slog.Error("productRepo.Create.br.Close, error while closing batch results", slog.Any("error", err))
		}
	}()
	var createdProduct []ProductData
	for _, product := range productDataList {
		var id int
		err := br.QueryRow().Scan(&id)
		if err != nil {
			return nil, err
		}
		product.ID = id
		createdProduct = append(createdProduct, product)
	}
	var resultProduct ProductData
	return resultProduct.MapToEntityList(createdProduct), nil
}

func (p productRepo) GetBy(ctx context.Context, productQuery entity.ProductQuery) (entity.ProductItems, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqb := psql.Select("id", "name", "description", "tag", "price", "quantity", "created_at").From("products")
	sqb = p.generateUpdateWhere(sqb, productQuery)

	query, args, err := sqb.ToSql()
	if err != nil {
		return entity.ProductItems{}, commonentity.NewDatabaseError(
			errors.Wrapf(err, "productRepo.GetBy.ToSql, error while getting products"),
			commonentity.ErrUnknownStatus,
			"error while getting products",
		)
	}
	row := p.pool.QueryRow(ctx, query, args...)
	var productData ProductData
	err = row.Scan(&productData.ID, &productData.Name, &productData.Description, &productData.Tag, &productData.Price,
		&productData.Quantity, &productData.CreatedAt)
	if err != nil {
		return entity.ProductItems{}, commonentity.NewDatabaseError(
			errors.Wrapf(err, "productRepo.GetBy.Scan, error while getting products"),
			commonentity.ErrUnknownStatus,
			"error while getting products",
		)
	}
	return productData.MapToEntity(), nil
}

func (p productRepo) GetsBy(ctx context.Context, productQuery entity.ProductQuery) ([]entity.ProductItems, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqb := psql.Select("id", "name", "description", "tag", "price", "quantity", "created_at").From("products")
	sqb = p.generateUpdateWhere(sqb, productQuery)

	query, args, err := sqb.ToSql()
	if err != nil {
		return nil, commonentity.NewDatabaseError(
			errors.Wrapf(err, "productRepo.GetsBy.ToSql, error while getting products"),
			commonentity.ErrUnknownStatus,
			"error while getting products",
		)
	}
	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, commonentity.NewDatabaseError(
			errors.Wrapf(err, "productRepo.GetsBy.Query, error while getting products"),
			commonentity.ErrUnknownStatus,
			"error while getting products",
		)
	}
	defer rows.Close()

	var products []ProductData
	for rows.Next() {
		var product ProductData
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Tag, &product.Price,
			&product.Quantity, &product.CreatedAt)
		if err != nil {
			return nil, commonentity.NewDatabaseError(
				errors.Wrapf(err, "productRepo.GetsBy.Scan, error while getting products"),
				commonentity.ErrUnknownStatus,
				"error while getting products",
			)
		}
		products = append(products, product)
	}
	var resultProduct ProductData
	return resultProduct.MapToEntityList(products), nil
}

func (p productRepo) Update(ctx context.Context, productIn ...entity.ProductItems) error {
	var productData ProductData
	productDataList := productData.MapToRepoList(productIn)

	_, err := p.pool.Exec(ctx, "CREATE TEMP TABLE temp_products (LIKE products INCLUDING DEFAULTS)")
	if err != nil {
		return commonentity.NewDatabaseError(
			errors.Wrapf(err, "productRepo.Update.Exec, error while creating temp table"),
			commonentity.ErrUnknownStatus,
			"error while creating temp table",
		)
	}
	defer func() {
		_, err = p.pool.Exec(ctx, "DROP TABLE temp_products")
		if err != nil {
			slog.Error("productRepo.Update.Exec, error while dropping temp table", slog.Any("error", err))
		}
	}()

	for _, product := range productDataList {
		_, err = p.pool.Exec(ctx, "INSERT INTO temp_products (name, description, tag, price, quantity, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
			product.Name, product.Description, product.Tag, product.Price, product.Quantity, product.CreatedAt)
		if err != nil {
			return commonentity.NewDatabaseError(
				errors.Wrapf(err, "productRepo.Update.Exec, error while inserting into temp table"),
				commonentity.ErrUnknownStatus,
				"error while updating products",
			)
		}
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqb := psql.Update("products")

	for _, product := range productDataList {
		sqb = p.generateUpdateSet(sqb, product)
	}
	sqb = sqb.Suffix("FROM temp_products WHERE products.id = temp_products.id")

	query, args, err := sqb.ToSql()
	if err != nil {
		return commonentity.NewDatabaseError(
			errors.Wrapf(err, "productRepo.Update.ToSql, error while updating products"),
			commonentity.ErrUnknownStatus,
			"error while updating products",
		)
	}
	_, err = p.pool.Exec(ctx, query, args...)
	if err != nil {
		return commonentity.NewDatabaseError(
			errors.Wrapf(err, "productRepo.Update.Query, error while updating products"),
			commonentity.ErrUnknownStatus,
			"error while updating products",
		)
	}
	return nil
}

func (p productRepo) Delete(ctx context.Context, productID int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := p.pool.Exec(ctx, query, productID)
	if err != nil {
		return commonentity.NewDatabaseError(
			errors.Wrapf(err, "productRepo.Delete.Exec, error while deleting product"),
			commonentity.ErrUnknownStatus,
			"error while deleting product",
		)
	}
	return nil
}

func (p productRepo) generateUpdateSet(sqb sq.UpdateBuilder, productDataList ProductData) sq.UpdateBuilder {
	if productDataList.Name.Valid {
		sqb = sqb.Set("name", productDataList.Name)
	}
	if productDataList.Description.Valid {
		sqb = sqb.Set("description", productDataList.Description)
	}
	if productDataList.Tag.Valid {
		sqb = sqb.Set("tag", productDataList.Tag)
	}
	if productDataList.Price.Valid {
		sqb = sqb.Set("price", productDataList.Price)
	}
	if productDataList.Quantity.Valid {
		sqb = sqb.Set("quantity", productDataList.Quantity)
	}
	return sqb
}

func (p productRepo) generateUpdateWhere(sqb sq.SelectBuilder, productQuery entity.ProductQuery) sq.SelectBuilder {
	if len(productQuery.IDs()) > 0 {
		sqb = sqb.Where(sq.Eq{"id": productQuery.IDs()})
	}
	if len(productQuery.Description()) > 0 {
		sqb = sqb.Where(sq.Eq{"description": productQuery.Description()})
	}
	if len(productQuery.Tag()) > 0 {
		sqb = sqb.Where(sq.Eq{"tag": productQuery.Tag()})
	}
	if len(productQuery.Quantity()) > 0 {
		sqb = sqb.Where(sq.Eq{"quantity": productQuery.Quantity()})
	}
	if len(productQuery.CreatedAt()) > 0 {
		sqb = sqb.Where(sq.Eq{"created_at": productQuery.CreatedAt()})
	}
	return sqb

}
