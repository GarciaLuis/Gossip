package mocks

import nutriportclient_models "github.com/garcialuis/Nutriport/sdk/models"

type NutriportClientMock struct {
	BMIClientMock
	TEEClientMock
	FoodClientMock
}

type BMIClientMock interface {
	CalculateImperialBMI(weight, height float64) nutriportclient_models.Person
}

type TEEClientMock interface {
	CalculateTotalEnergyExpenditure(age int, gender int, weight float64, activityLevel string) nutriportclient_models.Person
}

type FoodClientMock interface {
	CreateFoodItem(foodItem nutriportclient_models.FoodItem) nutriportclient_models.FoodItem
	GetAllFoodItems() []nutriportclient_models.FoodItem
	DeleteFoodItem(foodItemName string) int
	GetFoodItemByName(foodItemName string) nutriportclient_models.FoodItem
}

func NewNutriportClientMock() NutriportClientMock {

	// cli := new(NutriportClientMock)
	cli := NutriportClientMock{}
	cli.BMIClientMock = NewBMIClientMock()
	cli.TEEClientMock = NewTEEClientMock()
	cli.FoodClientMock = NewFoodClientMock()

	return cli
}

func CalculateImperialBMI(weight, height float64) nutriportclient_models.Person {

	person := nutriportclient_models.Person{
		BMI:    24.5,
		Weight: weight,
		Height: height,
	}

	return person
}
