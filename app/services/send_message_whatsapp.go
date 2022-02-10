package services

import (
	"fmt"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

// 6282334976825
func (s *GenerateRepository) SendMessage(no string, txt string, id uint) error {
	s.Wac.SetClientVersion(3, 2123, 7)
	text := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: no + "@s.whatsapp.net",
		},
		Text: txt,
	}

	tes, err := s.Wac.Send(text)
	if err != nil {
		return err
	}
	fmt.Println(tes)
	return nil
}
