package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration: %v \n", err)
	}
}
