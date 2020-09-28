package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garcialuis/Gossip/api/models"

	nutriportclient "github.com/garcialuis/Nutriport/sdk/client"
	nutriportclient_models "github.com/garcialuis/Nutriport/sdk/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB              *gorm.DB
	Router          *mux.Router
	NutriportClient NutriportClientService
}

type NutriportClientService interface {
	BMIClient
	TEEClient
	FoodClient
}

type BMIClient interface {
	CalculateImperialBMI(weight, height float64) nutriportclient_models.Person
}

type TEEClient interface {
	CalculateTotalEnergyExpenditure(age int, gender int, weight float64, activityLevel string) nutriportclient_models.Person
}

type FoodClient interface {
	CreateFoodItem(foodItem nutriportclient_models.FoodItem) nutriportclient_models.FoodItem
	GetAllFoodItems() []nutriportclient_models.FoodItem
	DeleteFoodItem(foodItemName string) int
	GetFoodItemByName(foodItemName string) nutriportclient_models.FoodItem
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

	server.NutriportClient = nutriportclient.NewClient()

	server.InitializeRoutes()

}

func (server *Server) InitializeTestServer(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

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

	// TODO: Initialize NutriportClient to Mock
	server.NutriportClient = nutriportclient.NewClient()

	server.InitializeRoutes()

}

func (server *Server) Run(addr string) {
	fmt.Println("Listen to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
