package configs

type redis struct {
	Address  string
	Password string
	DB       int
}

type mysql struct {
	Address string
}

var Redis redis = redis{}
var Mysql mysql = mysql{}
