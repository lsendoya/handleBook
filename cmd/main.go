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

	err = e.Start(":" + Config("SERVER_PORT"))

}
