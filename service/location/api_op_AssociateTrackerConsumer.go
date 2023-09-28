// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an association between a geofence collection and a tracker resource.
// This allows the tracker resource to communicate location data to the linked
// geofence collection. You can associate up to five geofence collections to each
// tracker resource. Currently not supported — Cross-account configurations, such
// as creating associations between a tracker resource in one account and a
// geofence collection in another account.
func (c *Client) AssociateTrackerConsumer(ctx context.Context, params *AssociateTrackerConsumerInput, optFns ...func(*Options)) (*AssociateTrackerConsumerOutput, error) {
	if params == nil {
		params = &AssociateTrackerConsumerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateTrackerConsumer", params, optFns, c.addOperationAssociateTrackerConsumerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateTrackerConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTrackerConsumerInput struct {

	// The Amazon Resource Name (ARN) for the geofence collection to be associated to
	// tracker resource. Used when you need to specify a resource across all oss.
	//
	// *
	// Format example:
	// arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollectionConsumer
	//
	// This member is required.
	ConsumerArn *string

	// The name of the tracker resource to be associated with a geofence collection.
	//
	// This member is required.
	TrackerName *string

	noSmithyDocumentSerde
}

type AssociateTrackerConsumerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateTrackerConsumerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateTrackerConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateTrackerConsumer{}, middleware.After)
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
	if err = addEndpointPrefix_opAssociateTrackerConsumerMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateTrackerConsumerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTrackerConsumer(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opAssociateTrackerConsumerMiddleware struct {
}

func (*endpointPrefix_opAssociateTrackerConsumerMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opAssociateTrackerConsumerMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opAssociateTrackerConsumerMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opAssociateTrackerConsumerMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateTrackerConsumer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "AssociateTrackerConsumer",
	}
}