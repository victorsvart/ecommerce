package database

import (
	"log"
	"os"

	"github.com/victorsvart/go-ecommerce/internal/user/domain"
	"github.com/victorsvart/go-ecommerce/pkg/rbac"
	"github.com/victorsvart/go-ecommerce/pkg/utils"
	"gorm.io/gorm"
)

func InitData(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{},
	)

	SeedAdmin(db)
	SeedSampleUser(db)
}

func SeedAdmin(db *gorm.DB) {
	email := "admin@godmode.com"
	password := os.Getenv("LOCAL_ADMIN_PASSWORD")

	if password == "" {
		log.Fatal("LOCAL_ADMIN_PASSWORD environment variable is not set")
		return
	}

	var user domain.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		log.Println("Admin user already exists, skipping seeding")
		return
	}

	utils.HashPassword(&password)
	admin := domain.User{
		Name:     "admin",
		Surname:  "admin",
		Email:    email,
		RoleID:   rbac.AdminRoleID,
		Password: password,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Error creating admin user: %v", err)
		return
	}

	log.Println("Admin user seeded successfully")
}

func SeedSampleUser(db *gorm.DB) {
	email := "victor@sample.com"
	password := os.Getenv("LOCAL_USER_PASSWORD")

	if password == "" {
		log.Fatal("LOCAL_USER_PASSWORD environment variable is not set")
		return
	}

	var user domain.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		log.Println("Sample user already exists, skipping seeding")
		return
	}

	utils.HashPassword(&password)
	admin := domain.User{
		Name:     "Victor",
		Surname:  "Moraes",
		Email:    email,
		RoleID:   rbac.UserRoleID,
		Password: password,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Error creating admin user: %v", err)
		return
	}

	log.Println("Sample user seeded successfully")
}
