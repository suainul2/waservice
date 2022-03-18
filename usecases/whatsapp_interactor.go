package usecases

import "waservice/helpers"

type GenerateInteractor struct {
	GenerateRepository GenerateRepository
}

// Index is display a listing of the resource.
func (gi *GenerateInteractor) Generate(waUser *helpers.DataWAUser) (data string, err error) {
	data, err = gi.GenerateRepository.Generate(waUser)
	return
}
func (gi *GenerateInteractor) Resession(waUser *helpers.DataWAUser) (data string, err error) {
	data, err = gi.GenerateRepository.Resession(waUser)
	return
}
func (gi *GenerateInteractor) SendMessage(no string, txt string) (err error) {
	err = gi.GenerateRepository.SendMessage(no, txt)
	return
}
func (gi *GenerateInteractor) SendImage(no string, imgUrl string, txt string) (err error) {
	err = gi.GenerateRepository.SendImage(no, imgUrl, txt)
	return
}
