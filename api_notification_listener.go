/*
Party Management

TMF API Reference : TMF 632 - Party  Release: 22.5 The party API provides standardized mechanism for party management such as creation, update, retrieval, deletion, and notification of events. Party can be an individual or an organization that has any kind of relation with the enterprise. Party is created to record individual or organization information before the assignment of any role. For example, within the context of a split billing mechanism, Party API allows creation of the individual or organization that will play the role of 3rd payer for a given offer and, then, allows consultation or update of his information. Resources - Party (abstract base class with concrete subclasses Individual and Organization) Party API performs the following operations: - Retrieve an organization or an individual - Retrieve a collection of organizations or individuals according to given criteria - Create a new organization or a new individual - Update an existing organization or an existing individual - Delete an existing organization or an existing individual - Notify events on organization or individual

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// NotificationListenerAPIService NotificationListenerAPI service
type NotificationListenerAPIService service

type ApiIndividualAttributeValueChangeEventRequest struct {
	ctx context.Context
	ApiService *NotificationListenerAPIService
	individualAttributeValueChangeEvent *IndividualAttributeValueChangeEvent
}

// Individual attributeValueChange Event payload
func (r ApiIndividualAttributeValueChangeEventRequest) IndividualAttributeValueChangeEvent(individualAttributeValueChangeEvent IndividualAttributeValueChangeEvent) ApiIndividualAttributeValueChangeEventRequest {
	r.individualAttributeValueChangeEvent = &individualAttributeValueChangeEvent
	return r
}

func (r ApiIndividualAttributeValueChangeEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.IndividualAttributeValueChangeEventExecute(r)
}

/*
IndividualAttributeValueChangeEvent Client listener for entity IndividualAttributeValueChangeEvent

Example of a client listener for receiving the notification IndividualAttributeValueChangeEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIndividualAttributeValueChangeEventRequest
*/
func (a *NotificationListenerAPIService) IndividualAttributeValueChangeEvent(ctx context.Context) ApiIndividualAttributeValueChangeEventRequest {
	return ApiIndividualAttributeValueChangeEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationListenerAPIService) IndividualAttributeValueChangeEventExecute(r ApiIndividualAttributeValueChangeEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationListenerAPIService.IndividualAttributeValueChangeEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listener/individualAttributeValueChangeEvent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.individualAttributeValueChangeEvent == nil {
		return nil, reportError("individualAttributeValueChangeEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.individualAttributeValueChangeEvent
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIndividualCreateEventRequest struct {
	ctx context.Context
	ApiService *NotificationListenerAPIService
	individualCreateEvent *IndividualCreateEvent
}

// Individual create Event payload
func (r ApiIndividualCreateEventRequest) IndividualCreateEvent(individualCreateEvent IndividualCreateEvent) ApiIndividualCreateEventRequest {
	r.individualCreateEvent = &individualCreateEvent
	return r
}

func (r ApiIndividualCreateEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.IndividualCreateEventExecute(r)
}

/*
IndividualCreateEvent Client listener for entity IndividualCreateEvent

Example of a client listener for receiving the notification IndividualCreateEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIndividualCreateEventRequest
*/
func (a *NotificationListenerAPIService) IndividualCreateEvent(ctx context.Context) ApiIndividualCreateEventRequest {
	return ApiIndividualCreateEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationListenerAPIService) IndividualCreateEventExecute(r ApiIndividualCreateEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationListenerAPIService.IndividualCreateEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listener/individualCreateEvent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.individualCreateEvent == nil {
		return nil, reportError("individualCreateEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.individualCreateEvent
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIndividualDeleteEventRequest struct {
	ctx context.Context
	ApiService *NotificationListenerAPIService
	individualDeleteEvent *IndividualDeleteEvent
}

// Individual delete Event payload
func (r ApiIndividualDeleteEventRequest) IndividualDeleteEvent(individualDeleteEvent IndividualDeleteEvent) ApiIndividualDeleteEventRequest {
	r.individualDeleteEvent = &individualDeleteEvent
	return r
}

func (r ApiIndividualDeleteEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.IndividualDeleteEventExecute(r)
}

/*
IndividualDeleteEvent Client listener for entity IndividualDeleteEvent

Example of a client listener for receiving the notification IndividualDeleteEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIndividualDeleteEventRequest
*/
func (a *NotificationListenerAPIService) IndividualDeleteEvent(ctx context.Context) ApiIndividualDeleteEventRequest {
	return ApiIndividualDeleteEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationListenerAPIService) IndividualDeleteEventExecute(r ApiIndividualDeleteEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationListenerAPIService.IndividualDeleteEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listener/individualDeleteEvent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.individualDeleteEvent == nil {
		return nil, reportError("individualDeleteEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.individualDeleteEvent
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIndividualStateChangeEventRequest struct {
	ctx context.Context
	ApiService *NotificationListenerAPIService
	individualStateChangeEvent *IndividualStateChangeEvent
}

// Individual stateChange Event payload
func (r ApiIndividualStateChangeEventRequest) IndividualStateChangeEvent(individualStateChangeEvent IndividualStateChangeEvent) ApiIndividualStateChangeEventRequest {
	r.individualStateChangeEvent = &individualStateChangeEvent
	return r
}

func (r ApiIndividualStateChangeEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.IndividualStateChangeEventExecute(r)
}

/*
IndividualStateChangeEvent Client listener for entity IndividualStateChangeEvent

Example of a client listener for receiving the notification IndividualStateChangeEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIndividualStateChangeEventRequest
*/
func (a *NotificationListenerAPIService) IndividualStateChangeEvent(ctx context.Context) ApiIndividualStateChangeEventRequest {
	return ApiIndividualStateChangeEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationListenerAPIService) IndividualStateChangeEventExecute(r ApiIndividualStateChangeEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationListenerAPIService.IndividualStateChangeEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listener/individualStateChangeEvent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.individualStateChangeEvent == nil {
		return nil, reportError("individualStateChangeEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.individualStateChangeEvent
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiOrganizationAttributeValueChangeEventRequest struct {
	ctx context.Context
	ApiService *NotificationListenerAPIService
	organizationAttributeValueChangeEvent *OrganizationAttributeValueChangeEvent
}

// Organization attributeValueChange Event payload
func (r ApiOrganizationAttributeValueChangeEventRequest) OrganizationAttributeValueChangeEvent(organizationAttributeValueChangeEvent OrganizationAttributeValueChangeEvent) ApiOrganizationAttributeValueChangeEventRequest {
	r.organizationAttributeValueChangeEvent = &organizationAttributeValueChangeEvent
	return r
}

func (r ApiOrganizationAttributeValueChangeEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrganizationAttributeValueChangeEventExecute(r)
}

/*
OrganizationAttributeValueChangeEvent Client listener for entity OrganizationAttributeValueChangeEvent

Example of a client listener for receiving the notification OrganizationAttributeValueChangeEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrganizationAttributeValueChangeEventRequest
*/
func (a *NotificationListenerAPIService) OrganizationAttributeValueChangeEvent(ctx context.Context) ApiOrganizationAttributeValueChangeEventRequest {
	return ApiOrganizationAttributeValueChangeEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationListenerAPIService) OrganizationAttributeValueChangeEventExecute(r ApiOrganizationAttributeValueChangeEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationListenerAPIService.OrganizationAttributeValueChangeEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listener/organizationAttributeValueChangeEvent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationAttributeValueChangeEvent == nil {
		return nil, reportError("organizationAttributeValueChangeEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.organizationAttributeValueChangeEvent
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiOrganizationCreateEventRequest struct {
	ctx context.Context
	ApiService *NotificationListenerAPIService
	organizationCreateEvent *OrganizationCreateEvent
}

// Organization create Event payload
func (r ApiOrganizationCreateEventRequest) OrganizationCreateEvent(organizationCreateEvent OrganizationCreateEvent) ApiOrganizationCreateEventRequest {
	r.organizationCreateEvent = &organizationCreateEvent
	return r
}

func (r ApiOrganizationCreateEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrganizationCreateEventExecute(r)
}

/*
OrganizationCreateEvent Client listener for entity OrganizationCreateEvent

Example of a client listener for receiving the notification OrganizationCreateEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrganizationCreateEventRequest
*/
func (a *NotificationListenerAPIService) OrganizationCreateEvent(ctx context.Context) ApiOrganizationCreateEventRequest {
	return ApiOrganizationCreateEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationListenerAPIService) OrganizationCreateEventExecute(r ApiOrganizationCreateEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationListenerAPIService.OrganizationCreateEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listener/organizationCreateEvent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationCreateEvent == nil {
		return nil, reportError("organizationCreateEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.organizationCreateEvent
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiOrganizationDeleteEventRequest struct {
	ctx context.Context
	ApiService *NotificationListenerAPIService
	organizationDeleteEvent *OrganizationDeleteEvent
}

// Organization delete Event payload
func (r ApiOrganizationDeleteEventRequest) OrganizationDeleteEvent(organizationDeleteEvent OrganizationDeleteEvent) ApiOrganizationDeleteEventRequest {
	r.organizationDeleteEvent = &organizationDeleteEvent
	return r
}

func (r ApiOrganizationDeleteEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrganizationDeleteEventExecute(r)
}

/*
OrganizationDeleteEvent Client listener for entity OrganizationDeleteEvent

Example of a client listener for receiving the notification OrganizationDeleteEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrganizationDeleteEventRequest
*/
func (a *NotificationListenerAPIService) OrganizationDeleteEvent(ctx context.Context) ApiOrganizationDeleteEventRequest {
	return ApiOrganizationDeleteEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationListenerAPIService) OrganizationDeleteEventExecute(r ApiOrganizationDeleteEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationListenerAPIService.OrganizationDeleteEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listener/organizationDeleteEvent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationDeleteEvent == nil {
		return nil, reportError("organizationDeleteEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.organizationDeleteEvent
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiOrganizationStateChangeEventRequest struct {
	ctx context.Context
	ApiService *NotificationListenerAPIService
	organizationStateChangeEvent *OrganizationStateChangeEvent
}

// Organization stateChange Event payload
func (r ApiOrganizationStateChangeEventRequest) OrganizationStateChangeEvent(organizationStateChangeEvent OrganizationStateChangeEvent) ApiOrganizationStateChangeEventRequest {
	r.organizationStateChangeEvent = &organizationStateChangeEvent
	return r
}

func (r ApiOrganizationStateChangeEventRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrganizationStateChangeEventExecute(r)
}

/*
OrganizationStateChangeEvent Client listener for entity OrganizationStateChangeEvent

Example of a client listener for receiving the notification OrganizationStateChangeEvent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrganizationStateChangeEventRequest
*/
func (a *NotificationListenerAPIService) OrganizationStateChangeEvent(ctx context.Context) ApiOrganizationStateChangeEventRequest {
	return ApiOrganizationStateChangeEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationListenerAPIService) OrganizationStateChangeEventExecute(r ApiOrganizationStateChangeEventRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationListenerAPIService.OrganizationStateChangeEvent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listener/organizationStateChangeEvent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationStateChangeEvent == nil {
		return nil, reportError("organizationStateChangeEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.organizationStateChangeEvent
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
