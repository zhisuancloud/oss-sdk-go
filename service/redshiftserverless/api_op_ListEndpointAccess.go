// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of EndpointAccess objects and relevant information.
func (c *Client) ListEndpointAccess(ctx context.Context, params *ListEndpointAccessInput, optFns ...func(*Options)) (*ListEndpointAccessOutput, error) {
	if params == nil {
		params = &ListEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpointAccess", params, optFns, c.addOperationListEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointAccessInput struct {

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults *int32

	// If your initial ListEndpointAccess operation returns a nextToken, you can
	// include the returned nextToken in subsequent ListEndpointAccess operations,
	// which returns results in the next page.
	NextToken *string

	// The unique identifier of the virtual private cloud with access to Amazon
	// Redshift Serverless.
	VpcId *string

	// The name of the workgroup associated with the VPC endpoint to return.
	WorkgroupName *string

	noSmithyDocumentSerde
}

type ListEndpointAccessOutput struct {

	// The returned VPC endpoints.
	//
	// This member is required.
	Endpoints []types.EndpointAccess

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpointAccess{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpointAccess(options.Region), middleware.Before); err != nil {
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

// ListEndpointAccessAPIClient is a client that implements the ListEndpointAccess
// operation.
type ListEndpointAccessAPIClient interface {
	ListEndpointAccess(context.Context, *ListEndpointAccessInput, ...func(*Options)) (*ListEndpointAccessOutput, error)
}

var _ ListEndpointAccessAPIClient = (*Client)(nil)

// ListEndpointAccessPaginatorOptions is the paginator options for
// ListEndpointAccess
type ListEndpointAccessPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEndpointAccessPaginator is a paginator for ListEndpointAccess
type ListEndpointAccessPaginator struct {
	options   ListEndpointAccessPaginatorOptions
	client    ListEndpointAccessAPIClient
	params    *ListEndpointAccessInput
	nextToken *string
	firstPage bool
}

// NewListEndpointAccessPaginator returns a new ListEndpointAccessPaginator
func NewListEndpointAccessPaginator(client ListEndpointAccessAPIClient, params *ListEndpointAccessInput, optFns ...func(*ListEndpointAccessPaginatorOptions)) *ListEndpointAccessPaginator {
	if params == nil {
		params = &ListEndpointAccessInput{}
	}

	options := ListEndpointAccessPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEndpointAccessPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEndpointAccessPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEndpointAccess page.
func (p *ListEndpointAccessPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEndpointAccessOutput, error) {
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

	result, err := p.client.ListEndpointAccess(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-serverless",
		OperationName: "ListEndpointAccess",
	}
}