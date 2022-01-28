package control

import (
	"errors"

	"github.com/mustafakocatepe/go-unit-test-example/model"
)

func ControlMedicalStatus(enjecteEdilenBarinak model.Shelter) (string, error) {

	getirilenHayvan, err := enjecteEdilenBarinak.GetAnimal()

	if err != nil {
		return "", errors.New("sorun")
	}

	if getirilenHayvan.MedicalStatus == "İyi" {
		return "Sağlıklı", nil
	}

	return "Sağlıksız", nil

}
