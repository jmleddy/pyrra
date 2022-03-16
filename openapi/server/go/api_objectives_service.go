/*
 * Pyrra
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// ObjectivesApiService is a service that implents the logic for the ObjectivesApiServicer
// This service should implement the business logic for every endpoint for the ObjectivesApi API.
// Include any external packages or services that will be required by this service.
type ObjectivesApiService struct {
}

// NewObjectivesApiService creates a default api service
func NewObjectivesApiService() ObjectivesApiServicer {
	return &ObjectivesApiService{}
}

// GetMultiBurnrateAlerts - Get the MultiBurnrateAlerts for the Objective
func (s *ObjectivesApiService) GetMultiBurnrateAlerts(ctx context.Context, expr string, grouping string, inactive bool) (ImplResponse, error) {
	// TODO - update GetMultiBurnrateAlerts with the required logic for this service method.
	// Add api_objectives_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []MultiBurnrateAlert{}) or use other options such as http.Ok ...
	//return Response(200, []MultiBurnrateAlert{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetMultiBurnrateAlerts method not implemented")
}

// GetObjectiveErrorBudget - Get ErrorBudget graph sample pairs
func (s *ObjectivesApiService) GetObjectiveErrorBudget(ctx context.Context, expr string, grouping string, start int32, end int32) (ImplResponse, error) {
	// TODO - update GetObjectiveErrorBudget with the required logic for this service method.
	// Add api_objectives_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, QueryRange{}) or use other options such as http.Ok ...
	//return Response(200, QueryRange{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetObjectiveErrorBudget method not implemented")
}

// GetObjectiveStatus - Get objective status
func (s *ObjectivesApiService) GetObjectiveStatus(ctx context.Context, expr string, grouping string) (ImplResponse, error) {
	// TODO - update GetObjectiveStatus with the required logic for this service method.
	// Add api_objectives_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []ObjectiveStatus{}) or use other options such as http.Ok ...
	//return Response(200, []ObjectiveStatus{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetObjectiveStatus method not implemented")
}

// GetREDErrors - Get a matrix of error percentage by label
func (s *ObjectivesApiService) GetREDErrors(ctx context.Context, expr string, grouping string, start int32, end int32) (ImplResponse, error) {
	// TODO - update GetREDErrors with the required logic for this service method.
	// Add api_objectives_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, QueryRange{}) or use other options such as http.Ok ...
	//return Response(200, QueryRange{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetREDErrors method not implemented")
}

// GetREDRequests - Get a matrix of requests by label
func (s *ObjectivesApiService) GetREDRequests(ctx context.Context, expr string, grouping string, start int32, end int32) (ImplResponse, error) {
	// TODO - update GetREDRequests with the required logic for this service method.
	// Add api_objectives_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, QueryRange{}) or use other options such as http.Ok ...
	//return Response(200, QueryRange{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetREDRequests method not implemented")
}

// ListObjectives - List Objectives
func (s *ObjectivesApiService) ListObjectives(ctx context.Context, expr string) (ImplResponse, error) {
	// TODO - update ListObjectives with the required logic for this service method.
	// Add api_objectives_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Objective{}) or use other options such as http.Ok ...
	//return Response(200, []Objective{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListObjectives method not implemented")
}
