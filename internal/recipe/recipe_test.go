package recipe

import (
	"fmt"
	"testing"
)

func TestRecipeNew(t *testing.T) {
	expectedId := uint64(1)
	expectedUserId := uint64(2)
	expectedName := "test name"
	expectedDescription := "test description"
	r := New(expectedId, expectedUserId, expectedName, expectedDescription, nil)
	if r.Id() != expectedId {
		t.Errorf("New recipe Id %d not equal expected Id %d", r.Id(), expectedId)
	}
	if r.UserId() != expectedUserId {
		t.Errorf("New recipe UserId %d not equal expected UserId %d", r.Id(), expectedUserId)
	}
	if r.Name() != expectedName {
		t.Errorf("New recipe Name '%s' not equal expected Name '%s'", r.Name(), expectedName)
	}
	if r.Description() != expectedDescription {
		t.Errorf("New recipe Description '%s' not equal expected Description '%s'", r.Description(), expectedDescription)
	}
}

func TestRecipeString(t *testing.T) {
	expectedName := "test name"
	r := New(1, 1, expectedName, "test description", nil)
	expectedString := fmt.Sprintf("Recipe('%s')", expectedName)
	if r.String() != expectedString {
		t.Errorf("Recipe String '%s' not equal expected string '%s'", r.String(), expectedString)
	}
}
