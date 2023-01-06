package main

import (
	"github/practice/pkg/web"
	"log"
	"os"
)

func WebAppSetup() {

	// Create a PhoneBook web application
	app := web.WebPhoneBook{
		InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Host: ":9090",
	}

	// Set the routers for web application
	app.RoutersSetup()

	// Start up the web server
	app.ServerStartup()
}


func main() {
	appType := os.Args[1]
	
	if appType == "Web" {
		WebAppSetup()
	}

}