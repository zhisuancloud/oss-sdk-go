// Code generated by smithy-go-codegen DO NOT EDIT.

// Package mwaa provides the API client, operations, and parameter types for
// AmazonMWAA.
//
// Amazon Managed Workflows for Apache Airflow This section contains the Amazon
// Managed Workflows for Apache Airflow (MWAA) API reference documentation. For
// more information, see What Is Amazon MWAA?
// (https://docs.aws.amazon.com/mwaa/latest/userguide/what-is-mwaa.html).
// Endpoints
//
// * api.airflow.{region}.amazonaws.com - This endpoint is used for
// environment management.
//
// * CreateEnvironment
// (https://docs.aws.amazon.com/mwaa/latest/API/API_CreateEnvironment.html)
//
// *
// DeleteEnvironment
// (https://docs.aws.amazon.com/mwaa/latest/API/API_DeleteEnvironment.html)
//
// *
// GetEnvironment
// (https://docs.aws.amazon.com/mwaa/latest/API/API_GetEnvironment.html)
//
// *
// ListEnvironments
// (https://docs.aws.amazon.com/mwaa/latest/API/API_ListEnvironments.html)
//
// *
// ListTagsForResource
// (https://docs.aws.amazon.com/mwaa/latest/API/API_ListTagsForResource.html)
//
// *
// TagResource
// (https://docs.aws.amazon.com/mwaa/latest/API/API_TagResource.html)
//
// *
// UntagResource
// (https://docs.aws.amazon.com/mwaa/latest/API/API_UntagResource.html)
//
// *
// UpdateEnvironment
// (https://docs.aws.amazon.com/mwaa/latest/API/API_UpdateEnvironment.html)
//
// *
// env.airflow.{region}.amazonaws.com - This endpoint is used to operate the
// Airflow environment.
//
// * CreateCliToken
// (https://docs.aws.amazon.com/mwaa/latest/API/API_CreateCliToken.html)
//
// *
// CreateWebLoginToken
// (https://docs.aws.amazon.com/mwaa/latest/API/API_CreateWebLoginToken.html)
//
// *
// ops.airflow.{region}.amazonaws.com - This endpoint is used to push environment
// metrics that track environment health.
//
// * PublishMetrics
// (https://docs.aws.amazon.com/mwaa/latest/API/API_PublishMetrics.html)
//
// Regions
// For a list of regions that Amazon MWAA supports, see Region availability
// (https://docs.aws.amazon.com/mwaa/latest/userguide/what-is-mwaa.html#regions-mwaa)
// in the Amazon MWAA User Guide.
package mwaa
