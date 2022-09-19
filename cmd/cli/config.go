package cli

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig AppConfig
	DB        DB
}

type AppConfig struct {
	Address         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type DB struct {
	Driver   string
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	Schema   string
}

func GetConfig() (Config, error) {
	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name(), file.IsDir())
	// }

	var conf Config
	viper.AddConfigPath(".")
	viper.SetConfigFile("app.env")

	// viper.AutomaticEnv()

	if err := viper.ReadInConfig(); nil != err {
		return conf, err
	}

	conf.AppConfig = AppConfig{
		Address:         viper.GetString("APP_ADDRESS"),
		WriteTimeout:    viper.GetDuration("APP_WRITE_TIMEOUT"),
		ReadTimeout:     viper.GetDuration("APP_READ_TIMEOUT"),
		ShutdownTimeout: viper.GetDuration("APP_SHUTDOWN_TIMEOUT"),
	}

	conf.DB = DB{
		Driver:   viper.GetString("DB_DRIVER"),
		Host:     viper.GetString("DB_HOST"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		Name:     viper.GetString("DB_NAME"),
		Port:     viper.GetString("DB_PORT"),
		Schema:   viper.GetString("DB_SCHEMA"),
	}

	// if err := viper.Unmarshal(&conf); nil != err {
	// 	return conf, err
	// }

	return conf, nil
}

func (d DB) String() string {
	return fmt.Sprintf("DB_DRIVER: %s\nDB_HOST: %s\nDB_USER: %s\nDB_PASSWORD: %s\nDB_NAME: %s\nDB_PORT: %s\nDB_SCHEMA: %s\n",
		d.Driver, d.Host, d.User, d.Password, d.Name, d.Port, d.Schema)
}

func (w AppConfig) String() string {
	return fmt.Sprintf("app_ADDRESS: %s\napp_READ_TIMEOUT: %s\napp_WRITE_TIMEOUT: %s\n"+
		"app_SHUTDOWN_TIMEOUT: %s\n", w.Address, w.ReadTimeout, w.WriteTimeout, w.ShutdownTimeout)
}
