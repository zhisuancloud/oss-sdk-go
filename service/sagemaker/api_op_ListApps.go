// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists apps.
func (c *Client) ListApps(ctx context.Context, params *ListAppsInput, optFns ...func(*Options)) (*ListAppsOutput, error) {
	if params == nil {
		params = &ListAppsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApps", params, optFns, c.addOperationListAppsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppsInput struct {

	// A parameter to search for the domain ID.
	DomainIdEquals *string

	// Returns a list up to a specified limit.
	MaxResults *int32

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy types.AppSortKey

	// The sort order for the results. The default is Ascending.
	SortOrder types.SortOrder

	// A parameter to search by user profile name.
	UserProfileNameEquals *string

	noSmithyDocumentSerde
}

type ListAppsOutput struct {

	// The list of apps.
	Apps []types.AppDetails

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApps{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApps(options.Region), middleware.Before); err != nil {
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

// ListAppsAPIClient is a client that implements the ListApps operation.
type ListAppsAPIClient interface {
	ListApps(context.Context, *ListAppsInput, ...func(*Options)) (*ListAppsOutput, error)
}

var _ ListAppsAPIClient = (*Client)(nil)

// ListAppsPaginatorOptions is the paginator options for ListApps
type ListAppsPaginatorOptions struct {
	// Returns a list up to a specified limit.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppsPaginator is a paginator for ListApps
type ListAppsPaginator struct {
	options   ListAppsPaginatorOptions
	client    ListAppsAPIClient
	params    *ListAppsInput
	nextToken *string
	firstPage bool
}

// NewListAppsPaginator returns a new ListAppsPaginator
func NewListAppsPaginator(client ListAppsAPIClient, params *ListAppsInput, optFns ...func(*ListAppsPaginatorOptions)) *ListAppsPaginator {
	if params == nil {
		params = &ListAppsInput{}
	}

	options := ListAppsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApps page.
func (p *ListAppsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppsOutput, error) {
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

	result, err := p.client.ListApps(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListApps",
	}
}
