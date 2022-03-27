package helper

import "gokes/app/models"

func ValidateFakes(fakesType string) bool {

	validFakesType := []string{
		models.FakesStatusKlinik,
		models.FakesStatusPosyandu,
		models.FakesStatusPuskesmas,
		models.FakesStatusRumahSakit,
	}

	for i := range validFakesType {
		if fakesType == validFakesType[i] {
			return true
		}
	}

	return false

}
