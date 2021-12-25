package handler

import (
	"common/errors"
	"hasherapi/domain/hash"
	"hasherapi/system/restapi/middlewares"
	"hasherapi/system/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func newHashHandler(
	hashService *hash.Service,
	errorResponderFactory *middlewares.ResponderFactory,
) *hashHandler {
	return &hashHandler{
		hashService:           hashService,
		errorResponderFactory: errorResponderFactory,
	}
}

type hashHandler struct {
	hashService *hash.Service

	errorResponderFactory *middlewares.ResponderFactory
}

func (h *Handler) GetCheck(params operations.GetCheckParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	hashIDs, err := hash.NewIDsFromStrings(params.Ids)
	if err != nil {
		return h.errorResponderFactory.NewBadRequestErrorResponder(
			ctx, operations.NewGetCheckBadRequest(), errors.Errorf("failed to parse hash ids: %w", err),
		)
	}

	identifiedHashes, err := h.hashService.FindHashes(ctx, hashIDs)
	if err != nil {
		return h.errorResponderFactory.NewInternalErrorResponder(
			ctx, operations.NewGetCheckInternalServerError(), err,
		)
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
		return h.errorResponderFactory.NewInternalErrorResponder(
			ctx, operations.NewPostSendInternalServerError(), err,
		)
	}

	return operations.NewPostSendOK().WithPayload(identifiedHashesToArrayOfHash(identifiedHashes))
}
