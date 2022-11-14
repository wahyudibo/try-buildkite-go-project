package getuserbyid

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	userdao "github.com/wahyudibo/try-buildkite-go-project/internal/application/user/dao"
	"github.com/wahyudibo/try-buildkite-go-project/internal/config"
)

func TestGetUserByIDFeatureSuccess(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	expected := userdao.User{
		ID:   1,
		Name: "John Doe",
	}

	url := fmt.Sprintf("http://localhost:%d/api/users/%d", cfg.HTTPServerPort, expected.ID)
	resp, err := http.Get(url)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var user userdao.User
	err = json.Unmarshal(body, &user)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expected.ID, user.ID)
	assert.Equal(t, expected.Name, user.Name)
}
