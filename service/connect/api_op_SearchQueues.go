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

// This API is in preview release for Amazon Connect and is subject to change.
// Searches queues in an Amazon Connect instance, with optional filtering.
func (c *Client) SearchQueues(ctx context.Context, params *SearchQueuesInput, optFns ...func(*Options)) (*SearchQueuesOutput, error) {
	if params == nil {
		params = &SearchQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchQueues", params, optFns, c.addOperationSearchQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchQueuesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The search criteria to be used to return queues.
	SearchCriteria *types.QueueSearchCriteria

	// Filters to be applied to search results.
	SearchFilter *types.QueueSearchFilter

	noSmithyDocumentSerde
}

type SearchQueuesOutput struct {

	// The total number of queues which matched your search query.
	ApproximateTotalCount *int64

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the queues.
	Queues []types.Queue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchQueues{}, middleware.After)
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
	if err = addOpSearchQueuesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchQueues(options.Region), middleware.Before); err != nil {
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

// SearchQueuesAPIClient is a client that implements the SearchQueues operation.
type SearchQueuesAPIClient interface {
	SearchQueues(context.Context, *SearchQueuesInput, ...func(*Options)) (*SearchQueuesOutput, error)
}

var _ SearchQueuesAPIClient = (*Client)(nil)

// SearchQueuesPaginatorOptions is the paginator options for SearchQueues
type SearchQueuesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchQueuesPaginator is a paginator for SearchQueues
type SearchQueuesPaginator struct {
	options   SearchQueuesPaginatorOptions
	client    SearchQueuesAPIClient
	params    *SearchQueuesInput
	nextToken *string
	firstPage bool
}

// NewSearchQueuesPaginator returns a new SearchQueuesPaginator
func NewSearchQueuesPaginator(client SearchQueuesAPIClient, params *SearchQueuesInput, optFns ...func(*SearchQueuesPaginatorOptions)) *SearchQueuesPaginator {
	if params == nil {
		params = &SearchQueuesInput{}
	}

	options := SearchQueuesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchQueuesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchQueuesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchQueues page.
func (p *SearchQueuesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchQueuesOutput, error) {
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

	result, err := p.client.SearchQueues(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchQueues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "SearchQueues",
	}
}
