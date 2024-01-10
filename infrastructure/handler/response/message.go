package response

import (
	"github.com/lsendoya/handleBook/constants"
	"github.com/lsendoya/handleBook/model"
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

func (r Response) Deleted() (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:   nil,
		Errors: nil,
		Messages: model.Responses{
			{
				Code: constants.MessageOK, Message: "Entity deleted successfully",
			},
		},
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

func (r Response) BadRequest(err error) (int, model.MessageResponse) {
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

func (r Response) InternalServerError(err error) (int, model.MessageResponse) {
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
