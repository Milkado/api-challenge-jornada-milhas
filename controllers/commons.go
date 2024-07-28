package controllers

import (
	"context"
	"fmt"

	"github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

var ctx = context.Background()
var dailer = gomail.NewDialer(helpers.Env("MAIL_HOST"), helpers.EnvInt("MAIL_PORT"), helpers.Env("MAIL_USER"), helpers.Env("MAIL_PASS"))

func withTx(c context.Context, client ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(c)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			if helpers.Env("DEBUG") == "false" {
				panic(v)
			}
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}

		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func hashPass(pass string) string {
	hashed, hashErr := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if hashErr != nil {
		return hashErr.Error()
	}
	return string(hashed)
}

func checkHash(pass string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	if err != nil {
		return false
	}

	return true
}
