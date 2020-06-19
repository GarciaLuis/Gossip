package api

import (
	"os"

	"github.com/garcialuis/Gossip/api/controllers"
)

var server = controllers.Server{}

func Run() {

	server.Initialize(os.Getenv("DB_POSTGRES_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Uncomment to seed the database with data from seeder
	//seed.Load(server.DB)

	server.Run(":8080")
}
