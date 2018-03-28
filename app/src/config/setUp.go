package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

const (
	wait_time = 15000
)

var (
	count = 0
)

func ReadInConfig() *viper.Viper {

	// TODO: add error handling later, return err instead of constructing new one

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

func RabbitmqConnect(amqpUri string) *amqp.Connection {

	// TODO: add error handling later, return err instead of constructing new one

	if count == 0 { // wait until Rabbitmq is ready
		time.Sleep(wait_time * time.Millisecond)
		count = 1
	}
	conn, err := amqp.Dial(amqpUri)
	if err != nil {
		fmt.Errorf("Dial: %s", err)
	}
	return conn
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
