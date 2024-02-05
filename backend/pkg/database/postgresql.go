package database

import (
	"context"
	"fmt"
	"time"

	"github.com/a631807682/zerofield"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgreSQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type PostgreSQL struct {
	DB *gorm.DB
}

type PostgreSQLModel struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}

// Option - represents PostgreSQL service option.
type Option func(*PostgreSQL)

// SetLogger - configures logger
func SetLogger(log logger.Interface) Option {
	return func(p *PostgreSQL) {
		p.DB.Logger = log
	}
}

func NewPostgreSQL(cfg PostgreSQLConfig, opts ...Option) (*PostgreSQL, error) {
	sql := &PostgreSQL{}

	var err error
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port,
	)
	sql.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgresql: %w", err)
	}

	for _, opt := range opts {
		opt(sql)
	}

	// https://github.com/a631807682/zerofield
	// allow update zero value field
	err = sql.DB.Use(zerofield.NewPlugin())
	if err != nil {
		return nil, fmt.Errorf("failed to use zerofield plugin: %w", err)
	}

	return sql, nil
}

func (p *PostgreSQL) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	pgSql, err := p.DB.WithContext(ctx).DB()
	if err != nil {
		return fmt.Errorf("failed to connect")
	}
	err = pgSql.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect")
	}
	return nil
}

// Close - closes mysql service database connection.
func (p *PostgreSQL) Close() error {
	if p.DB != nil {
		db, _ := p.DB.DB()
		err := db.Close()
		if err != nil {
			return fmt.Errorf("failed to close postgresql connection: %w", err)
		}
	}
	return nil
}
