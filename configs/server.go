package configs

type server struct {
	Port    string `mapstructure:"port"`
	AppName string `mapstructure:"appName"`
	Key     string `mapstructure:"key"`
}

var Server server = server{}
