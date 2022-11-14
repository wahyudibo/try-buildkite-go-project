package getuserbyid

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	userdao "github.com/wahyudibo/try-buildkite-go-project/internal/application/user/dao"
	postgrespkg "github.com/wahyudibo/try-buildkite-go-project/internal/pkg/postgres"
)

func TestGetUserByIDFeatureSuccess(t *testing.T) {
	ctx := context.Background()

	dbCfg, err := postgrespkg.NewConfig()
	require.NoError(t, err)

	dbConn, err := postgrespkg.NewConnection(ctx, dbCfg)
	require.NoError(t, err)

	userDAO := userdao.New(dbConn)
	handler := New(userDAO)

	app := fiber.New()
	app.Get("/api/users/:userId", handler.Handler)

	expected := userdao.User{
		ID:   1,
		Name: "John Doe",
	}

	url := fmt.Sprintf("/api/users/%d", expected.ID)
	req := httptest.NewRequest(http.MethodGet, url, nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	fmt.Printf("respBody: %+v\n", string(body))

	var user userdao.User
	err = json.Unmarshal(body, &user)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expected.ID, user.ID)
	assert.Equal(t, expected.Name, user.Name)
}
