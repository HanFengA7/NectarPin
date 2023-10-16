package conf

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
}

type Mysql struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DbName     string `yaml:"dbName"`
	DbUser     string `yaml:"dbUser"`
	DbPassword string `yaml:"dbPassword"`
}

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"showLine"`
	LogInConsole bool   `yaml:"logInConsole"`
}
