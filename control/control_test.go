package control

import (
	"testing"

	"github.com/mustafakocatepe/go-unit-test-example/model"
)

func TestControlMedicalStatus(t *testing.T) {

	t.Run("If animal is healthy", func(t *testing.T) {
		mockShelter := model.MockShelter{MedicalStatus: "İyi"}

		want := "Sağlıklı"

		got, _ := ControlMedicalStatus(&mockShelter)
		if got != want {
			t.Errorf("Test fail! want: '%s', got: '%s'", want, got)
		}
	})

	t.Run("If animal is not healthy", func(t *testing.T) {
		mockShelter := model.MockShelter{MedicalStatus: "Kötü"}

		want := "Sağlıksız"

		got, _ := ControlMedicalStatus(&mockShelter)

		if got != want {
			t.Errorf("Test fail! want: '%s', got: '%s'", want, got)
		}
	})
}
