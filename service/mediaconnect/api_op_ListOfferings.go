// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays a list of all offerings that are available to this account in the
// current AWS Region. If you have an active reservation (which means you've
// purchased an offering that has already started and hasn't expired yet), your
// account isn't eligible for other offerings.
func (c *Client) ListOfferings(ctx context.Context, params *ListOfferingsInput, optFns ...func(*Options)) (*ListOfferingsOutput, error) {
	if params == nil {
		params = &ListOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOfferings", params, optFns, c.addOperationListOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOfferingsInput struct {

	// The maximum number of results to return per API request. For example, you submit
	// a ListOfferings request with MaxResults set at 5. Although 20 items match your
	// request, the service returns no more than the first 5 items. (The service also
	// returns a NextToken value that you can use to fetch the next batch of results.)
	// The service might return fewer results than the MaxResults value. If MaxResults
	// is not included in the request, the service defaults to pagination with a
	// maximum of 10 results per page.
	MaxResults int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListOfferings request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListOfferings request a second
	// time and specify the NextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOfferingsOutput struct {

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListOfferings request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListOfferings request a second
	// time and specify the NextToken value.
	NextToken *string

	// A list of offerings that are available to this account in the current AWS
	// Region.
	Offerings []types.Offering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOfferings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferings(options.Region), middleware.Before); err != nil {
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

// ListOfferingsAPIClient is a client that implements the ListOfferings operation.
type ListOfferingsAPIClient interface {
	ListOfferings(context.Context, *ListOfferingsInput, ...func(*Options)) (*ListOfferingsOutput, error)
}

var _ ListOfferingsAPIClient = (*Client)(nil)

// ListOfferingsPaginatorOptions is the paginator options for ListOfferings
type ListOfferingsPaginatorOptions struct {
	// The maximum number of results to return per API request. For example, you submit
	// a ListOfferings request with MaxResults set at 5. Although 20 items match your
	// request, the service returns no more than the first 5 items. (The service also
	// returns a NextToken value that you can use to fetch the next batch of results.)
	// The service might return fewer results than the MaxResults value. If MaxResults
	// is not included in the request, the service defaults to pagination with a
	// maximum of 10 results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOfferingsPaginator is a paginator for ListOfferings
type ListOfferingsPaginator struct {
	options   ListOfferingsPaginatorOptions
	client    ListOfferingsAPIClient
	params    *ListOfferingsInput
	nextToken *string
	firstPage bool
}

// NewListOfferingsPaginator returns a new ListOfferingsPaginator
func NewListOfferingsPaginator(client ListOfferingsAPIClient, params *ListOfferingsInput, optFns ...func(*ListOfferingsPaginatorOptions)) *ListOfferingsPaginator {
	if params == nil {
		params = &ListOfferingsInput{}
	}

	options := ListOfferingsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOfferings page.
func (p *ListOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOfferingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListOfferings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "ListOfferings",
	}
}
