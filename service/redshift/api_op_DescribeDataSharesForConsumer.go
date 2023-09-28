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

// Returns a list of datashares where the account identifier being called is a
// consumer account identifier.
func (c *Client) DescribeDataSharesForConsumer(ctx context.Context, params *DescribeDataSharesForConsumerInput, optFns ...func(*Options)) (*DescribeDataSharesForConsumerOutput, error) {
	if params == nil {
		params = &DescribeDataSharesForConsumerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataSharesForConsumer", params, optFns, c.addOperationDescribeDataSharesForConsumerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataSharesForConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataSharesForConsumerInput struct {

	// The Amazon Resource Name (ARN) of the consumer that returns in the list of
	// datashares.
	ConsumerArn *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeDataSharesForConsumer request
	// exceed the value specified in MaxRecords, Amazon Web Services returns a value in
	// the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32

	// An identifier giving the status of a datashare in the consumer cluster. If this
	// field is specified, Amazon Redshift returns the list of datashares that have the
	// specified status.
	Status types.DataShareStatusForConsumer

	noSmithyDocumentSerde
}

type DescribeDataSharesForConsumerOutput struct {

	// Shows the results of datashares available for consumers.
	DataShares []types.DataShare

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeDataSharesForConsumer request
	// exceed the value specified in MaxRecords, Amazon Web Services returns a value in
	// the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDataSharesForConsumerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDataSharesForConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDataSharesForConsumer{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataSharesForConsumer(options.Region), middleware.Before); err != nil {
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

// DescribeDataSharesForConsumerAPIClient is a client that implements the
// DescribeDataSharesForConsumer operation.
type DescribeDataSharesForConsumerAPIClient interface {
	DescribeDataSharesForConsumer(context.Context, *DescribeDataSharesForConsumerInput, ...func(*Options)) (*DescribeDataSharesForConsumerOutput, error)
}

var _ DescribeDataSharesForConsumerAPIClient = (*Client)(nil)

// DescribeDataSharesForConsumerPaginatorOptions is the paginator options for
// DescribeDataSharesForConsumer
type DescribeDataSharesForConsumerPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDataSharesForConsumerPaginator is a paginator for
// DescribeDataSharesForConsumer
type DescribeDataSharesForConsumerPaginator struct {
	options   DescribeDataSharesForConsumerPaginatorOptions
	client    DescribeDataSharesForConsumerAPIClient
	params    *DescribeDataSharesForConsumerInput
	nextToken *string
	firstPage bool
}

// NewDescribeDataSharesForConsumerPaginator returns a new
// DescribeDataSharesForConsumerPaginator
func NewDescribeDataSharesForConsumerPaginator(client DescribeDataSharesForConsumerAPIClient, params *DescribeDataSharesForConsumerInput, optFns ...func(*DescribeDataSharesForConsumerPaginatorOptions)) *DescribeDataSharesForConsumerPaginator {
	if params == nil {
		params = &DescribeDataSharesForConsumerInput{}
	}

	options := DescribeDataSharesForConsumerPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDataSharesForConsumerPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDataSharesForConsumerPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDataSharesForConsumer page.
func (p *DescribeDataSharesForConsumerPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDataSharesForConsumerOutput, error) {
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

	result, err := p.client.DescribeDataSharesForConsumer(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDataSharesForConsumer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeDataSharesForConsumer",
	}
}