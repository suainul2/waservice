package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Rhymen/go-whatsapp"
)

type DataWAUser struct {
	UserId  string
	Session whatsapp.Session
	Token   string
	Status  int
	Msg     string
}

func UintToStr(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}

func StrToInt(s string) int {
	d, _ := strconv.Atoi(s)
	return d
}

func ObToSTring(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func StrToObWA(wa string) *DataWAUser {
	data := DataWAUser{}
	json.Unmarshal([]byte(wa), &data)
	return &data
}
