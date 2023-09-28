// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of core network policy versions.
func (c *Client) ListCoreNetworkPolicyVersions(ctx context.Context, params *ListCoreNetworkPolicyVersionsInput, optFns ...func(*Options)) (*ListCoreNetworkPolicyVersionsOutput, error) {
	if params == nil {
		params = &ListCoreNetworkPolicyVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCoreNetworkPolicyVersions", params, optFns, c.addOperationListCoreNetworkPolicyVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCoreNetworkPolicyVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCoreNetworkPolicyVersionsInput struct {

	// The ID of a core network.
	//
	// This member is required.
	CoreNetworkId *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCoreNetworkPolicyVersionsOutput struct {

	// Describes core network policy versions.
	CoreNetworkPolicyVersions []types.CoreNetworkPolicyVersion

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCoreNetworkPolicyVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCoreNetworkPolicyVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCoreNetworkPolicyVersions{}, middleware.After)
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
	if err = addOpListCoreNetworkPolicyVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCoreNetworkPolicyVersions(options.Region), middleware.Before); err != nil {
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

// ListCoreNetworkPolicyVersionsAPIClient is a client that implements the
// ListCoreNetworkPolicyVersions operation.
type ListCoreNetworkPolicyVersionsAPIClient interface {
	ListCoreNetworkPolicyVersions(context.Context, *ListCoreNetworkPolicyVersionsInput, ...func(*Options)) (*ListCoreNetworkPolicyVersionsOutput, error)
}

var _ ListCoreNetworkPolicyVersionsAPIClient = (*Client)(nil)

// ListCoreNetworkPolicyVersionsPaginatorOptions is the paginator options for
// ListCoreNetworkPolicyVersions
type ListCoreNetworkPolicyVersionsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCoreNetworkPolicyVersionsPaginator is a paginator for
// ListCoreNetworkPolicyVersions
type ListCoreNetworkPolicyVersionsPaginator struct {
	options   ListCoreNetworkPolicyVersionsPaginatorOptions
	client    ListCoreNetworkPolicyVersionsAPIClient
	params    *ListCoreNetworkPolicyVersionsInput
	nextToken *string
	firstPage bool
}

// NewListCoreNetworkPolicyVersionsPaginator returns a new
// ListCoreNetworkPolicyVersionsPaginator
func NewListCoreNetworkPolicyVersionsPaginator(client ListCoreNetworkPolicyVersionsAPIClient, params *ListCoreNetworkPolicyVersionsInput, optFns ...func(*ListCoreNetworkPolicyVersionsPaginatorOptions)) *ListCoreNetworkPolicyVersionsPaginator {
	if params == nil {
		params = &ListCoreNetworkPolicyVersionsInput{}
	}

	options := ListCoreNetworkPolicyVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCoreNetworkPolicyVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCoreNetworkPolicyVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCoreNetworkPolicyVersions page.
func (p *ListCoreNetworkPolicyVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCoreNetworkPolicyVersionsOutput, error) {
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

	result, err := p.client.ListCoreNetworkPolicyVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCoreNetworkPolicyVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "ListCoreNetworkPolicyVersions",
	}
}