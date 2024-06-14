package category

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.POST("/category", func(c echo.Context) error {
		name := c.QueryParam("name")

		id, err := dbAdd(name)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, fmt.Sprintf("Id: %v Name: %s", id, name))
	})

	e.GET("/category", func(c echo.Context) error {
		categories, err := dbGetAll()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, categories)
	})

	e.GET("/category/:id", func(c echo.Context) error {
		id := c.Param("id")
		category, err := dbGetById(id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, category)

	})

	e.PUT("/category/:id", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, name)
	})
	e.DELETE("/category/:id", func(c echo.Context) error {
		id := c.Param("id")

		err := dbDelete(id)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, fmt.Sprintf("Deleted category %s", id))
	})
}
