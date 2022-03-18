package usecases

import "waservice/helpers"

// A UserRepository belong to the usecases layer.
type GenerateRepository interface {
	Generate(waUser *helpers.DataWAUser) (string, error)
	Resession(waUser *helpers.DataWAUser) (string, error)
	SendMessage(no string, txt string) error
	SendImage(no string, imgUrl string, txt string) error
}
