package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

func (r *Role) GetValue() string {
	return string(*r)
}

func GetRoleFromString(role string) Role {
	switch role {
	case "admin":
		return Admin
	case "user":
		return User
	default:
		return User
	}
}

type Collaborator struct {
	BaseEntity
	FullName string  `json:"fullName"`
	Email    string  `gorm:"not null;uniqueIndex;size:255" json:"email"`
	Role     Role    `gorm:"column:role;type:enum('admin','user')" json:"role"`
	ImageUrl *string `gorm:"unique" json:"imageUrl,omitempty"`
	Linkedin *string `gorm:"unique" json:"linkedin,omitempty"`
	Github   *string `gorm:"unique" json:"github,omitempty"`
} // @name Collaborator

func (c *Collaborator) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
