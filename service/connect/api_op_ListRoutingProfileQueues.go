// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the queues associated with a routing profile.
func (c *Client) ListRoutingProfileQueues(ctx context.Context, params *ListRoutingProfileQueuesInput, optFns ...func(*Options)) (*ListRoutingProfileQueuesOutput, error) {
	if params == nil {
		params = &ListRoutingProfileQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoutingProfileQueues", params, optFns, c.addOperationListRoutingProfileQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoutingProfileQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoutingProfileQueuesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier of the routing profile.
	//
	// This member is required.
	RoutingProfileId *string

	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRoutingProfileQueuesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the routing profiles.
	RoutingProfileQueueConfigSummaryList []types.RoutingProfileQueueConfigSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRoutingProfileQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRoutingProfileQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRoutingProfileQueues{}, middleware.After)
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
	if err = addOpListRoutingProfileQueuesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRoutingProfileQueues(options.Region), middleware.Before); err != nil {
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

// ListRoutingProfileQueuesAPIClient is a client that implements the
// ListRoutingProfileQueues operation.
type ListRoutingProfileQueuesAPIClient interface {
	ListRoutingProfileQueues(context.Context, *ListRoutingProfileQueuesInput, ...func(*Options)) (*ListRoutingProfileQueuesOutput, error)
}

var _ ListRoutingProfileQueuesAPIClient = (*Client)(nil)

// ListRoutingProfileQueuesPaginatorOptions is the paginator options for
// ListRoutingProfileQueues
type ListRoutingProfileQueuesPaginatorOptions struct {
	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRoutingProfileQueuesPaginator is a paginator for ListRoutingProfileQueues
type ListRoutingProfileQueuesPaginator struct {
	options   ListRoutingProfileQueuesPaginatorOptions
	client    ListRoutingProfileQueuesAPIClient
	params    *ListRoutingProfileQueuesInput
	nextToken *string
	firstPage bool
}

// NewListRoutingProfileQueuesPaginator returns a new
// ListRoutingProfileQueuesPaginator
func NewListRoutingProfileQueuesPaginator(client ListRoutingProfileQueuesAPIClient, params *ListRoutingProfileQueuesInput, optFns ...func(*ListRoutingProfileQueuesPaginatorOptions)) *ListRoutingProfileQueuesPaginator {
	if params == nil {
		params = &ListRoutingProfileQueuesInput{}
	}

	options := ListRoutingProfileQueuesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRoutingProfileQueuesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRoutingProfileQueuesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRoutingProfileQueues page.
func (p *ListRoutingProfileQueuesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRoutingProfileQueuesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListRoutingProfileQueues(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRoutingProfileQueues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListRoutingProfileQueues",
	}
}
