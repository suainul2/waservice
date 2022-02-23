package configs

import (
	"waservice/helpers"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func Handle() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	err := config.LoadFiles(helpers.Pwd + "/config.yml")
	if err != nil {
		panic(err)
	}
	config.BindStruct("server", &Server)
	config.BindStruct("redis", &Redis)
	config.BindStruct("mysql", &Mysql)

}
