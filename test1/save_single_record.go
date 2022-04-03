package main

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Device struct {
	SerialNumber string `gorm:"primaryKey;not null"`
	Settings     int	`gorm:"not null"`
	Location     string `gorm:"size:20;not null"`
}

func SaveDevice(db *gorm.DB, device Device) error {
	persisted := &Device{}
	err := db.Where(&Device{SerialNumber: device.SerialNumber}).First(persisted).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.Create(&device).Error		// Creates a new record
	} else {
		err = db.Model(&device).Update("settings", device.Settings).Error
	}
	return err
}

func main() {
	// see: https://gorm.io/docs/connecting_to_the_database.html
	// and for loggers: https://gorm.io/docs/logger.html
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Panic(err.Error())
		return
	}
	db.AutoMigrate(Device{})


	SaveDevice(db, Device{SerialNumber: "SN0001", Settings: 2, Location: "94566"})
	SaveDevice(db, Device{SerialNumber: "SN0001", Settings: 1, Location: "94566"})
}
