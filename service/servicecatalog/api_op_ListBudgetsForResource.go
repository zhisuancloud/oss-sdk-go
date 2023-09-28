// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the budgets associated to the specified resource.
func (c *Client) ListBudgetsForResource(ctx context.Context, params *ListBudgetsForResourceInput, optFns ...func(*Options)) (*ListBudgetsForResourceOutput, error) {
	if params == nil {
		params = &ListBudgetsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBudgetsForResource", params, optFns, c.addOperationListBudgetsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBudgetsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBudgetsForResourceInput struct {

	// The resource identifier.
	//
	// This member is required.
	ResourceId *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	noSmithyDocumentSerde
}

type ListBudgetsForResourceOutput struct {

	// Information about the associated budgets.
	Budgets []types.BudgetDetail

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBudgetsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBudgetsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBudgetsForResource{}, middleware.After)
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
	if err = addOpListBudgetsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBudgetsForResource(options.Region), middleware.Before); err != nil {
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

// ListBudgetsForResourceAPIClient is a client that implements the
// ListBudgetsForResource operation.
type ListBudgetsForResourceAPIClient interface {
	ListBudgetsForResource(context.Context, *ListBudgetsForResourceInput, ...func(*Options)) (*ListBudgetsForResourceOutput, error)
}

var _ ListBudgetsForResourceAPIClient = (*Client)(nil)

// ListBudgetsForResourcePaginatorOptions is the paginator options for
// ListBudgetsForResource
type ListBudgetsForResourcePaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBudgetsForResourcePaginator is a paginator for ListBudgetsForResource
type ListBudgetsForResourcePaginator struct {
	options   ListBudgetsForResourcePaginatorOptions
	client    ListBudgetsForResourceAPIClient
	params    *ListBudgetsForResourceInput
	nextToken *string
	firstPage bool
}

// NewListBudgetsForResourcePaginator returns a new ListBudgetsForResourcePaginator
func NewListBudgetsForResourcePaginator(client ListBudgetsForResourceAPIClient, params *ListBudgetsForResourceInput, optFns ...func(*ListBudgetsForResourcePaginatorOptions)) *ListBudgetsForResourcePaginator {
	if params == nil {
		params = &ListBudgetsForResourceInput{}
	}

	options := ListBudgetsForResourcePaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBudgetsForResourcePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBudgetsForResourcePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBudgetsForResource page.
func (p *ListBudgetsForResourcePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBudgetsForResourceOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListBudgetsForResource(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListBudgetsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListBudgetsForResource",
	}
}
