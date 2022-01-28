package model

type Animal struct {
	Name          string
	Type          string
	MedicalStatus string
}

type Shelter interface {
	GetAnimal() (Animal, error)
}

type shelter struct {
	Animal Animal
}

func NewShelter() shelter {
	shelter := shelter{
		Animal: Animal{
			Name:          "Pamuk",
			Type:          "Kedi",
			MedicalStatus: "İyi",
		},
	}
	return shelter
}

func (s *shelter) GetAnimal() (Animal, error) {
	return s.Animal, nil
}

type MockShelter struct {
	MedicalStatus string
}

func (s *MockShelter) GetAnimal() (Animal, error) {
	return Animal{
		Name:          "test",
		Type:          "test",
		MedicalStatus: s.MedicalStatus,
	}, nil
}
