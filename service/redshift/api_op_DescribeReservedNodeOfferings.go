// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the available reserved node offerings by Amazon Redshift with
// their descriptions including the node type, the fixed and recurring costs of
// reserving the node and duration the node will be reserved for you. These
// descriptions help you determine which reserve node offering you want to
// purchase. You then use the unique offering ID in you call to
// PurchaseReservedNodeOffering to reserve one or more nodes for your Amazon
// Redshift cluster. For more information about reserved node offerings, go to
// Purchasing Reserved Nodes
// (https://docs.aws.amazon.com/redshift/latest/mgmt/purchase-reserved-node-instance.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) DescribeReservedNodeOfferings(ctx context.Context, params *DescribeReservedNodeOfferingsInput, optFns ...func(*Options)) (*DescribeReservedNodeOfferingsOutput, error) {
	if params == nil {
		params = &DescribeReservedNodeOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedNodeOfferings", params, optFns, c.addOperationDescribeReservedNodeOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedNodeOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReservedNodeOfferingsInput struct {

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeReservedNodeOfferings request
	// exceed the value specified in MaxRecords, Amazon Web Services returns a value in
	// the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The unique identifier for the offering.
	ReservedNodeOfferingId *string

	noSmithyDocumentSerde
}

type DescribeReservedNodeOfferingsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of ReservedNodeOffering objects.
	ReservedNodeOfferings []types.ReservedNodeOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedNodeOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReservedNodeOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReservedNodeOfferings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedNodeOfferings(options.Region), middleware.Before); err != nil {
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

// DescribeReservedNodeOfferingsAPIClient is a client that implements the
// DescribeReservedNodeOfferings operation.
type DescribeReservedNodeOfferingsAPIClient interface {
	DescribeReservedNodeOfferings(context.Context, *DescribeReservedNodeOfferingsInput, ...func(*Options)) (*DescribeReservedNodeOfferingsOutput, error)
}

var _ DescribeReservedNodeOfferingsAPIClient = (*Client)(nil)

// DescribeReservedNodeOfferingsPaginatorOptions is the paginator options for
// DescribeReservedNodeOfferings
type DescribeReservedNodeOfferingsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedNodeOfferingsPaginator is a paginator for
// DescribeReservedNodeOfferings
type DescribeReservedNodeOfferingsPaginator struct {
	options   DescribeReservedNodeOfferingsPaginatorOptions
	client    DescribeReservedNodeOfferingsAPIClient
	params    *DescribeReservedNodeOfferingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedNodeOfferingsPaginator returns a new
// DescribeReservedNodeOfferingsPaginator
func NewDescribeReservedNodeOfferingsPaginator(client DescribeReservedNodeOfferingsAPIClient, params *DescribeReservedNodeOfferingsInput, optFns ...func(*DescribeReservedNodeOfferingsPaginatorOptions)) *DescribeReservedNodeOfferingsPaginator {
	if params == nil {
		params = &DescribeReservedNodeOfferingsInput{}
	}

	options := DescribeReservedNodeOfferingsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedNodeOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedNodeOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedNodeOfferings page.
func (p *DescribeReservedNodeOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedNodeOfferingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeReservedNodeOfferings(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeReservedNodeOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeReservedNodeOfferings",
	}
}
