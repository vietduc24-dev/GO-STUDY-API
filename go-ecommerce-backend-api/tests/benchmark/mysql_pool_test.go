package benchmark

import (
	"go-ecommerce-backend-api/internal/po"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dsn = "root:root1234@tcp(127.0.0.1:3306)/shopducteo?charset=utf8mb4&parseTime=True&loc=Local"

// initDB khởi tạo kết nối GORM, tắt log query và auto-migrate bảng User
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Đảm bảo bảng go_users đã tồn tại trước khi benchmark insert
	if err := db.AutoMigrate(&po.User{}); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	return db
}

// =========================================================================
// PHẦN 1: BENCHMARK INSERT RECORD (GHI DỮ LIỆU THỰC TẾ VÀO DATABASE)
// =========================================================================

// BenchmarkInsert_NoPool: Mỗi lần INSERT lại mở một connection mới và đóng lại
func BenchmarkInsert_NoPool(b *testing.B) {
	initDB() // Đảm bảo bảng đã tồn tại
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			b.Fatal(err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			b.Fatal(err)
		}

		user := po.User{
			UUID:     uuid.New(),
			UserName: "user_nopool",
			IsActive: true,
		}
		if err := db.Create(&user).Error; err != nil {
			b.Error(err)
		}

		sqlDB.Close()
	}
}

// BenchmarkInsert_WithPool: Dùng Connection Pool để INSERT liên tục
func BenchmarkInsert_WithPool(b *testing.B) {
	db := initDB()
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := po.User{
			UUID:     uuid.New(),
			UserName: "user_pool",
			IsActive: true,
		}
		if err := db.Create(&user).Error; err != nil {
			b.Error(err)
		}
	}
}

// BenchmarkParallel_Insert_Pool1: INSERT đồng thời với Pool chỉ có 1 connection
func BenchmarkParallel_Insert_Pool1(b *testing.B) {
	db := initDB()
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			user := po.User{
				UUID:     uuid.New(),
				UserName: "user_parallel",
				IsActive: true,
			}
			if err := db.Create(&user).Error; err != nil {
				b.Error(err)
			}
		}
	})
}

// BenchmarkParallel_Insert_Pool10: INSERT đồng thời với Pool tối đa 10 connections
func BenchmarkParallel_Insert_Pool10(b *testing.B) {
	db := initDB()
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			user := po.User{
				UUID:     uuid.New(),
				UserName: "user_parallel",
				IsActive: true,
			}
			if err := db.Create(&user).Error; err != nil {
				b.Error(err)
			}
		}
	})
}

// BenchmarkParallel_Insert_Pool50: INSERT đồng thời với Pool tối đa 50 connections
func BenchmarkParallel_Insert_Pool50(b *testing.B) {
	db := initDB()
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			user := po.User{
				UUID:     uuid.New(),
				UserName: "user_parallel",
				IsActive: true,
			}
			if err := db.Create(&user).Error; err != nil {
				b.Error(err)
			}
		}
	})
}

// BenchmarkParallel_Insert_Pool100: INSERT đồng thời với Pool tối đa 100 connections
func BenchmarkParallel_Insert_Pool100(b *testing.B) {
	db := initDB()
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			user := po.User{
				UUID:     uuid.New(),
				UserName: "user_parallel",
				IsActive: true,
			}
			if err := db.Create(&user).Error; err != nil {
				b.Error(err)
			}
		}
	})
}

// =========================================================================
// PHẦN 2: BENCHMARK QUERY NHANH (SELECT 1 - ĐO OVERHEAD CỦA CONNECTION POOL)
// =========================================================================

// BenchmarkSelect_NoPool: Mỗi lần SELECT 1 lại mở connection mới
func BenchmarkSelect_NoPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			b.Fatal(err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			b.Fatal(err)
		}

		var result int
		if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
			b.Error(err)
		}

		sqlDB.Close()
	}
}

// BenchmarkSelect_WithPool: Tái sử dụng connection trong Pool để SELECT 1
func BenchmarkSelect_WithPool(b *testing.B) {
	db := initDB()
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result int
		if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
			b.Error(err)
		}
	}
}
