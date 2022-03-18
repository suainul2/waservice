package services

import (
	"fmt"
	"waservice/helpers"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

// 6282334976825
func (s *GenerateRepository) SendImage(no string, imgUrl string, txt string) error {
	s.Wac.SetClientVersion(3, 2123, 7)
	image, err := helpers.DownloadMediaFromURL(imgUrl)
	if err != nil {
		return err
	}
	text := whatsapp.ImageMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: no + "@s.whatsapp.net",
		},
		Type:    "image/jpeg",
		Caption: txt,
		Content: image,
	}
	tes, err := s.Wac.Send(text)
	if err != nil {
		return err
	}
	fmt.Println(tes)
	return nil
}
