package jwt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/google/uuid"
)

func GenerateSecret() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	fileContent := ""
	scanner := bufio.NewScanner(file)
	updated := false

	for scanner.Scan() {
		line := scanner.Text()
		if !updated && strings.Contains(line, "JWT_SECRET") {
			line = `JWT_SECRET="` + uuid.NewString() + `"`
			updated = true
		}

		fileContent += line + "\n"
	}

	if !updated {
		fileContent += `JWT_SECRET="` + uuid.NewString() + `"` + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(".env", []byte(fileContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.Green + "JWT_SECRET generated" + helpers.Reset)
}
