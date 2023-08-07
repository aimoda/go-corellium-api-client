/*
Corellium API

Testing LicensingApiService

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

func Test_corellium_LicensingApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LicensingApiService V1GetSupportedFeatures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LicensingApi.V1GetSupportedFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
