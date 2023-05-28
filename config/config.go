package config

type ServerConfig struct {
	Name         string      `mapstructure:"name"`
	Port         int         `mapstructure:"port"`
	MysqlInfo    MysqlConfig `mapstructure:"mysql"`
	RedisInfo    RedisConfig `mapstructure:"redis"`
	LogsLocation string      `mapstructure:"logsLocation"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
