/*
Corellium API

Testing AuthenticationApiService

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

func Test_corellium_AuthenticationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuthenticationApiService V1AuthLogin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuthenticationApi.V1AuthLogin(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}