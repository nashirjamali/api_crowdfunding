package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Username     string `yaml:"user"`
	Password     string `yaml:"pass"`
	Port         int    `yaml:"port"`
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"db"`
}

type Configuration struct {
	Database Database `yaml:"database"`
}

var Config Configuration
var DB *gorm.DB

func Init() {
	confFilePath := "config/config.yml"
	f, err := os.Open(confFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}

	initDB()
}

func initDB() {
	if Config.Database.Username == "" && Config.Database.Password == "" {
		panic("Database username and password are empty")
	}
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.DatabaseName,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// tables := []interface{}{
	// 	user.User{},
	// 	campaign.Campaign{},
	// 	// campaign.CampaignImage{},
	// }

	// err = DB.AutoMigrate(tables...)

	// if err != nil {
	// 	panic(err.Error())
	// }

}
