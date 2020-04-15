package question

import (
	"testing"
)

func TestTheLengthOfTheLastWordInTheString(t *testing.T) {
	params := "hello world\n"
	value := TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "   hello world\n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "hello   world\n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "hello world   \n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "   hello   world   \n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "hello-world\n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 11 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 11, value)
	}

	params = "\n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 0 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 0, value)
	}
}

func TestCountTheNumberOfCharacters(t *testing.T) {
	param1 := "ABCDEF"
	param2 := "A"
	your := CountTheNumberOfCharacters(param1, param2)
	correct := 1
	if your != correct {
		t.Fatalf("param1: %s, param2: %s, correct: %d, your: %d", param1, param2, correct, your)
	}

	param1 = "ABCDEFAABDD"
	param2 = "A"
	your = CountTheNumberOfCharacters(param1, param2)
	correct = 3
	if your != correct {
		t.Fatalf("param1: %s, param2: %s, correct: %d, your: %d", param1, param2, correct, your)
	}

}

func TestObviouslyRandomNumbers(t *testing.T) {
	param := []int{3, 3, 2, 4}
	result := ObviouslyRandomNumbers(param)
	t.Logf("%v", result)
}
