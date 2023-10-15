package conf

type Mysql struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DbName     string `yaml:"dbName"`
	DbUser     string `yaml:"dbUser"`
	DbPassword string `yaml:"dbPassword"`
}
