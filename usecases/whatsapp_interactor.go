package usecases

type GenerateInteractor struct {
	GenerateRepository GenerateRepository
}

// Index is display a listing of the resource.
func (gi *GenerateInteractor) Generate(id uint) (data string, err error) {
	data, err = gi.GenerateRepository.Generate(id)
	return
}
func (gi *GenerateInteractor) Resession(id uint) (data string, err error) {
	data, err = gi.GenerateRepository.Resession(id)
	return
}
func (gi *GenerateInteractor) SendMessage(no string, txt string, id uint) (err error) {
	err = gi.GenerateRepository.SendMessage(no, txt, id)
	return
}
func (gi *GenerateInteractor) SendImage(no string, imgUrl string, txt string, id uint) (err error) {
	err = gi.GenerateRepository.SendImage(no, imgUrl, txt, id)
	return
}
