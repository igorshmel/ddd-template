package postgres

import (
	"default_ddd/app/internal/adapters/repository/models"
	"default_ddd/app/pkg/config"
	"default_ddd/app/pkg/errs"
	log "default_ddd/app/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MigrateEnum - интерфейс для перечислений (enum)
type MigrateEnum interface {
	MigrationSQL() string
}

// SQLStore fulfills the Extractor and Persister document interfaces
type SQLStore struct {
	db  *gorm.DB
	log log.Logger
}

// NewPostgresRepository returns a repository instance
func NewPostgresRepository(cfg config.Config, log log.Logger, migrate bool) (*SQLStore, error) {
	dbGorm, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.CreateDSN(),
	}))
	if err != nil {
		return nil, err
	}

	postgresDb, err := dbGorm.DB()
	if err != nil {
		return nil, err
	}

	postgresDb.SetMaxIdleConns(10)
	postgresDb.SetMaxOpenConns(100)

	if migrate {
		if err = migrateData(dbGorm); err != nil {
			return nil, err
		}
	}

	return &SQLStore{
		db:  dbGorm,
		log: log,
	}, nil
}

// migrateData --
func migrateData(db *gorm.DB) error {

	// Создать в БД перечисляемый тип для статуса сущности CartItem
	var cartItemModel models.CartItemModel
	db.Exec(cartItemModel.CreateCartItemStatusEnum())

	// Создать в БД перечисляемый тип для статуса сущности Order
	var orderModel models.OrderModel
	db.Exec(orderModel.CreateOrderStatusEnum())

	// создаёт основные таблицы в базе данных
	return db.AutoMigrate(
		&models.UserModel{},
		&models.ProductModel{},
		&models.OrderModel{},
		&models.CartItemModel{},
	)
}

// verifyPointer --
func (ths *SQLStore) verifyPointer(dbo interface{}) error {
	if ths == nil || ths.db == nil {
		return errs.ErrEmptyPointer
	}
	if dbo == nil {
		return errs.ErrEmptyData
	}
	return nil
}
