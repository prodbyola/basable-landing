package collaborator

import (
	"log"
	"net/http"

	ent "basable/internal/infrastructure/database/entities"
	ex "basable/pkg/exceptions"
)

type CollaboratorService interface {
	GetAllCollaborators(query *GetAllCollaboratorsQuery) ([]*ent.Collaborator, error)
	GetByID(id string) (*ent.Collaborator, error)
	CreateCollaborator(request *CreateCollaboratorRequest) (*ent.Collaborator, *ex.Exception)
	UpdateCollaborator(id string, request *UpdateCollaboratorRequest) (*ent.Collaborator, *ex.Exception)
	DeleteCollaborator(id string) *ex.Exception
}

type collaboratorService struct {
	collaboratorRepository CollaboratorRepository
}

func NewCollaboratorService(collaboratorRepository CollaboratorRepository) *collaboratorService {
	return &collaboratorService{
		collaboratorRepository: collaboratorRepository,
	}
}

func (cs *collaboratorService) GetAllCollaborators(query *GetAllCollaboratorsQuery) ([]*ent.Collaborator, error) {
	allColaborators, err := cs.collaboratorRepository.GetCollaborators(query)
	if err != nil {
		return nil, err
	}

	return allColaborators, nil
}

func (cs *collaboratorService) GetByID(id string) (*ent.Collaborator, error) {
	collaborator, err := cs.collaboratorRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return collaborator, nil
}

func (cs *collaboratorService) CreateCollaborator(request *CreateCollaboratorRequest) (*ent.Collaborator, *ex.Exception) {
	collaborator := &ent.Collaborator{
		FullName: request.FullName,
		Email:    request.Email,
		Role:     ent.Role(request.Role),
	}

	if request.ImageUrl != nil {
		collaborator.ImageUrl = request.ImageUrl
	}
	if request.Linkedin != nil {
		collaborator.Linkedin = request.Linkedin
	}
	if request.Github != nil {
		collaborator.Github = request.Github
	}

	if err := cs.collaboratorRepository.CreateCollaborator(collaborator); err != nil {
		return nil, ex.NewException(http.StatusConflict, err)
	}

	return collaborator, nil
}

func (cs *collaboratorService) UpdateCollaborator(id string, request *UpdateCollaboratorRequest) (*ent.Collaborator, *ex.Exception) {
	if _, err := cs.collaboratorRepository.FindByID(id); err != nil {
		return nil, ex.NewException(http.StatusNotFound, err)
	}

	collaboratorUpdate, fields := request.FormatCollaboratorUpdate()
	updatedCollaborator, err := cs.collaboratorRepository.UpdateCollaborator(id, fields, collaboratorUpdate)
	if err != nil {
		log.Println(err)
		return nil, ex.NewException(http.StatusInternalServerError, err)
	}

	return updatedCollaborator, nil
}

func (cs *collaboratorService) DeleteCollaborator(id string) *ex.Exception {
	if _, err := cs.collaboratorRepository.FindByID(id); err != nil {
		return ex.NewException(http.StatusNotFound, err)
	}

	if err := cs.collaboratorRepository.DeleteCollaborator(id); err != nil {
		return ex.NewException(http.StatusInternalServerError, err)
	}

	return nil
}
