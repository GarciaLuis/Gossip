package mocks

import nutriportclient_models "github.com/garcialuis/Nutriport/sdk/models"

type BMIClient struct {
}

func NewBMIClientMock() *BMIClient {
	return &BMIClient{}
}

func (c *BMIClient) CalculateImperialBMI(weight, height float64) nutriportclient_models.Person {

	person := nutriportclient_models.Person{
		BMI:    25.06,
		Weight: weight,
		Height: height,
	}

	return person
}
