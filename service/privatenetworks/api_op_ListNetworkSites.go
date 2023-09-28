// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/privatenetworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists network sites. Add filters to your request to return a more specific list
// of results. Use filters to match the status of the network site.
func (c *Client) ListNetworkSites(ctx context.Context, params *ListNetworkSitesInput, optFns ...func(*Options)) (*ListNetworkSitesOutput, error) {
	if params == nil {
		params = &ListNetworkSitesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNetworkSites", params, optFns, c.addOperationListNetworkSitesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNetworkSitesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNetworkSitesInput struct {

	// The Amazon Resource Name (ARN) of the network.
	//
	// This member is required.
	NetworkArn *string

	// The filters. Add filters to your request to return a more specific list of
	// results. Use filters to match the status of the network sites.
	//
	// * STATUS - The
	// status (AVAILABLE | CREATED | DELETED | DEPROVISIONING | PROVISIONING).
	//
	// Filter
	// values are case sensitive. If you specify multiple values for a filter, the
	// values are joined with an OR, and the request returns all results that match any
	// of the specified values.
	Filters map[string][]string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	StartToken *string

	noSmithyDocumentSerde
}

type ListNetworkSitesOutput struct {

	// Information about the network sites.
	NetworkSites []types.NetworkSite

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNetworkSitesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNetworkSites{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNetworkSites{}, middleware.After)
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
	if err = addOpListNetworkSitesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNetworkSites(options.Region), middleware.Before); err != nil {
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

// ListNetworkSitesAPIClient is a client that implements the ListNetworkSites
// operation.
type ListNetworkSitesAPIClient interface {
	ListNetworkSites(context.Context, *ListNetworkSitesInput, ...func(*Options)) (*ListNetworkSitesOutput, error)
}

var _ ListNetworkSitesAPIClient = (*Client)(nil)

// ListNetworkSitesPaginatorOptions is the paginator options for ListNetworkSites
type ListNetworkSitesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNetworkSitesPaginator is a paginator for ListNetworkSites
type ListNetworkSitesPaginator struct {
	options   ListNetworkSitesPaginatorOptions
	client    ListNetworkSitesAPIClient
	params    *ListNetworkSitesInput
	nextToken *string
	firstPage bool
}

// NewListNetworkSitesPaginator returns a new ListNetworkSitesPaginator
func NewListNetworkSitesPaginator(client ListNetworkSitesAPIClient, params *ListNetworkSitesInput, optFns ...func(*ListNetworkSitesPaginatorOptions)) *ListNetworkSitesPaginator {
	if params == nil {
		params = &ListNetworkSitesInput{}
	}

	options := ListNetworkSitesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNetworkSitesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.StartToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNetworkSitesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNetworkSites page.
func (p *ListNetworkSitesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNetworkSitesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.StartToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListNetworkSites(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNetworkSites(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "private-networks",
		OperationName: "ListNetworkSites",
	}
}
