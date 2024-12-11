package migrations

import (
	"log"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Migrate the schema
	log.Println("Running database migration...")

	// Pastikan ekstensi uuid-ossp diaktifkan jika belum ada
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatalf("Error activating uuid-ossp extension: %v", err)
		return err
	}

	if err := db.AutoMigrate(&entity.Merek{}, &entity.Mobil{}); err != nil {
		log.Fatalf("Error during migration: %v", err)
		return err
	}

	log.Println("Database migration completed successfully.")
	return nil
}

func Seeder(db *gorm.DB) error {
	log.Println("Running database seeder...")

	// Seed Merek data
	merekSeed := []entity.Merek{
		{ID: uuid.New(), Nama: "Toyota"},
		{ID: uuid.New(), Nama: "Honda"},
		{ID: uuid.New(), Nama: "BMW"},
	}

	for _, merek := range merekSeed {
		if err := db.FirstOrCreate(&merek, "nama = ?", merek.Nama).Error; err != nil {
			return err
		}
	}

	// Seed Mobil data
	mobilSeed := []entity.Mobil{
		{
			ID:              uuid.New(),
			MerekID:         merekSeed[0].ID,
			Type:            "Avanza",
			NoPlat:          "B 1234 ABC",
			Warna:           "Hitam",
			InitialCondition: "Baru",
			Harga:           200000000,
			Deskripsi:       "Mobil keluarga yang nyaman.",
			ImageUrl:        "https://example.com/avanza.jpg",
		},
		{
			ID:              uuid.New(),
			MerekID:         merekSeed[1].ID,
			Type:            "Civic",
			NoPlat:          "B 5678 XYZ",
			Warna:           "Putih",
			InitialCondition: "Baru",
			Harga:           350000000,
			Deskripsi:       "Sedan sport dengan performa tinggi.",
			ImageUrl:        "https://example.com/civic.jpg",
		},
		{
			ID:              uuid.New(),
			MerekID:         merekSeed[2].ID,
			Type:            "X5",
			NoPlat:          "B 8765 QWE",
			Warna:           "Silver",
			InitialCondition: "Bekas",
			Harga:           800000000,
			Deskripsi:       "SUV premium dengan fitur canggih.",
			ImageUrl:        "https://example.com/x5.jpg",
		},
	}

	for _, mobil := range mobilSeed {
		if err := db.FirstOrCreate(&mobil, "type = ?", mobil.Type).Error; err != nil {
			return err
		}
	}

	log.Println("Database seeder completed successfully.")
	return nil
}
