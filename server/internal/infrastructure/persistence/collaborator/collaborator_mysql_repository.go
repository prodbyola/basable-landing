package collaborator

import (
	"gorm.io/gorm"

	cd "basable/internal/domain/collaborator"
	ent "basable/internal/infrastructure/database/entities"
)

type CollaboratorMySQLRepository struct {
	DB *gorm.DB
}

func NewCollaboratorMySQLRepository(db *gorm.DB) *CollaboratorMySQLRepository {
	return &CollaboratorMySQLRepository{
		DB: db,
	}
}

func (cr *CollaboratorMySQLRepository) GetCollaborators(query *cd.GetAllCollaboratorsQuery) ([]*ent.Collaborator, error) {
	var collaborators []*ent.Collaborator
	offset := (query.Page - 1) * query.Size
	collabiratorQuery := cr.DB.Table("collaborators").
		Order("created_at desc").
		Limit(query.Size).
		Offset(offset)

	switch {
	case query.FullName != "":
		collabiratorQuery = collabiratorQuery.Where("full_name LIKE ?", "%"+query.FullName+"%")
	case query.Email != "":
		collabiratorQuery = collabiratorQuery.Where("email LIKE ?", "%"+query.Email+"%")
	case query.Role != "":
		collabiratorQuery = collabiratorQuery.Where("role = ?", query.Role)
	}
	err := collabiratorQuery.Scan(&collaborators).Error
	if err != nil {
		return nil, err
	}

	return collaborators, nil
}

func (cr *CollaboratorMySQLRepository) FindByID(id string) (*ent.Collaborator, error) {
	var collaborator *ent.Collaborator
	err := cr.DB.Where("id = ?", id).First(&collaborator).Error
	if err != nil {
		return nil, err
	}

	return collaborator, nil
}

func (cr *CollaboratorMySQLRepository) CreateCollaborator(collaborator *ent.Collaborator) error {
	return cr.DB.Create(collaborator).Error
}

func (cr *CollaboratorMySQLRepository) UpdateCollaborator(id string, fields []string, collaboratorUpdate *ent.Collaborator) (*ent.Collaborator, error) {
	var collaborator *ent.Collaborator
	if err := cr.DB.Model(&collaborator).
		Select(fields).
		Where("id = ?", id).
		Updates(&collaboratorUpdate).Error; err != nil {
		return nil, err
	}

	return collaborator, nil
}

func (cr *CollaboratorMySQLRepository) DeleteCollaborator(id string) error {
	err := cr.DB.Unscoped().Delete(&ent.Collaborator{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
