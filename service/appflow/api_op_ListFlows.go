// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the flows associated with your account.
func (c *Client) ListFlows(ctx context.Context, params *ListFlowsInput, optFns ...func(*Options)) (*ListFlowsOutput, error) {
	if params == nil {
		params = &ListFlowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFlows", params, optFns, c.addOperationListFlowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFlowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlowsInput struct {

	// Specifies the maximum number of items that should be returned in the result set.
	MaxResults *int32

	// The pagination token for next page of data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFlowsOutput struct {

	// The list of flows associated with your account.
	Flows []types.FlowDefinition

	// The pagination token for next page of data.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFlowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFlows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFlows{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFlows(options.Region), middleware.Before); err != nil {
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

// ListFlowsAPIClient is a client that implements the ListFlows operation.
type ListFlowsAPIClient interface {
	ListFlows(context.Context, *ListFlowsInput, ...func(*Options)) (*ListFlowsOutput, error)
}

var _ ListFlowsAPIClient = (*Client)(nil)

// ListFlowsPaginatorOptions is the paginator options for ListFlows
type ListFlowsPaginatorOptions struct {
	// Specifies the maximum number of items that should be returned in the result set.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFlowsPaginator is a paginator for ListFlows
type ListFlowsPaginator struct {
	options   ListFlowsPaginatorOptions
	client    ListFlowsAPIClient
	params    *ListFlowsInput
	nextToken *string
	firstPage bool
}

// NewListFlowsPaginator returns a new ListFlowsPaginator
func NewListFlowsPaginator(client ListFlowsAPIClient, params *ListFlowsInput, optFns ...func(*ListFlowsPaginatorOptions)) *ListFlowsPaginator {
	if params == nil {
		params = &ListFlowsInput{}
	}

	options := ListFlowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFlowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFlowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFlows page.
func (p *ListFlowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFlowsOutput, error) {
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

	result, err := p.client.ListFlows(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFlows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "ListFlows",
	}
}
