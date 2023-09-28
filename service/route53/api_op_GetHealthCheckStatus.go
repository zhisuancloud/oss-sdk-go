// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets status of a specified health check. This API is intended for use during
// development to diagnose behavior. It doesn’t support production use-cases with
// high query rates that require immediate and actionable responses.
func (c *Client) GetHealthCheckStatus(ctx context.Context, params *GetHealthCheckStatusInput, optFns ...func(*Options)) (*GetHealthCheckStatusOutput, error) {
	if params == nil {
		params = &GetHealthCheckStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHealthCheckStatus", params, optFns, c.addOperationGetHealthCheckStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHealthCheckStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get the status for a health check.
type GetHealthCheckStatusInput struct {

	// The ID for the health check that you want the current status for. When you
	// created the health check, CreateHealthCheck returned the ID in the response, in
	// the HealthCheckId element. If you want to check the status of a calculated
	// health check, you must use the Amazon Route 53 console or the CloudWatch
	// console. You can't use GetHealthCheckStatus to get the status of a calculated
	// health check.
	//
	// This member is required.
	HealthCheckId *string

	noSmithyDocumentSerde
}

// A complex type that contains the response to a GetHealthCheck request.
type GetHealthCheckStatusOutput struct {

	// A list that contains one HealthCheckObservation element for each Amazon Route 53
	// health checker that is reporting a status about the health check endpoint.
	//
	// This member is required.
	HealthCheckObservations []types.HealthCheckObservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHealthCheckStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetHealthCheckStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetHealthCheckStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetHealthCheckStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHealthCheckStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetHealthCheckStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHealthCheckStatus",
	}
}
