package services

import (
	"context"
	"fmt"
	"os"
	"waservice/helpers"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/go-redis/redis/v8"
)

const StatusOffline = 0
const StatusAktif = 1
const StatusLogin = 2
const StatusLoginFail = 3

type GenerateRepository struct {
	Rdb *redis.Client
	Ctx context.Context
	Wac *whatsapp.Conn
}

func (s *GenerateRepository) Generate(waUser *helpers.DataWAUser) (string, error) {
	s.Wac.SetClientVersion(3, 2123, 7)
	var qr = make(chan string)
	go func() {
		waUser.Token = <-qr
		waUser.Status = StatusLogin
		err := s.Rdb.Set(s.Ctx, "go_wa_user_"+waUser.UserId, helpers.ObToSTring(waUser), 0).Err()
		if err != nil {
			fmt.Println(err)
		}
	}()
	session, err := s.Wac.Login(qr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
		waUser.Status = StatusLoginFail
		waUser.Msg = err.Error()
		err := s.Rdb.Set(s.Ctx, "go_wa_user_"+waUser.UserId, helpers.ObToSTring(waUser), 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		return "", err
	}
	waUser.Session = session
	waUser.Status = StatusAktif

	err = s.Rdb.Set(s.Ctx, "go_wa_user_"+waUser.UserId, helpers.ObToSTring(waUser), 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	return "ok", nil
}
