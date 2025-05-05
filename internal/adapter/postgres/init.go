package postgres

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"github.com/victorsvart/go-ecommerce/pkg/rbac"
	"github.com/victorsvart/go-ecommerce/pkg/utils"
	"gorm.io/gorm"
)

func InitData(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
	)

	SeedAdmin(db)
	SeedSampleUser(db)
	SeedProduct(db)
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
		Contact:  "(21) 98873-3943",
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
		Contact:  "(21) 98873-9942",
		RoleID:   rbac.UserRoleID,
		Password: password,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Error creating admin user: %v", err)
		return
	}

	log.Println("Sample user seeded successfully")
}

func SeedProduct(db *gorm.DB) {
	var count int64
	err := db.Model(&domain.Product{}).Count(&count).Error
	if err != nil {
		log.Fatalf("Error querying db during product: %v", err)
	}

	if count > 0 {
		log.Println("Products already seeded")
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	filePath := filepath.Join(dir, "internal", "adapter", "postgres", "json", "products.json")
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading products.json seed file: %v", err)
	}

	var products []domain.Product
	if err := json.Unmarshal(file, &products); err != nil {
		log.Fatalf("Error unmarshalling products.json: %v", err)
	}

	if len(products) == 0 {
		log.Println("No products found in seed file.")
		return
	}

	if err := db.Create(&products).Error; err != nil {
		log.Fatalf("Error seeding products: %v", err)
	}

	log.Printf("Seeded %d products successfully.", len(products))

}
