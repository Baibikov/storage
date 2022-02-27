package configs

type Conf struct {
	DB      DB      `yaml:"db"`
	HTTP    HTTP    `yaml:"http"`
	Storage Storage `yaml:"storage"`
}

type DB struct {
	Conn string `yaml:"conn"`
}

type HTTP struct {
	Port string `yaml:"port"`
}

type Storage struct {
	SRC string `yaml:"src"`
}
