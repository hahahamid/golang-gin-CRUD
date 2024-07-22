package main

import (
	models "github.com/hahahamid/go-crud/Models"
	"github.com/hahahamid/go-crud/initializers"
)

func init() {
	initializers.LoadEnvironmentVaraible()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.BlogPost{})
}
