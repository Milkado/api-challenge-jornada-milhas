package atlas

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
)

// Command to migrate from atlas

func Migrate() {
	dir, url := buildMigrate()
	cmd := exec.Command("atlas", "migrate", "apply", "--dir", dir, "--url", url)

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

func buildMigrate() (string, string) {
	dir := "file://ent/migrate/migrations"
	url := helpers.Env("DB_CONNECTION") + "://" + helpers.Env("DB_USER") + ":" + helpers.Env("DB_PASSWORD") + "@" + helpers.Env("DB_HOST") + ":" + helpers.Env("DB_PORT") + "/" + helpers.Env("DB_NAME")

	return dir, url
}
