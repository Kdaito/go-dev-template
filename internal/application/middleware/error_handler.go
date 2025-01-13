package middleware

import (
	"net/http"

	"github.com/Kdaito/kinodokuna-be/internal/lib/errors"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code int `json:"code"`
	Message interface{} `json:"message"`
}

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	// カスタムエラーの場合
	if ce, ok := err.(*errors.Error); ok {
		c.Logger().Error(ce)
		if JSONErr := c.JSON(ce.Code, ErrorResponse{
			Code: ce.Code,
			Message: http.StatusText(ce.Code),
		}); JSONErr != nil {
			c.Logger().Error(JSONErr)
		}
		return
	}

	// echo.HTTPErrorの場合
	if he, ok := err.(*echo.HTTPError); ok {
		c.Logger().Error(he)
		if JSONErr := c.JSON(he.Code, ErrorResponse{
			Code: he.Code,
			Message: http.StatusText(he.Code),
		}); JSONErr != nil {
			c.Logger().Error(JSONErr)
		}
		return
	}

	// その他のエラーの場合
	c.Logger().Error(err)
	if JSONErr := c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code: http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}); JSONErr != nil {
		c.Logger().Error(JSONErr)
	}
}