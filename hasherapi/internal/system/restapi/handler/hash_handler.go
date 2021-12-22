package handler

import (
	"hasherapi/internal/domain/hash"
	"hasherapi/internal/system/restapi/middlewares"
	"hasherapi/internal/system/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func newHashHandler(hashService *hash.Service) *hashHandler {
	return &hashHandler{
		hashService: hashService,
	}
}

type hashHandler struct {
	hashService *hash.Service
}

func (h *Handler) GetCheck(params operations.GetCheckParams) middleware.Responder {
	hashIDs, err := hash.NewIDsFromStrings(params.Ids)
	if err != nil {
		return middlewares.NewBadRequestErrorResponder(operations.NewGetCheckBadRequest(), err)
	}

	ctx := params.HTTPRequest.Context()

	identifiedHashes, err := h.hashService.FindHashes(ctx, hashIDs)
	if err != nil {
		return middlewares.NewInternalErrorResponder(operations.NewGetCheckInternalServerError(), err)
	}

	if len(identifiedHashes) == 0 {
		return operations.NewGetCheckNoContent()
	}

	return operations.NewGetCheckOK().WithPayload(identifiedHashesToArrayOfHash(identifiedHashes))
}

func (h *Handler) PostSend(params operations.PostSendParams) middleware.Responder {
	hashInputs := postSendParamsToHashInputs(params)

	ctx := params.HTTPRequest.Context()

	identifiedHashes, err := h.hashService.CreateHashes(ctx, hashInputs)
	if err != nil {
		return operations.NewPostSendInternalServerError()
	}

	return operations.NewPostSendOK().WithPayload(identifiedHashesToArrayOfHash(identifiedHashes))
}
