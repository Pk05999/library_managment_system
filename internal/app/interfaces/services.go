package interfaces

import (
	"net/url"

	"github.com/Pk05999/library_managment_system/internal/app/models/users"
	lms "github.com/Pk05999/library_managment_system/internal/library_mangment_system"
)

type IServerUsers interface {
	UserSignUp(request users.Users, query url.Values) *lms.HelperResponse
}
