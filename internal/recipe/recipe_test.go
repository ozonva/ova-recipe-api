package recipe

import (
	"fmt"
	"github.com/golang/mock/gomock"
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

func TestRecipeCookCallDoActionWithoutError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAction := NewMockAction(mockCtrl)
	mockAction.EXPECT().DoAction().Return(nil).Times(1)
	r := New(1, 1, "test name", "test description", []Action{mockAction})
	if err := r.Cook(); err != nil {
		t.Errorf("Recipe Cook method returns not nil error: '%s'", err)
	}
}

func TestRecipeCookDoActionReturnsError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockActionOne := NewMockAction(mockCtrl)
	expectedError := fmt.Errorf("test expected error")
	mockActionOne.EXPECT().DoAction().Return(expectedError).Times(1)
	r := New(1, 1, "test name", "test description", []Action{mockActionOne})
	if err := r.Cook(); err != expectedError {
		t.Errorf("Error '%s' not equal expected error '%s'", err, expectedError)
	}
}

func TestRecipeCookSeveralActionsWOErrors(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockActionOne := NewMockAction(mockCtrl)
	mockActionOne.EXPECT().DoAction().Return(nil).Times(1)
	mockActionTwo := NewMockAction(mockCtrl)
	mockActionTwo.EXPECT().DoAction().Return(nil).Times(1)
	mockActionThree := NewMockAction(mockCtrl)
	mockActionThree.EXPECT().DoAction().Return(nil).Times(1)
	r := New(1, 1, "test name", "test description", []Action{mockActionOne, mockActionTwo, mockActionThree})
	if err := r.Cook(); err != nil {
		t.Errorf("Recipe Cook method returns not nil error: '%s'", err)
	}
}

func TestRecipeCookSeveralActionsWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockActionOne := NewMockAction(mockCtrl)
	mockActionOne.EXPECT().DoAction().Return(nil).Times(1)
	expectedError := fmt.Errorf("test expected error")
	mockActionTwo := NewMockAction(mockCtrl)
	mockActionTwo.EXPECT().DoAction().Return(expectedError).Times(1)
	mockActionThree := NewMockAction(mockCtrl)
	mockActionThree.EXPECT().DoAction().Return(nil).Times(0)
	r := New(1, 1, "test name", "test description", []Action{mockActionOne, mockActionTwo, mockActionThree})
	if err := r.Cook(); err != expectedError {
		t.Errorf("Error '%s' not equal expected error '%s'", err, expectedError)
	}
}
