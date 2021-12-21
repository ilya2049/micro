package handler

import (
	"hasherapi/internal/system/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type HashHandler struct {
}

func (h Handler) GetCheck(params operations.GetCheckParams) middleware.Responder {
	return nil
}

func (h Handler) PostSend(params operations.PostSendParams) middleware.Responder {
	return nil
}
