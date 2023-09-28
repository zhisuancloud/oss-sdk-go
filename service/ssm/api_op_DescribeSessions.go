// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of all active sessions (both connected and disconnected) or
// terminated sessions from the past 30 days.
func (c *Client) DescribeSessions(ctx context.Context, params *DescribeSessionsInput, optFns ...func(*Options)) (*DescribeSessionsOutput, error) {
	if params == nil {
		params = &DescribeSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSessions", params, optFns, c.addOperationDescribeSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSessionsInput struct {

	// The session status to retrieve a list of sessions for. For example, "Active".
	//
	// This member is required.
	State types.SessionState

	// One or more filters to limit the type of sessions returned by the request.
	Filters []types.SessionFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeSessionsOutput struct {

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// A list of sessions meeting the request parameters.
	Sessions []types.Session

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSessions{}, middleware.After)
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
	if err = addOpDescribeSessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSessions(options.Region), middleware.Before); err != nil {
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

// DescribeSessionsAPIClient is a client that implements the DescribeSessions
// operation.
type DescribeSessionsAPIClient interface {
	DescribeSessions(context.Context, *DescribeSessionsInput, ...func(*Options)) (*DescribeSessionsOutput, error)
}

var _ DescribeSessionsAPIClient = (*Client)(nil)

// DescribeSessionsPaginatorOptions is the paginator options for DescribeSessions
type DescribeSessionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSessionsPaginator is a paginator for DescribeSessions
type DescribeSessionsPaginator struct {
	options   DescribeSessionsPaginatorOptions
	client    DescribeSessionsAPIClient
	params    *DescribeSessionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSessionsPaginator returns a new DescribeSessionsPaginator
func NewDescribeSessionsPaginator(client DescribeSessionsAPIClient, params *DescribeSessionsInput, optFns ...func(*DescribeSessionsPaginatorOptions)) *DescribeSessionsPaginator {
	if params == nil {
		params = &DescribeSessionsInput{}
	}

	options := DescribeSessionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSessions page.
func (p *DescribeSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSessionsOutput, error) {
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

	result, err := p.client.DescribeSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeSessions",
	}
}
