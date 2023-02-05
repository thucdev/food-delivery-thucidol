package restaurantmodel

import "testing"

type testData struct {
	Input  RestaurantCreate
	Expect error
}

func TestRestaurantCreate_Validate(t *testing.T) {
	dataTable := []testData{
		{Input: RestaurantCreate{Name: ""}, Expect: ErrNameIsEmpty},
		{Input: RestaurantCreate{Name: "thucidol"}, Expect: nil},
	}

	for _, item := range dataTable {
		err := item.Input.Validate()

		if err != item.Expect {
			// t.Error("Validate restaurant. Input: %v, Expect: %v, Output: %v", item.Input, item.Expect, err)
			t.Error("loi roi")
		}
	}

	t.Log("Validate passed")

}

// func TestRestaurantCreate_Validate(t *testing.T) {
// 	dataTest := RestaurantCreate{
// 		Name: "",
// 	}

// 	err := dataTest.Validate()

// 	if err == ErrNameIsEmpty {
// 		t.Error("Input name can not empty")
// 		return
// 	}

// 	t.Log("Validate passed")

// }
