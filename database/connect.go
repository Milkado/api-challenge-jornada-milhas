package database

import (
	"log"

	"github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() ent.Client {
	connection := helpers.Env("DB_USER") + ":" + helpers.Env("DB_PASSWORD") + "@tcp(" + helpers.Env("DB_HOST") + ":" + helpers.Env("DB_PORT") + ")/" + helpers.Env("DB_NAME") + "?parseTime=true"
	client, err := ent.Open("mysql", connection)
	if err != nil {
		log.Fatalf(helpers.Red+"failed to open connection to mysql: %v"+helpers.Reset, err)
	}

	return *client
}
