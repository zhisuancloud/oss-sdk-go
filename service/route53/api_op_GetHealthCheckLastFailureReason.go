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

// Gets the reason that a specified health check failed most recently.
func (c *Client) GetHealthCheckLastFailureReason(ctx context.Context, params *GetHealthCheckLastFailureReasonInput, optFns ...func(*Options)) (*GetHealthCheckLastFailureReasonOutput, error) {
	if params == nil {
		params = &GetHealthCheckLastFailureReasonInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHealthCheckLastFailureReason", params, optFns, c.addOperationGetHealthCheckLastFailureReasonMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHealthCheckLastFailureReasonOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for the reason that a health check failed most recently.
type GetHealthCheckLastFailureReasonInput struct {

	// The ID for the health check for which you want the last failure reason. When you
	// created the health check, CreateHealthCheck returned the ID in the response, in
	// the HealthCheckId element. If you want to get the last failure reason for a
	// calculated health check, you must use the Amazon Route 53 console or the
	// CloudWatch console. You can't use GetHealthCheckLastFailureReason for a
	// calculated health check.
	//
	// This member is required.
	HealthCheckId *string

	noSmithyDocumentSerde
}

// A complex type that contains the response to a GetHealthCheckLastFailureReason
// request.
type GetHealthCheckLastFailureReasonOutput struct {

	// A list that contains one Observation element for each Amazon Route 53 health
	// checker that is reporting a last failure reason.
	//
	// This member is required.
	HealthCheckObservations []types.HealthCheckObservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHealthCheckLastFailureReasonMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetHealthCheckLastFailureReason{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetHealthCheckLastFailureReason{}, middleware.After)
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
	if err = addOpGetHealthCheckLastFailureReasonValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHealthCheckLastFailureReason(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetHealthCheckLastFailureReason(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHealthCheckLastFailureReason",
	}
}
