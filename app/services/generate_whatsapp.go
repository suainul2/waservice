package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"waservice/helpers"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/go-redis/redis/v8"
)

type GenerateRepository struct {
	Rdb *redis.Client
	Ctx context.Context
	Wac *whatsapp.Conn
}

func (s *GenerateRepository) Generate(id uint) (string, error) {
	s.Wac.SetClientVersion(3, 2123, 7)
	var qr = make(chan string)
	var newId string = helpers.UintToStr(id)
	go func() {
		err := s.Rdb.Set(s.Ctx, "go_con_"+newId, "C", 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		err = s.Rdb.Set(s.Ctx, "go_red_"+newId, <-qr, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
	}()
	session, err := s.Wac.Login(qr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
		err := s.Rdb.Set(s.Ctx, "go_red_"+newId, "sessionFail ", 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		err = s.Rdb.Set(s.Ctx, "go_con_"+newId, "X", 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		return "", err
	}
	b, err := json.Marshal(session)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = s.Rdb.Set(s.Ctx, "go_ses_"+newId, string(b), 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	return "ok", nil
}
