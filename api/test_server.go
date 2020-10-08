package api

import (
	"os"

	"github.com/garcialuis/Gossip/api/controllers"
)

var testServer = controllers.Server{}

func RunTestServer() {

	// TODO: Need to implement InitializeTestServer function to initialize server with Mocks to external services
	// Using non-test Initialize function for now.
	testServer.InitializeTestServer(os.Getenv("DB_POSTGRES_DRIVER"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_NAME"))

	testServer.Run(":8080")
}
