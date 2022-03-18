package services

import (
	"fmt"
	"waservice/helpers"
)

func (s *GenerateRepository) Resession(waUser *helpers.DataWAUser) (string, error) {
	s.Wac.SetClientVersion(3, 2123, 7)
	newSess, err := s.Wac.RestoreWithSession(waUser.Session)
	if err != nil {
		if err.Error() == "already logged in" {
			return "ok", nil
		}
		if err.Error() == "restore session connection timed out" {
			waUser.Msg = "WA anda Offline"
			return "", err

		}
		fmt.Println("error session " + err.Error())
		waUser.Msg = err.Error()
		waUser.Status = StatusOffline
		errr := s.Rdb.Set(s.Ctx, "go_wa_user_"+waUser.UserId, helpers.ObToSTring(waUser), 0).Err()
		if errr != nil {
			fmt.Println(err)
		}
		return "", err
	}
	waUser.Session = newSess
	waUser.Status = StatusAktif
	err = s.Rdb.Set(s.Ctx, "go_wa_user_"+waUser.UserId, helpers.ObToSTring(waUser), 0).Err()
	if err != nil {
		return "", err
	}
	return "ok", nil
}
