package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/b1g-nguyx/ADRDMS/shared/go/pkg/config"
	"github.com/b1g-nguyx/ADRDMS/shared/go/pkg/health"
	"github.com/b1g-nguyx/ADRDMS/shared/go/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	KafkaBroker  string `mapstructure:"KAFKA_BROKER"`
	KafkaTopic   string `mapstructure:"KAFKA_TOPIC"`
	PostgresDSN  string `mapstructure:"POSTGRES_DSN"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`
}

func main() {
	logger.InitLogger(true)
	log := logger.Get()
	defer logger.Sync()

	cfg := Config{
		Port:         "8080",
		KafkaBroker:  "localhost:9092",
		KafkaTopic:   "test-topic",
		PostgresDSN:  "host=localhost user=postgres password=postgres dbname=postgres port=5005 sslmode=disable",
		MigrationURL: "postgres://postgres:postgres@localhost:5005/postgres?sslmode=disable",
	}

	if err := config.LoadConfig(".", "config", "yaml", &cfg); err != nil {
		log.Warn("Failed to load config, using defaults", zap.Error(err))
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.KafkaBroker == "" {
		cfg.KafkaBroker = "localhost:9092"
	}
	if cfg.KafkaTopic == "" {
		cfg.KafkaTopic = "test-topic"
	}
	if cfg.PostgresDSN == "" {
		cfg.PostgresDSN = "host=localhost user=postgres password=postgres dbname=postgres port=5005 sslmode=disable"
	}
	if cfg.MigrationURL == "" {
		cfg.MigrationURL = "postgres://postgres:postgres@localhost:5005/postgres?sslmode=disable"
	}

	log.Info(fmt.Sprintf("🤖 Robot Service is starting on port %s...", cfg.Port))

	// Kafka connection
	kafkaWriter := &kafka.Writer{
		Addr:  kafka.TCP(cfg.KafkaBroker),
		Topic: cfg.KafkaTopic,
	}
	_ = kafkaWriter // mock usage
	log.Info("Kafka connection initialized (mock/stub)")

	// Postgres connection via Gorm
	_, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Warn("Failed to connect to Postgres (mock), ignoring...", zap.Error(err))
	} else {
		log.Info("Postgres connection initialized (mock/stub)")
	}

	// Run migrations
	m, err := migrate.New("file://database/migrations", cfg.MigrationURL)
	if err == nil {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Error("Migration failed", zap.Error(err))
		} else {
			log.Info("Database migrations applied successfully")
		}
	} else {
		log.Warn("Failed to setup migration (mock), ignoring...", zap.Error(err))
	}

	// HTTP Server
	http.HandleFunc("/health", health.HTTPHealthHandler)
	addr := ":" + strings.TrimPrefix(cfg.Port, ":")
	log.Info("Starting HTTP server on " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Error("Failed to start server", zap.Error(err))
	}
}
