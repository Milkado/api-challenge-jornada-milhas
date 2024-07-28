package controllers

import (
	"net/http"

	"github.com/Milkado/api-challenge-jornada-milhas/database"
	"github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/user"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/Milkado/api-challenge-jornada-milhas/mail"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tjarratt/babble"
)

type (
	CreatePassword struct {
		Password        string `json:"password" xml:"password" form:"password" validate:"required,min=6"`
		PasswordConfirm string `json:"password_confirm" xml:"password_confirm" form:"password_confirm" validate:"required"`
	}
	PasswordDTO struct {
		Password     string
		RandSecurity string
	}
	RequestPassword struct {
		Email string `json:"email" xml:"email" form:"email" validate:"required,email"`
	}
	MailData struct {
		Name string
		Link string
	}
)

func StorePassword(c echo.Context) error {
	client := database.ConnectDB()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		createPass := new(CreatePassword)
		if err := c.Bind(createPass); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := helpers.Validate(createPass, c); err != nil {
			return err
		}
		if createPass.Password != createPass.PasswordConfirm {
			return c.JSON(http.StatusBadRequest, "The password confirmation does't match")
		}
		token := c.Param("token")
		userGet, uErr := client.User.Query().Where(user.PasswordToken(token)).Only(ctx)

		if uErr != nil {
			return c.JSON(http.StatusBadRequest, uErr.Error())
		}

		newPass := new(PasswordDTO)
		babbler := babble.NewBabbler()
		newPass.RandSecurity = babbler.Babble()
		newPass.Password = hashPass(createPass.Password + newPass.RandSecurity)
		updateErr := userGet.Update().ClearPasswordToken().SetPassword(newPass.Password).SetRandSecurity(newPass.RandSecurity).Exec(ctx)

		if updateErr != nil {
			return c.JSON(http.StatusInternalServerError, updateErr.Error())
		}

		return c.JSON(http.StatusOK, "Passwotd created")
	}); err != nil {
		defer client.Close()
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	defer client.Close()
	return nil
}

func RequestNewPassword(c echo.Context) error {
	client := database.ConnectDB()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		email := new(RequestPassword)
		if err := c.Bind(email); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := helpers.Validate(email, c); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		userGet, getErr := client.User.Query().Where(user.Email(email.Email)).Only(ctx)
		if getErr != nil {
			return c.JSON(http.StatusBadRequest, getErr)
		}
		newToken := uuid.New()
		userGet, updateErr := userGet.Update().SetPasswordToken(newToken.String()).Save(ctx)
		if updateErr != nil {
			return c.JSON(http.StatusInternalServerError, updateErr)
		}

		mailData := new(MailData)
		mailData.Name = userGet.Name
		mailData.Link = helpers.Env("FRONT_URL") + *userGet.PasswordToken

		mailErr := mail.SendMail("Criar nova senha", helpers.ResetMail, []string{userGet.Email}, mailData)

		if mailErr != nil {
			panic(mailErr.Error())
		}
		// if err := mail.SendMail("Criar nova senha", helpers.ResetMail, []string{userGet.Email}, mailData); err != nil {
		// 	return c.JSON(http.StatusInternalServerError, err)
		// }

		return c.JSON(http.StatusAccepted, "Email enviado")
	}); err != nil {

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	defer client.Close()
	return nil
}
