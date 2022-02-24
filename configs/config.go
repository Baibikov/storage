package configs

type Conf struct {
	DB DB `yaml:"db"`
}

type DB struct {
	Conn string `yaml:"conn"`
}