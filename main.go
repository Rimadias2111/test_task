package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"os"
	"test_task/logic"
	"test_task/logic/cashier"
	"test_task/logic/manager"
	"test_task/service"
	"test_task/storage"
	"time"
)

func setupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Connecting to DB with DSN: %s", dsn)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	log.Println("Database connection established")
	if err := storage.Migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: error loading .env file:", err)
	}
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	store := storage.New(db)
	serv := service.New(store, logger)
	logicInstance := logic.NewLogic(serv)
	cashierInstance := cashier.NewCashier(logicInstance)
	managerInstance := manager.NewManager(logicInstance)
	Run(cashierInstance, managerInstance)
}

func Run(c *cashier.Cashier, m *manager.Manager) {
	for {
		fmt.Println("\nSelect role:")
		fmt.Println("1 - Manager")
		fmt.Println("2 - Cashier")
		fmt.Println("0 - Exit")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			m.Menu()
		case 2:
			c.Menu()
		case 0:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
