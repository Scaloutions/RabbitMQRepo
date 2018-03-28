package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func ReadInConfig() *viper.Viper {

	v := viper.New()
	v.AddConfigPath("./app/config")
	v.SetConfigType("toml")

	v.SetConfigName("app")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	readInConfigHelper(v, "app.docker")
	readInConfigHelper(v, "app.development")
	readInConfigHelper(v, "app.production")
	readInConfigHelper(v, "app.qa")
	readInConfigHelper(v, "messages")
	readInConfigHelper(v, "maps")

	return v
}

func RabbitmqConnect() {
	time.Sleep(1 * time.Minute)
	connection, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		fmt.Errorf("Dial: %s", err)
	}
	fmt.Println("Connection: ", *connection)
}

/*
	Private methods
*/

func readInConfigHelper(v *viper.Viper, fileName string) {

	v.SetConfigName(fileName)
	err := v.MergeInConfig()
	if err != nil {
		log.Fatalln(err)
	}

}
