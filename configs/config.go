package configs

type Conf struct {
	DB   DB   `yaml:"db"`
	HTTP HTTP `yaml:"http"`
}

type DB struct {
	Conn string `yaml:"conn"`
}

type HTTP struct {
	Port string `yaml:"port"`
}
