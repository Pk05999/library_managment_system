package users

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email            string `json:"email"`
	Password         string `json:"password"`
	Status           string `json:"status"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	OrganizationName string `json:"organization_name"`
}
