package interfaces

import (
	lms "github.com/Pk05999/library_managment_system/internal/library_mangment_system"
	// lms "github.com/Pk05999/library_managment_system/internal/library_mangment_system"
	"github.com/gin-gonic/gin"
)

type IControllerUser interface {
	UserSignUp(c *gin.Context) *lms.HelperResponse
}
