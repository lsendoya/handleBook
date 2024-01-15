package main

import (
	"github.com/lsendoya/handleBook/infrastructure/handler"
	"log"
)

func main() {
	err := validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	e := newHTTP()

	db, errConn := ConnectDB()
	if err != nil {
		log.Fatal(errConn)
	}

	err = MigrateDB(db)
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(e, db)

	port := Config("SERVER_PORT")

	if Config("IS_HTTPS") == "true" {
		err = e.StartTLS(port, Config("CERT_PEM_FILE"), Config("KEY_PEM"))
	} else {
		err = e.Start(":" + port)
		if err != nil {
			log.Fatal(err)
		}
	}

}
