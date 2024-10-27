package collaborator

import (
	"strconv"

	"github.com/gin-gonic/gin"

	ent "basable/internal/infrastructure/database/entities"
)

type CreateCollaboratorRequest struct {
	FullName string  `json:"fullName" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Role     string  `json:"role" default:"user"`
	ImageUrl *string `json:"imageUrl,"`
	Linkedin *string `json:"linkedin"`
	Github   *string `json:"github"`
} // @name CreateCollaboratorRequest

type UpdateCollaboratorRequest struct {
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	ImageUrl string `json:"imageUrl"`
	Linkedin string `json:"linkedin"`
	Github   string `json:"github"`
} // @name UpdateCollaborator

func (ucr *UpdateCollaboratorRequest) FormatCollaboratorUpdate() (*ent.Collaborator, []string) {
	updatedCollaborator := &ent.Collaborator{}
	fields := []string{}

	switch {
	case len(ucr.FullName) > 0:
		updatedCollaborator.FullName = ucr.FullName
		fields = append(fields, "full_name")
	case len(ucr.Role) > 0:
		updatedCollaborator.Role = ent.GetRoleFromString(ucr.Role)
		fields = append(fields, "role")
	case len(ucr.ImageUrl) > 0:
		updatedCollaborator.ImageUrl = &ucr.ImageUrl
		fields = append(fields, "image_url")
	case len(ucr.Linkedin) > 0:
		updatedCollaborator.Linkedin = &ucr.Linkedin
		fields = append(fields, "linkedin")
	case len(ucr.Github) > 0:
		updatedCollaborator.Github = &ucr.Github
		fields = append(fields, "github")
	}

	return updatedCollaborator, fields
}

type GetAllCollaboratorsQuery struct {
	Page     int    `form:"page" default:"1"`
	Size     int    `form:"size" default:"10"`
	FullName string `form:"fullName"`
	Email    string `form:"email"`
	Role     string `form:"role"`
} // @name GetAllCollaboratorsQuery

func GetAllCollaboratorsQueryFromCtx(ctx *gin.Context) *GetAllCollaboratorsQuery {
	pageStr := ctx.DefaultQuery("page", "1")
	sizeStr := ctx.DefaultQuery("size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	return &GetAllCollaboratorsQuery{
		Page:     page,
		Size:     size,
		FullName: ctx.Query("fullName"),
		Email:    ctx.Query("email"),
		Role:     ctx.Query("role"),
	}
}
