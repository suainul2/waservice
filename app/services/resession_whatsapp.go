package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"waservice/databases"
	"waservice/helpers"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

func (s *GenerateRepository) Resession(id uint) (string, error) {
	s.Wac.SetClientVersion(3, 2123, 7)
	var newId string = helpers.UintToStr(id)
	dataSession, err := databases.Rdb.Get(databases.Ctx, "go_ses_"+newId).Result()
	if err != nil {
		panic(err.Error())
	}
	handleErrConnection, err := databases.Rdb.Get(databases.Ctx, "go_con_"+newId).Result()
	if err != nil {
		panic(err.Error())
	}
	if handleErrConnection == "X" {
		return "", errors.New("session fail")
	}
	data := whatsapp.Session{}
	json.Unmarshal([]byte(dataSession), &data)
	newSess, err := s.Wac.RestoreWithSession(data)
	if err != nil {
		if err.Error() == "already logged in" {
			return "ok", nil
		}
		fmt.Println("error session " + err.Error())

		err = s.Rdb.Set(s.Ctx, "go_con_"+newId, "X", 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		return "", err
	}

	b, err := json.Marshal(newSess)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = s.Rdb.Set(s.Ctx, "go_ses_"+newId, string(b), 0).Err()
	if err != nil {
		return "", err
	}
	return "ok", nil
}
