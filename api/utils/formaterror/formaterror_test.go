package formaterror

import "testing"

func TestFormatError(t *testing.T) {
	inputString := "nickname error"
	expectedOutput := "Nickname is already taken"

	err := FormatError(inputString)

	if expectedOutput != err.Error() {
		t.Errorf("Expected to be %s, but got %s", expectedOutput, err.Error())
	}
}
