// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of failures for the specified import.
func (c *Client) ListImportFailures(ctx context.Context, params *ListImportFailuresInput, optFns ...func(*Options)) (*ListImportFailuresOutput, error) {
	if params == nil {
		params = &ListImportFailuresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImportFailures", params, optFns, c.addOperationListImportFailuresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImportFailuresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImportFailuresInput struct {

	// The ID of the import.
	//
	// This member is required.
	ImportId *string

	// The maximum number of failures to display on a single page.
	MaxResults *int32

	// A token you can use to get the next page of import failures.
	NextToken *string

	noSmithyDocumentSerde
}

type ListImportFailuresOutput struct {

	// Contains information about the import failures.
	Failures []types.ImportFailureListItem

	// A token you can use to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImportFailuresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListImportFailures{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListImportFailures{}, middleware.After)
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
	if err = addOpListImportFailuresValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImportFailures(options.Region), middleware.Before); err != nil {
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

// ListImportFailuresAPIClient is a client that implements the ListImportFailures
// operation.
type ListImportFailuresAPIClient interface {
	ListImportFailures(context.Context, *ListImportFailuresInput, ...func(*Options)) (*ListImportFailuresOutput, error)
}

var _ ListImportFailuresAPIClient = (*Client)(nil)

// ListImportFailuresPaginatorOptions is the paginator options for
// ListImportFailures
type ListImportFailuresPaginatorOptions struct {
	// The maximum number of failures to display on a single page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImportFailuresPaginator is a paginator for ListImportFailures
type ListImportFailuresPaginator struct {
	options   ListImportFailuresPaginatorOptions
	client    ListImportFailuresAPIClient
	params    *ListImportFailuresInput
	nextToken *string
	firstPage bool
}

// NewListImportFailuresPaginator returns a new ListImportFailuresPaginator
func NewListImportFailuresPaginator(client ListImportFailuresAPIClient, params *ListImportFailuresInput, optFns ...func(*ListImportFailuresPaginatorOptions)) *ListImportFailuresPaginator {
	if params == nil {
		params = &ListImportFailuresInput{}
	}

	options := ListImportFailuresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImportFailuresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImportFailuresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImportFailures page.
func (p *ListImportFailuresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImportFailuresOutput, error) {
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

	result, err := p.client.ListImportFailures(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImportFailures(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "ListImportFailures",
	}
}