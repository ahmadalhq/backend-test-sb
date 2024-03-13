package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB struct {
	DbPostgres *gorm.DB
}

var (
	onceDbPostgres sync.Once
	instanceDB     *DB
)

func GetInstancePostgresDb() *gorm.DB {
	onceDbPostgres.Do(func() {
		postgreInfo := FileConfig.Database.Postgre
		lf := logrus.WithField("host", postgreInfo.Host)

		dbConfig := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s fallback_application_name=diyo_pos_app",
			postgreInfo.Host, postgreInfo.Port, postgreInfo.Username, postgreInfo.Password, postgreInfo.Name, postgreInfo.Schema, postgreInfo.SSLMode,
		)

		sqlDB, err := sql.Open("postgres", dbConfig)
		if err != nil {
			lf.WithError(err).Fatal("failed connect to postgres db")
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetConnMaxLifetime(10 * time.Minute)

		dialect := postgres.New(postgres.Config{Conn: sqlDB})
		loggerLevel := logger.Error
		if postgreInfo.LogMode {
			loggerLevel = logger.Info
		}

		dbConnection, err := gorm.Open(dialect, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      loggerLevel,
				},
			),
		})
		if err != nil {
			lf.WithError(err).Fatal("failed open gorm session")
		}

		lf.WithField("log_mode", postgreInfo.LogMode).Info("connected to postgres db")
		instanceDB = &DB{DbPostgres: dbConnection}
	})
	return instanceDB.DbPostgres
}
