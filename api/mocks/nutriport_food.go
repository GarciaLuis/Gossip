package mocks

import nutriportclient_models "github.com/garcialuis/Nutriport/sdk/models"

type FoodClient struct{}

func NewFoodClientMock() *FoodClient {
	return &FoodClient{}
}

func (c *FoodClient) CreateFoodItem(foodItem nutriportclient_models.FoodItem) nutriportclient_models.FoodItem {

	newFoodItem := nutriportclient_models.FoodItem{}

	return newFoodItem
}

func (c *FoodClient) GetAllFoodItems() []nutriportclient_models.FoodItem {

	foodItemList := []nutriportclient_models.FoodItem{
		nutriportclient_models.FoodItem{},
		nutriportclient_models.FoodItem{},
		nutriportclient_models.FoodItem{},
	}

	return foodItemList
}

func (c *FoodClient) DeleteFoodItem(foodItemName string) int {
	return 1
}

func (c *FoodClient) GetFoodItemByName(foodItemName string) nutriportclient_models.FoodItem {

	foodItem := nutriportclient_models.FoodItem{}

	return foodItem
}
