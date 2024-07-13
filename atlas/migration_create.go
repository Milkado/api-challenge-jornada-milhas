package atlas

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
)

//create aux to create atlas migrations with env variables

func CreateMigration(migrationName string) {
	dir, to, devUrl := buildCreate()
	cmd := exec.Command("atlas", "migrate", "diff", migrationName, "--dir", dir, "--to", to, "--dev-url", devUrl)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(helpers.Red + err.Error())
		fmt.Println(stderr.String() + helpers.Reset)
		return
	}

	fmt.Println(out.String())
}

func buildCreate() (string, string, string) {
	dir := "file://ent/migrate/migrations"
	to := "ent://ent/schema"
	devUrl := helpers.Env("DB_CONNECTION") + "://" + helpers.Env("DB_USER") + ":" + helpers.Env("DB_PASSWORD") + "@" + helpers.Env("DB_HOST") + ":" + helpers.Env("DB_PORT") + "/" + helpers.Env("DB_DEV")

	return dir, to, devUrl
}
