package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	env := os.Getenv("ENV")
	fmt.Println(env)

	port := os.Getenv("PORT")
	fmt.Println(port)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err = viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// Memungkinkan viper dapat otomatis membaca config dari .env
	// contoh MYSQL_DBHOST akan dibaca viper.GetString("mysql.dbhost")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	fmt.Println(viper.GetString("env"))
	fmt.Println(viper.GetString("mysql.db_host"))
}
