// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcampaigns

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/connectcampaigns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides summary information about the campaigns under the specified Amazon
// Connect account.
func (c *Client) ListCampaigns(ctx context.Context, params *ListCampaignsInput, optFns ...func(*Options)) (*ListCampaignsOutput, error) {
	if params == nil {
		params = &ListCampaignsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCampaigns", params, optFns, c.addOperationListCampaignsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCampaignsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListCampaignsRequest
type ListCampaignsInput struct {

	// Filter model by type
	Filters *types.CampaignFilters

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

// ListCampaignsResponse
type ListCampaignsOutput struct {

	// A list of Amazon Connect campaigns.
	CampaignSummaryList []types.CampaignSummary

	// The token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCampaignsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCampaigns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCampaigns{}, middleware.After)
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
	if err = addOpListCampaignsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCampaigns(options.Region), middleware.Before); err != nil {
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

// ListCampaignsAPIClient is a client that implements the ListCampaigns operation.
type ListCampaignsAPIClient interface {
	ListCampaigns(context.Context, *ListCampaignsInput, ...func(*Options)) (*ListCampaignsOutput, error)
}

var _ ListCampaignsAPIClient = (*Client)(nil)

// ListCampaignsPaginatorOptions is the paginator options for ListCampaigns
type ListCampaignsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCampaignsPaginator is a paginator for ListCampaigns
type ListCampaignsPaginator struct {
	options   ListCampaignsPaginatorOptions
	client    ListCampaignsAPIClient
	params    *ListCampaignsInput
	nextToken *string
	firstPage bool
}

// NewListCampaignsPaginator returns a new ListCampaignsPaginator
func NewListCampaignsPaginator(client ListCampaignsAPIClient, params *ListCampaignsInput, optFns ...func(*ListCampaignsPaginatorOptions)) *ListCampaignsPaginator {
	if params == nil {
		params = &ListCampaignsInput{}
	}

	options := ListCampaignsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCampaignsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCampaignsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCampaigns page.
func (p *ListCampaignsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCampaignsOutput, error) {
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

	result, err := p.client.ListCampaigns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCampaigns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect-campaigns",
		OperationName: "ListCampaigns",
	}
}
