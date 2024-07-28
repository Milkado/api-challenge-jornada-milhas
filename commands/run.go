package main

import (
	"flag"
	"fmt"

	"github.com/Milkado/api-challenge-jornada-milhas/commands/atlas"
	"github.com/Milkado/api-challenge-jornada-milhas/commands/jwt"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
)

// Just to eliminate the need to write a big command with big args
func main() {
	var command string
	var name string

	flag.StringVar(&command, "command", "", "Choose the command to run.")
	flag.StringVar(&name, "migration_name", "", "Name of the migration")
	flag.Parse()

	if command == "" {
		fmt.Println(helpers.Red + "--command flag missing, choose one of the avaible:" + helpers.Reset)
		fmt.Println(helpers.Green + "migration_create, required the --name flag")
		fmt.Println("migrate" + helpers.Reset)
		return
	}

	switch command {
	case "migration_create":
		atlas.CreateMigration(name)
	case "migrate":
		atlas.Migrate()
	case "generate_secret":
		jwt.GenerateSecret()
	default:
		fmt.Println(helpers.Red + "Command not avaible" + helpers.Reset)
	}
}
