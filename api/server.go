package api

import (
	"fmt"
	"log"
	"os"

	"github.com/jnanendraveer/transactions-golang-app/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func doesFileExist(fileName string) {

	// check if error is "file not exists"
	if _, err := os.Stat(fileName); err == nil {
		if err = godotenv.Load(fileName); err != nil {
			log.Fatalf("Error getting env, not comming through %v", err)
		} else {
			fmt.Println(fmt.Sprintf("We are getting the %s values", fileName))
		}
	}
}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// doesFileExist(".fp_live")
	doesFileExist(".fp_dev")
	// doesFileExist(".fp_local")
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run(":9072")

}
