// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a batch of geofences from a geofence collection. This operation deletes
// the resource permanently.
func (c *Client) BatchDeleteGeofence(ctx context.Context, params *BatchDeleteGeofenceInput, optFns ...func(*Options)) (*BatchDeleteGeofenceOutput, error) {
	if params == nil {
		params = &BatchDeleteGeofenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteGeofence", params, optFns, c.addOperationBatchDeleteGeofenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteGeofenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteGeofenceInput struct {

	// The geofence collection storing the geofences to be deleted.
	//
	// This member is required.
	CollectionName *string

	// The batch of geofences to be deleted.
	//
	// This member is required.
	GeofenceIds []string

	noSmithyDocumentSerde
}

type BatchDeleteGeofenceOutput struct {

	// Contains error details for each geofence that failed to delete.
	//
	// This member is required.
	Errors []types.BatchDeleteGeofenceError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteGeofenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDeleteGeofence{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDeleteGeofence{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchDeleteGeofenceMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchDeleteGeofenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteGeofence(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchDeleteGeofenceMiddleware struct {
}

func (*endpointPrefix_opBatchDeleteGeofenceMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchDeleteGeofenceMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "geofencing." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchDeleteGeofenceMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchDeleteGeofenceMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opBatchDeleteGeofence(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "BatchDeleteGeofence",
	}
}
