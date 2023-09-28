// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the SIP media applications under the administrator's AWS account.
func (c *Client) ListSipMediaApplications(ctx context.Context, params *ListSipMediaApplicationsInput, optFns ...func(*Options)) (*ListSipMediaApplicationsOutput, error) {
	if params == nil {
		params = &ListSipMediaApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSipMediaApplications", params, optFns, c.addOperationListSipMediaApplicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSipMediaApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSipMediaApplicationsInput struct {

	// The maximum number of results to return in a single call. Defaults to 100.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSipMediaApplicationsOutput struct {

	// The token to use to retrieve the next page of results.
	NextToken *string

	// List of SIP media applications and application details.
	SipMediaApplications []types.SipMediaApplication

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSipMediaApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSipMediaApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSipMediaApplications{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSipMediaApplications(options.Region), middleware.Before); err != nil {
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

// ListSipMediaApplicationsAPIClient is a client that implements the
// ListSipMediaApplications operation.
type ListSipMediaApplicationsAPIClient interface {
	ListSipMediaApplications(context.Context, *ListSipMediaApplicationsInput, ...func(*Options)) (*ListSipMediaApplicationsOutput, error)
}

var _ ListSipMediaApplicationsAPIClient = (*Client)(nil)

// ListSipMediaApplicationsPaginatorOptions is the paginator options for
// ListSipMediaApplications
type ListSipMediaApplicationsPaginatorOptions struct {
	// The maximum number of results to return in a single call. Defaults to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSipMediaApplicationsPaginator is a paginator for ListSipMediaApplications
type ListSipMediaApplicationsPaginator struct {
	options   ListSipMediaApplicationsPaginatorOptions
	client    ListSipMediaApplicationsAPIClient
	params    *ListSipMediaApplicationsInput
	nextToken *string
	firstPage bool
}

// NewListSipMediaApplicationsPaginator returns a new
// ListSipMediaApplicationsPaginator
func NewListSipMediaApplicationsPaginator(client ListSipMediaApplicationsAPIClient, params *ListSipMediaApplicationsInput, optFns ...func(*ListSipMediaApplicationsPaginatorOptions)) *ListSipMediaApplicationsPaginator {
	if params == nil {
		params = &ListSipMediaApplicationsInput{}
	}

	options := ListSipMediaApplicationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSipMediaApplicationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSipMediaApplicationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSipMediaApplications page.
func (p *ListSipMediaApplicationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSipMediaApplicationsOutput, error) {
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

	result, err := p.client.ListSipMediaApplications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSipMediaApplications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListSipMediaApplications",
	}
}
