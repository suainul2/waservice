package configs

type redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"DB"`
}

type mysql struct {
	Address string `mapstructure:"address"`
}

var Redis redis = redis{}
var Mysql mysql = mysql{}
