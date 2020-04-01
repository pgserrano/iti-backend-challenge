package enviroment

import "os"

const SYSTEM_PORT = "3000"


func GetSystemPort() string {

	system_port := os.Getenv("PORT")

	if system_port == "" {
		system_port = SYSTEM_PORT
	}

	return ":" + system_port

}