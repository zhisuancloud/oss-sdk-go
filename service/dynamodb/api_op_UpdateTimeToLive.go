// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/dynamodb/types"
	internalEndpointDiscovery "oss-sdk-go/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The UpdateTimeToLive method enables or disables Time to Live (TTL) for the
// specified table. A successful UpdateTimeToLive call returns the current
// TimeToLiveSpecification. It can take up to one hour for the change to fully
// process. Any additional UpdateTimeToLive calls for the same table during this
// one hour duration result in a ValidationException. TTL compares the current time
// in epoch time format to the time stored in the TTL attribute of an item. If the
// epoch time value stored in the attribute is less than the current time, the item
// is marked as expired and subsequently deleted. The epoch time format is the
// number of seconds elapsed since 12:00:00 AM January 1, 1970 UTC. DynamoDB
// deletes expired items on a best-effort basis to ensure availability of
// throughput for other data operations. DynamoDB typically deletes expired items
// within two days of expiration. The exact duration within which an item gets
// deleted after expiration is specific to the nature of the workload. Items that
// have expired and not been deleted will still show up in reads, queries, and
// scans. As items are deleted, they are removed from any local secondary index and
// global secondary index immediately in the same eventually consistent way as a
// standard delete operation. For more information, see Time To Live
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/TTL.html) in
// the Amazon DynamoDB Developer Guide.
func (c *Client) UpdateTimeToLive(ctx context.Context, params *UpdateTimeToLiveInput, optFns ...func(*Options)) (*UpdateTimeToLiveOutput, error) {
	if params == nil {
		params = &UpdateTimeToLiveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTimeToLive", params, optFns, c.addOperationUpdateTimeToLiveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTimeToLiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an UpdateTimeToLive operation.
type UpdateTimeToLiveInput struct {

	// The name of the table to be configured.
	//
	// This member is required.
	TableName *string

	// Represents the settings used to enable or disable Time to Live for the specified
	// table.
	//
	// This member is required.
	TimeToLiveSpecification *types.TimeToLiveSpecification

	noSmithyDocumentSerde
}

type UpdateTimeToLiveOutput struct {

	// Represents the output of an UpdateTimeToLive operation.
	TimeToLiveSpecification *types.TimeToLiveSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTimeToLiveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTimeToLive{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTimeToLive{}, middleware.After)
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
	if err = addOpUpdateTimeToLiveDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addOpUpdateTimeToLiveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTimeToLive(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func addOpUpdateTimeToLiveDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpUpdateTimeToLiveDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpUpdateTimeToLiveDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*UpdateTimeToLiveInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opUpdateTimeToLive(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "UpdateTimeToLive",
	}
}
