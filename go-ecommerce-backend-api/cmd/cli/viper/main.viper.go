package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config") // path to config
	viper.SetConfigName("local")    // name of config file (without extension)
	viper.SetConfigType("yaml")     // type of config file
	//read file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}
	//read file server
	fmt.Println("Server Port: ", viper.GetInt("server.port"))
	fmt.Println("Sercu: ", viper.GetString("security.jwt.key"))

	// config structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration: %v \n", err)
	}
	fmt.Println("Server Port: ", config.Server.Port)
	for _, db := range config.Database {
		fmt.Println("User: ", db.User)
		fmt.Println("Password: ", db.Password)
		fmt.Println("Host: ", db.Host)

	}

}
