package handlers

import (
	"github.com/gin-gonic/gin"

	cd "basable/internal/domain/collaborator"
	_ "basable/internal/infrastructure/database/entities"
	ex "basable/pkg/exceptions"
	rs "basable/pkg/responses"
	cm "basable/shared/usecase"
)

type CollaboratorHandler interface {
	GetAllCollaborators(ctx *gin.Context)
	GetCollaboratorByID(ctx *gin.Context)
	CreateCollaborator(ctx *gin.Context)
	UpdateCollaborator(ctx *gin.Context)
	DeleteCollaborator(ctx *gin.Context)
}

type collaboratorHandler struct {
	collaboratorService cd.CollaboratorService
}

func NewCollaboratorHandler(collaboratorService cd.CollaboratorService) *collaboratorHandler {
	return &collaboratorHandler{
		collaboratorService: collaboratorService,
	}
}

// GetAllCollaborators godoc
// @Summary      returns all collaborators
// @Description  get all collaborators
// @Tags         Collaborator
// @Accept       json
// @Produce      json
// @Param        page  			query int  		false  "page number" default(1)
// @Param        size  			query int  		false  "page size" default(10)
// @Param		fullName 		query string 	false "fullName"
// @Param		email 			query string 	false "email"
// @Param		role 			query string 	false "role"
// @success 200 {object} cm.SuccessResponseDto{data=[]entities.Collaborator}	"all Collaborators returned"
// @Failure      500  {object}  cm.FailedResponseDto	"unexpected internal server error"
// @Router       /v1/collaborators [get]
func (ch *collaboratorHandler) GetAllCollaborators(ctx *gin.Context) {
	query := cd.GetAllCollaboratorsQueryFromCtx(ctx)

	allCollaborators, err := ch.collaboratorService.GetAllCollaborators(query)
	if err != nil {
		ex.HandleInternalServerException(ctx)
		return
	}

	paginateResponse := cm.PaginateResponseDto{
		Content: allCollaborators,
		Page:    query.Page,
		Size:    query.Size,
		Total:   len(allCollaborators),
	}

	rs.HandleOkDataResponse(ctx, "all Collaborators", paginateResponse)
}

// GetCollaboratorByID godoc
// @Summary      returns a collaborator by its 16 chaarcter uuid
// @Description  get collaborator by ID
// @Tags         Collaborator
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "collaborator ID(UUID)"
// @success 	200 {object} cm.SuccessResponseDto{data=entities.Collaborator} "desc"
// @Failure      400  {object}  cm.FailedResponseDto	"request param validation error"
// @Failure      404  {object}  cm.FailedResponseDto	"collaborator with the specified ID not found"
// @Failure      500  {object}  cm.FailedResponseDto	"unexpected internal server error"
// @Router       /v1/collaborators/{id} [get]
func (ch *collaboratorHandler) GetCollaboratorByID(ctx *gin.Context) {
	params := cm.EntityID{}
	if err := ctx.BindUri(&params); err != nil {
		ex.HandleValidationException(ctx, err)
		return
	}

	collaborator, err := ch.collaboratorService.GetByID(params.ID)
	if err != nil {
		ex.HandleNotFoundException(ctx, err)
		return
	}
	rs.HandleOkDataResponse(ctx, "Collaborator with ID: "+collaborator.ID.String()+"", collaborator)
}

// CreateCollaborator godoc
// @Summary      registers a new collaborator
// @Description  create new collaborator
// @Tags         Collaborator
// @Accept       json
// @Produce      json
// @Param 		 data	body	cd.CreateCollaboratorRequest	true	"New Collaborator Details JSON"
// @Success      201  {object}  cm.SuccessResponseDto{data=entities.Collaborator}	"Collaborator created successfully"
// @Failure      400  {object}  cm.FailedResponseDto "request body validation error"
// @Failure      409  {object}  cm.FailedResponseDto "another Collaborator with supplied email exists"
// @Failure      500  {object}  cm.FailedResponseDto "unexpected internal server error"
// @Router       /v1/collaborators [post]
func (ch *collaboratorHandler) CreateCollaborator(ctx *gin.Context) {
	var body cd.CreateCollaboratorRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ex.HandleValidationException(ctx, err)
		return
	}

	collaborator, err := ch.collaboratorService.CreateCollaborator(&body)
	if err != nil {
		switch statusCode := err.StatusCode; statusCode {
		case 400:
			ex.HandleBadRequestException(ctx, err.Error)
			return
		case 409:
			ex.HandleConflictException(ctx, err.Error.Error())
			return
		default:
			return
		}
	}
	rs.HandleOkDataResponse(ctx, "Collaborator created successfully", collaborator)
}

// UpdateCollaborator godoc
// @Summary      updates a collaborator
// @Description  update collaborator
// @Tags         Collaborator
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Collaborator ID(UUID)"
// @Param 		 data	body	cd.UpdateCollaboratorRequest	true	"Collaborator Details JSON"
// @success 200 {object} cm.SuccessResponseDto{data=entities.Collaborator}	"Collaborator updated successfully"
// @Failure      400  {object}  cm.FailedResponseDto	"request body/param validation error"
// @Failure      404  {object}  cm.FailedResponseDto	"Collaborator with specified ID not found"
// @Failure      500  {object}  cm.FailedResponseDto	"unexpected internal server error"
// @Router       /v1/collaborators/{id} [patch]
func (ch *collaboratorHandler) UpdateCollaborator(ctx *gin.Context) {
	params := cm.EntityID{}
	if err := ctx.BindUri(&params); err != nil {
		ex.HandleValidationException(ctx, err)
		return
	}

	body := cd.UpdateCollaboratorRequest{}
	if err := ctx.BindJSON(&body); err != nil {
		ex.HandleValidationException(ctx, err)
		return
	}

	collaborator, err := ch.collaboratorService.UpdateCollaborator(params.ID, &body)
	if err != nil {
		switch statusCode := err.StatusCode; statusCode {
		case 400:
			ex.HandleNotFoundException(ctx, err.Error)
			return
		case 500:
			ex.HandleInternalServerException(ctx)
			return
		default:
			return
		}
	}
	rs.HandleOkDataResponse(ctx, "Collaborator updated successfully", collaborator)
}

// DeleteCollaborator godoc
// @Summary      deletes a collaborator
// @Description  delete collaborator
// @Tags         Collaborator
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Collaborator ID(UUID)"
// @success 200 {object} cm.SuccessResponseDto	"Collaborator deleted suuceesfully"
// @Failure      400  {object}  cm.FailedResponseDto	"request param validation error"
// @Failure      500  {object}  cm.FailedResponseDto	"unexpected internal server error"
// @Router       /v1/collaborators/{id} [delete]
func (ch *collaboratorHandler) DeleteCollaborator(ctx *gin.Context) {
	params := cm.EntityID{}
	if err := ctx.BindUri(&params); err != nil {
		ex.HandleValidationException(ctx, err)
		return
	}

	err := ch.collaboratorService.DeleteCollaborator(params.ID)
	if err != nil {
		ex.HandleNotFoundException(ctx, err.Error)
		return
	}
	rs.HandleOkDataResponse(ctx, "Collaborator deleted successfully", nil)
}
