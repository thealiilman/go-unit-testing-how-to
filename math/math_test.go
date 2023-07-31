package math

import "testing"

type AddData struct {
	x, y 	int
	result 	int
}

func TestAdd(t *testing.T) {
	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Expected %d from Add(1, 3). Got %d\n", 4, result)
	}

	testData := []AddData {
		{ 1, 2, 3 },
		{ 3, 5, 8 },
		{ 3, -3, 0 },
	}

	for _, datum := range testData {
		result := Add(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("Expected %v from Add(%v, %v). Got %v\n", datum.result, datum.x, datum.y, result)
		}
	}
}

func TestDivide(t *testing.T) {
	result := Divide(5, 0)

	if result != 0 {
		t.Errorf("Expected %f from Divide(5, 0). Got %f\n", 0.0, result)
	}
}

