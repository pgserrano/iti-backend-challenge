package main

import (
	"fmt"
	"github.com/pgserrano/iti-backend-challenge/configs/version"
	"github.com/pgserrano/iti-backend-challenge/internal/routes"
)

func main() {

	fmt.Println("Starting Service Version: " + version.Version)
	routes.StartServer()

}
