package infrastructure

import "github.com/spf13/viper"

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
}

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type JWTConfig struct {
	Secret string
}

var AppCfg Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		panic("cannot read config: " + err.Error())
	}

	AppCfg = Config{
		App: AppConfig{
			Port: viper.GetString("app.port"),
		},
		DB: DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Name:     viper.GetString("db.name"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
		},
		JWT: JWTConfig{
			Secret: viper.GetString("jwt.secret"),
		},
	}
}
