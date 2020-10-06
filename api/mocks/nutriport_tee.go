package mocks

import nutriportclient_models "github.com/garcialuis/Nutriport/sdk/models"

type TEEClient struct{}

func NewTEEClientMock() *TEEClient {
	return &TEEClient{}
}

func (c *TEEClient) CalculateTotalEnergyExpenditure(age int, gender int, weight float64, activityLevel string) nutriportclient_models.Person {

	person := nutriportclient_models.Person{
		TEE:           2006.54,
		ActivityLevel: activityLevel,
		Age:           uint(age),
		Gender:        uint(gender),
		Weight:        weight,
	}

	return person
}
