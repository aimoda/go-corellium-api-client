/*
Corellium API

Testing WebPlayerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package corellium

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/aimoda/go-corellium-api-client"
)

func Test_corellium_WebPlayerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebPlayerApiService V1WebPlayerAllowedDomains", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WebPlayerApi.V1WebPlayerAllowedDomains(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebPlayerApiService V1WebPlayerCreateSession", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WebPlayerApi.V1WebPlayerCreateSession(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebPlayerApiService V1WebPlayerDestroySession", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		httpRes, err := apiClient.WebPlayerApi.V1WebPlayerDestroySession(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebPlayerApiService V1WebPlayerListSessions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WebPlayerApi.V1WebPlayerListSessions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebPlayerApiService V1WebPlayerSessionInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.WebPlayerApi.V1WebPlayerSessionInfo(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}