// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of executed tasks.
func (c *Client) ListTaskExecutions(ctx context.Context, params *ListTaskExecutionsInput, optFns ...func(*Options)) (*ListTaskExecutionsOutput, error) {
	if params == nil {
		params = &ListTaskExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTaskExecutions", params, optFns, c.addOperationListTaskExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTaskExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListTaskExecutions
type ListTaskExecutionsInput struct {

	// The maximum number of executed tasks to list.
	MaxResults *int32

	// An opaque string that indicates the position at which to begin the next list of
	// the executed tasks.
	NextToken *string

	// The Amazon Resource Name (ARN) of the task whose tasks you want to list.
	TaskArn *string

	noSmithyDocumentSerde
}

// ListTaskExecutionsResponse
type ListTaskExecutionsOutput struct {

	// An opaque string that indicates the position at which to begin returning the
	// next list of executed tasks.
	NextToken *string

	// A list of executed tasks.
	TaskExecutions []types.TaskExecutionListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTaskExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTaskExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTaskExecutions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTaskExecutions(options.Region), middleware.Before); err != nil {
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

// ListTaskExecutionsAPIClient is a client that implements the ListTaskExecutions
// operation.
type ListTaskExecutionsAPIClient interface {
	ListTaskExecutions(context.Context, *ListTaskExecutionsInput, ...func(*Options)) (*ListTaskExecutionsOutput, error)
}

var _ ListTaskExecutionsAPIClient = (*Client)(nil)

// ListTaskExecutionsPaginatorOptions is the paginator options for
// ListTaskExecutions
type ListTaskExecutionsPaginatorOptions struct {
	// The maximum number of executed tasks to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTaskExecutionsPaginator is a paginator for ListTaskExecutions
type ListTaskExecutionsPaginator struct {
	options   ListTaskExecutionsPaginatorOptions
	client    ListTaskExecutionsAPIClient
	params    *ListTaskExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListTaskExecutionsPaginator returns a new ListTaskExecutionsPaginator
func NewListTaskExecutionsPaginator(client ListTaskExecutionsAPIClient, params *ListTaskExecutionsInput, optFns ...func(*ListTaskExecutionsPaginatorOptions)) *ListTaskExecutionsPaginator {
	if params == nil {
		params = &ListTaskExecutionsInput{}
	}

	options := ListTaskExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTaskExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTaskExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTaskExecutions page.
func (p *ListTaskExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTaskExecutionsOutput, error) {
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

	result, err := p.client.ListTaskExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTaskExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "ListTaskExecutions",
	}
}
