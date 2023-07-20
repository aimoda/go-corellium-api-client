/*
Corellium API

REST API to manage your virtual devices.

API version: 5.2.0-17376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// ImagesApiService ImagesApi service
type ImagesApiService service

type ImagesApiV1CreateImageRequest struct {
	ctx context.Context
	ApiService *ImagesApiService
	type_ *string
	encoding *string
	encapsulated *bool
	name *string
	project *string
	instance *string
	file *os.File
}

// Image type
func (r ImagesApiV1CreateImageRequest) Type_(type_ string) ImagesApiV1CreateImageRequest {
	r.type_ = &type_
	return r
}

// How the file is stored
func (r ImagesApiV1CreateImageRequest) Encoding(encoding string) ImagesApiV1CreateImageRequest {
	r.encoding = &encoding
	return r
}

// set to false if the uploaded file is not encapsulated inside an outer zipfile
func (r ImagesApiV1CreateImageRequest) Encapsulated(encapsulated bool) ImagesApiV1CreateImageRequest {
	r.encapsulated = &encapsulated
	return r
}

// Image name
func (r ImagesApiV1CreateImageRequest) Name(name string) ImagesApiV1CreateImageRequest {
	r.name = &name
	return r
}

// Project ID
func (r ImagesApiV1CreateImageRequest) Project(project string) ImagesApiV1CreateImageRequest {
	r.project = &project
	return r
}

// Instance ID
func (r ImagesApiV1CreateImageRequest) Instance(instance string) ImagesApiV1CreateImageRequest {
	r.instance = &instance
	return r
}

// Optionally the actual file
func (r ImagesApiV1CreateImageRequest) File(file *os.File) ImagesApiV1CreateImageRequest {
	r.file = file
	return r
}

func (r ImagesApiV1CreateImageRequest) Execute() (*Image, *http.Response, error) {
	return r.ApiService.V1CreateImageExecute(r)
}

/*
V1CreateImage Create a new Image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ImagesApiV1CreateImageRequest
*/
func (a *ImagesApiService) V1CreateImage(ctx context.Context) ImagesApiV1CreateImageRequest {
	return ImagesApiV1CreateImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Image
func (a *ImagesApiService) V1CreateImageExecute(r ImagesApiV1CreateImageRequest) (*Image, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Image
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.V1CreateImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/images"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
	}
	if r.encoding == nil {
		return localVarReturnValue, nil, reportError("encoding is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "type", r.type_, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "encoding", r.encoding, "")
	if r.encapsulated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "encapsulated", r.encapsulated, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "")
	}
	if r.project != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "project", r.project, "")
	}
	if r.instance != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "instance", r.instance, "")
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ImagesApiV1DeleteImageRequest struct {
	ctx context.Context
	ApiService *ImagesApiService
	imageId string
}

func (r ImagesApiV1DeleteImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1DeleteImageExecute(r)
}

/*
V1DeleteImage Delete Image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageId Image ID - uuid
 @return ImagesApiV1DeleteImageRequest
*/
func (a *ImagesApiService) V1DeleteImage(ctx context.Context, imageId string) ImagesApiV1DeleteImageRequest {
	return ImagesApiV1DeleteImageRequest{
		ApiService: a,
		ctx: ctx,
		imageId: imageId,
	}
}

// Execute executes the request
func (a *ImagesApiService) V1DeleteImageExecute(r ImagesApiV1DeleteImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.V1DeleteImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ImagesApiV1GetImageRequest struct {
	ctx context.Context
	ApiService *ImagesApiService
	imageId string
}

func (r ImagesApiV1GetImageRequest) Execute() (*Image, *http.Response, error) {
	return r.ApiService.V1GetImageExecute(r)
}

/*
V1GetImage Get Image Metadata

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageId Image ID - uuid
 @return ImagesApiV1GetImageRequest
*/
func (a *ImagesApiService) V1GetImage(ctx context.Context, imageId string) ImagesApiV1GetImageRequest {
	return ImagesApiV1GetImageRequest{
		ApiService: a,
		ctx: ctx,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return Image
func (a *ImagesApiService) V1GetImageExecute(r ImagesApiV1GetImageRequest) (*Image, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Image
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.V1GetImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ImagesApiV1GetImagesRequest struct {
	ctx context.Context
	ApiService *ImagesApiService
	project *string
}

// Optionally filter by project - uuid
func (r ImagesApiV1GetImagesRequest) Project(project string) ImagesApiV1GetImagesRequest {
	r.project = &project
	return r
}

func (r ImagesApiV1GetImagesRequest) Execute() ([]Image, *http.Response, error) {
	return r.ApiService.V1GetImagesExecute(r)
}

/*
V1GetImages Get all Images Metadata

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ImagesApiV1GetImagesRequest
*/
func (a *ImagesApiService) V1GetImages(ctx context.Context) ImagesApiV1GetImagesRequest {
	return ImagesApiV1GetImagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Image
func (a *ImagesApiService) V1GetImagesExecute(r ImagesApiV1GetImagesRequest) ([]Image, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Image
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.V1GetImages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/images"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.project != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "project", r.project, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ImagesApiV1UploadImageDataRequest struct {
	ctx context.Context
	ApiService *ImagesApiService
	imageId string
	body *string
}

// Uploaded Image
func (r ImagesApiV1UploadImageDataRequest) Body(body string) ImagesApiV1UploadImageDataRequest {
	r.body = &body
	return r
}

func (r ImagesApiV1UploadImageDataRequest) Execute() (*Image, *http.Response, error) {
	return r.ApiService.V1UploadImageDataExecute(r)
}

/*
V1UploadImageData Upload Image Data

If the active project has enough remaining quota, updates an Image with the contents of the request body.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageId Image ID - uuid
 @return ImagesApiV1UploadImageDataRequest
*/
func (a *ImagesApiService) V1UploadImageData(ctx context.Context, imageId string) ImagesApiV1UploadImageDataRequest {
	return ImagesApiV1UploadImageDataRequest{
		ApiService: a,
		ctx: ctx,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return Image
func (a *ImagesApiService) V1UploadImageDataExecute(r ImagesApiV1UploadImageDataRequest) (*Image, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Image
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.V1UploadImageData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"binary"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiConflictError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
