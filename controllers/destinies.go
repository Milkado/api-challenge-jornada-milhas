package controllers

import (
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/Milkado/api-challenge-jornada-milhas/database"
	"github.com/Milkado/api-challenge-jornada-milhas/ent"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/destinies"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/testimonies"
	"github.com/Milkado/api-challenge-jornada-milhas/helpers"
	"github.com/labstack/echo/v4"
)

type (
	Destinies struct {
		Name        string  `json:"name" xml:"name" form:"name" validate:"required,min=1"`
		Price       float64 `json:"price" xml:"price" form:"price" validate:"required,numeric"`
		Meta        string  `json:"meta" xml:"meta" form:"meta" validate:"required,min=1,max=160"`
		Description *string `json:"description,omitempty" xml:"description,omitempty" form:"description,omitempty"`
	}
	Pictures struct {
		File []string `json:"file" xml:"file" form:"file" validate:"dive,base64"`
	}
	Picture struct {
		Picture string `json:"picture" xml:"picture" form:"picture" validate:"base64"`
	}
)

var (
	path = "/public/pictures/destinies/"
)

func IndexDestinies(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if c.QueryString() != "" {
		destinies, err := client.Destinies.Query().Where(destinies.NameContainsFold(c.QueryParam("name"))).WithTestimonies(
			func(q *ent.TestimoniesQuery) {
				q.Limit(3)
				q.Order(ent.Desc(testimonies.FieldCreatedAt))
			},
		).All(ctx)
		if err != nil {
			c.JSON(http.StatusAccepted, err)
		}
		return c.JSON(http.StatusOK, destinies)
	}

	destinies, err := client.Destinies.Query().WithTestimonies(
		func(q *ent.TestimoniesQuery) {
			q.Limit(3)
			q.Order(ent.Desc(testimonies.FieldCreatedAt))
		},
	).All(ctx)
	defer client.Close()
	if err != nil {
		c.JSON(http.StatusAccepted, err)
	}

	return c.JSON(http.StatusOK, destinies)
}

func ShowDestinies(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	client := database.ConnectDB()
	defer client.Close()
	destiny, sErr := client.Destinies.Query().Where(destinies.ID(id)).WithTestimonies(
		func(q *ent.TestimoniesQuery) {
			q.Limit(5)
			q.Order(ent.Desc(testimonies.FieldCreatedAt))
		},
	).Only(ctx)

	if sErr != nil {
		if ent.IsNotFound(sErr) {
			return c.JSON(http.StatusNotFound, sErr.Error())
		}
		c.JSON(http.StatusBadRequest, sErr.Error())
	}

	return c.JSON(http.StatusOK, destiny)
}

func StoreDestinies(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		d := new(Destinies)
		if bindErr := c.Bind(d); bindErr != nil {
			return c.JSON(http.StatusBadRequest, bindErr)
		}
		if validateErr := helpers.Validate(d, c); validateErr != nil {
			return c.JSON(http.StatusBadRequest, validateErr)
		}

		newDestiny, storeErr := client.Destinies.Create().SetName(d.Name).SetPrice(d.Price).SetMeta(d.Meta).SetNillableDescription(d.Description).Save(ctx)

		if storeErr != nil {
			return c.JSON(http.StatusBadRequest, storeErr)
		}

		return c.JSON(http.StatusCreated, newDestiny)
	}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return nil
}

func ChangeDestiny(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		id := idToInt(c, c.Param("id"))
		d := new(Destinies)
		if err := c.Bind(d); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		destiny, err := client.Destinies.UpdateOneID(id).SetName(d.Name).SetPrice(d.Price).SetMeta(d.Meta).SetNillableDescription(d.Description).Save(ctx)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, destiny)
	}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return nil
}

func DeleteDestiny(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		deleteErr := client.Destinies.DeleteOneID(id).Exec(ctx)
		if deleteErr != nil {
			return c.JSON(http.StatusBadRequest, deleteErr)
		}

		return c.JSON(http.StatusOK, "Deleted")
	}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return nil
}

func StoreDestinyPicture(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		p := new(Picture)
		if bindErr := c.Bind(p); bindErr != nil {
			return c.JSON(http.StatusBadRequest, bindErr)
		}

		destinyID := idToInt(c, c.Param("id"))
		var fileName string
		var err error
		if p.Picture != "" {
			if validateErr := helpers.Validate(p, c); validateErr != nil {
				return c.JSON(http.StatusBadRequest, validateErr)
			}

			base64File := p.Picture
			fileName, err = storeBase64Picture(base64File, c, path)
			if err != nil {
				return c.JSON(http.StatusBadRequest, []string{err.Error(), fileName})
			}
		} else {
			file, fErr := c.FormFile("picture")
			if fErr != nil {
				return c.JSON(http.StatusBadRequest, []string{fErr.Error(), "no file"})
			}
			fileName, err = storeFormPicture(file, c, path)
			if err != nil {
				return c.JSON(http.StatusBadRequest, []string{err.Error(), fileName})
			}
		}

		newPicture, err := client.DestinyPictures.Create().SetDestinyID(destinyID).SetPicture(fileName).SetPath(path).Save(ctx)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, newPicture)
	}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return nil
}

func StoreManyDestinyPictures(c echo.Context) error {
	client := database.ConnectDB()
	defer client.Close()

	if err := withTx(ctx, client, func(tx *ent.Tx) error {
		p := new(Pictures)
		if bindErr := c.Bind(p); bindErr != nil {
			return c.JSON(http.StatusBadRequest, bindErr)
		}
		if validateErr := helpers.Validate(p, c); validateErr != nil {
			return c.JSON(http.StatusBadRequest, validateErr)
		}

		destinyID := idToInt(c, c.Param("id"))
		if p.File != nil {
			for _, picture := range p.File {
				base64File := picture
				fileName, err := storeBase64Picture(base64File, c, path)
				if err != nil {
					return c.JSON(http.StatusBadRequest, []string{err.Error(), fileName})
				}
				_, err = client.DestinyPictures.Create().SetDestinyID(destinyID).SetPicture(fileName).SetPath(path).Save(ctx)

				if err != nil {
					return c.JSON(http.StatusBadRequest, []string{err.Error(), fileName})
				}
			}
		} else {
			var files []*multipart.FileHeader
			form, fErr := c.MultipartForm()
			if fErr != nil {
				return c.JSON(http.StatusBadRequest, []string{fErr.Error(), "no file"})
			}
			for _, fileHeaders := range form.File {
				files = append(files, fileHeaders...)
			}

			if len(files) == 0 {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "no files provided"})
			}

			for _, fileHeader := range files {
				fileName, err := storeFormPicture(fileHeader, c, path)

				if err != nil {
					return c.JSON(http.StatusBadRequest, []string{err.Error(), fileName})
				}

				_, err = client.DestinyPictures.Create().SetDestinyID(destinyID).SetPicture(fileName).SetPath(path).Save(ctx)

				if err != nil {
					return c.JSON(http.StatusBadRequest, err)
				}
			}
		}

		return c.JSON(http.StatusCreated, "Pictures created")
	}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return nil
}
