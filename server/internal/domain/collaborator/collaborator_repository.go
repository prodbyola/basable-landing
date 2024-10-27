package collaborator

import (
	ent "basable/internal/infrastructure/database/entities"
)

type CollaboratorRepository interface {
	GetCollaborators(query *GetAllCollaboratorsQuery) ([]*ent.Collaborator, error)
	FindByID(id string) (*ent.Collaborator, error)
	CreateCollaborator(collaborator *ent.Collaborator) error
	UpdateCollaborator(id string, fields []string, collaboratorUpdate *ent.Collaborator) (*ent.Collaborator, error)
	DeleteCollaborator(id string) error
}
