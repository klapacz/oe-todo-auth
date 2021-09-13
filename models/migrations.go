package models

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrations = []*gormigrate.Migration{
	// install uui-ossp extension
	{
		ID: "202109130454",
		Migrate: func(tx *gorm.DB) error {
			tx.Exec("CREATE EXTENSION \"uuid-ossp\";")
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			tx.Exec("DROP EXTENSION \"uuid-ossp\";")
			return nil
		},
	},
	// create users table
	{
		ID: "202109130509",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				gorm.Model
				ID       string `gorm:"type:uuid;default:uuid_generate_v4()"`
				Password string
				Email    string
			}
			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	},
}

func runMigrations() {
	m := gormigrate.New(DB, gormigrate.DefaultOptions, migrations)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
