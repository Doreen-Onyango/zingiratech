// Routes initialization with valid project root path returns nil error
package routes

import (
	"errors"
	"net/http"
	"testing"

	"bou.ke/monkey"
	"github.com/Doreen-Onyango/zingiratech/backend/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestInitRoutesWithValidPath(t *testing.T) {
	mux := http.NewServeMux()

	err := InitRoutes(mux)

	assert.NoError(t, err)
	assert.NotNil(t, mux)
}

// Handle invalid or non-existent project root path
func TestInitRoutesWithInvalidPath(t *testing.T) {
	mux := http.NewServeMux()

	// Mock utils.GetProjectRootPath to return error
	monkey.Patch(utils.GetProjectRootPath, func(paths ...string) (string, error) {
		return "", errors.New("invalid path")
	})
	defer monkey.Unpatch(utils.GetProjectRootPath)

	err := InitRoutes(mux)

	assert.Error(t, err)
	assert.Equal(t, "invalid path", err.Error())
}
