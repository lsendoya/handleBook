package response

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lsendoya/handleBook/constants"
	"github.com/lsendoya/handleBook/model"
	"gorm.io/gorm"
	"net/http"
)

type Response struct {
}

func (r Response) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:   data,
		Errors: nil,
		Messages: model.Responses{
			{
				Code: constants.MessageOK, Message: "successful",
			},
		},
	}

}

func (r Response) Updated(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:   data,
		Errors: nil,
		Messages: model.Responses{
			{
				Code: constants.MessageOK, Message: "Entity updated successfully",
			},
		},
	}
}

func (r Response) Deleted(id uuid.UUID) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:   nil,
		Errors: nil,
		Messages: model.Responses{
			{
				Code: constants.MessageOK, Message: fmt.Sprintf("Entity with id %v deleted successfully", id),
			},
		},
	}
}

func (r Response) Forbidden() (int, model.MessageResponse) {

	return http.StatusForbidden, model.MessageResponse{
		Data: nil,
		Errors: model.Responses{
			{
				Code: constants.MessageForbidden, Message: "Access denied. Admin permissions required",
			},
		},
		Messages: nil,
	}
}

func (r Response) Unauthorized(err error) (int, model.MessageResponse) {

	return http.StatusUnauthorized, model.MessageResponse{
		Data: nil,
		Errors: model.Responses{
			{
				Code: constants.MessageUnauthorized, Message: err.Error(),
			},
		},
		Messages: nil,
	}
}

func (r Response) BadRequest(err error) (int, model.MessageResponse) {

	log.Warnf("%s", err.Error())
	return http.StatusBadRequest, model.MessageResponse{
		Data: nil,
		Errors: model.Responses{
			{
				Code: constants.MessageBadRequest, Message: err.Error(),
			},
		},
		Messages: nil,
	}
}

func (r Response) NotFound(err error) (int, model.MessageResponse) {

	log.Warnf("%s", err.Error())
	return http.StatusNotFound, model.MessageResponse{
		Data: nil,
		Errors: model.Responses{
			{
				Code: constants.MessageNotFound, Message: err.Error(),
			},
		},
		Messages: nil,
	}
}
func (r Response) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusCreated, model.MessageResponse{
		Data:   data,
		Errors: nil,
		Messages: model.Responses{
			{
				Code: constants.MessageCreated, Message: "Entity created successfully",
			},
		},
	}
}

func (r Response) InternalServerError(c echo.Context, who string, err error) (int, model.MessageResponse) {

	e := model.NewError()
	e.Err = err
	e.APIMessage = "internal error"
	e.Code = constants.MessageInternalServerError
	e.StatusHTTP = http.StatusInternalServerError
	e.Who = who

	log.Errorf("%s", e.Error())

	return http.StatusInternalServerError, model.MessageResponse{
		Data: nil,
		Errors: model.Responses{
			{
				Code: constants.MessageInternalServerError, Message: err.Error(),
			},
		},
		Messages: nil,
	}
}

func (r Response) ValidateErr(c echo.Context, who string, err error) error {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return c.JSON(r.NotFound(err))
		}

		return c.JSON(r.InternalServerError(c, who, err))
	}
	return nil
}
