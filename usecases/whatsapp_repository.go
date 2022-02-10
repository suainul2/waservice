package usecases

// A UserRepository belong to the usecases layer.
type GenerateRepository interface {
	Generate(id uint) (string, error)
	Resession(id uint) (string, error)
	SendMessage(no string, txt string, id uint) error
	SendImage(no string, imgUrl string, txt string, id uint) error
}
