package dbconnection

import (
	"context"
	"io/fs"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
	pkgErrors "github.com/pkg/errors"

	"github.com/yerkinalagozov/clean-code-showcase.git/configs"
)

type Clients interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

func NewConnection(ctx context.Context, cfg configs.Config) (*pgxpool.Pool, error) {
	var err error

	poolConfig, err := pgxpool.ParseConfig(cfg.DBUrl())
	if err != nil {
		return nil, pkgErrors.Wrapf(err, "dbconnection.NewPostgres.pgxpool.ParseConfig, failed to parse postgres url")
	}
	poolConfig.MaxConns = cfg.DBMaxOpenConn
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, pkgErrors.Wrapf(err, "dbconnection.NewPostgres.pgxpool.Connect, failed to connect to postgres, "+
			"postgresURL: %s", cfg.DBUrl())
	}

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, pkgErrors.Wrapf(err, "dbconnection.NewPostgres.pool.Ping, failed to ping postgres")
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		pool.Close()
		return nil, pkgErrors.Wrapf(err, "dbconnection.NewPostgres-db.Acquire, error while acquiring connection")
	}

	err = migrateDatabase(ctx, conn.Conn(), cfg.DBMigrationPath(), cfg.DBMigrationScheme)
	if err != nil {
		conn.Release()
		pool.Close()
		return nil, pkgErrors.Wrapf(err, "dbconnection.NewPostgres.migrateDatabase, error while migrating database")
	}

	conn.Release()
	return pool, nil
}

func migrateDatabase(ctx context.Context, conn *pgx.Conn, path fs.FS, schemeTable string) error {

	migrator, err := migrate.NewMigrator(ctx, conn, schemeTable)
	if err != nil {
		return pkgErrors.Wrapf(err, "dbconnection.migrateDatabase.migrate.NewMigrator, failed to create migrator")
	}

	err = migrator.LoadMigrations(path)
	if err != nil {
		return pkgErrors.Wrapf(err, "dbconnection.migrateDatabase.migrator.LoadMigrations, "+
			"error while loading migrations")
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		return pkgErrors.Wrapf(err, "dbconnection.migrateDatabase.migrator.Migrate, error while migrating")
	}

	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		return pkgErrors.Wrap(err, "dbconnection.migrateDatabase.migrator.GetCurrentVersion, "+
			"error while getting current version")
	}
	slog.Debug("End migrateDatabase. ", slog.Any("version", ver))
	return nil
}
