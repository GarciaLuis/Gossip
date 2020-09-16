package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garcialuis/Gossip/api/models"
	"github.com/garcialuis/Nutriport/sdk/client/bmi"
	bmi_models "github.com/garcialuis/Nutriport/sdk/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
	// BMIClient *bmi.BMIClientService
	BMIClient BMIClientService
}

type BMIClientService interface {
	// TODO: Interface should require BMI Client functions:
	CalculateImperialBMI(weight, height float64) bmi_models.Person //models.Person
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	//database migration
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	server.Router = mux.NewRouter()

	server.BMIClient = bmi.NewBMIService()

	server.InitializeRoutes()

}

func (server *Server) Run(addr string) {
	fmt.Println("Listen to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
