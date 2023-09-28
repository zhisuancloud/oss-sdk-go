// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Amazon QuickSight analyses that exist in the specified Amazon Web Services
// account.
func (c *Client) ListAnalyses(ctx context.Context, params *ListAnalysesInput, optFns ...func(*Options)) (*ListAnalysesOutput, error) {
	if params == nil {
		params = &ListAnalysesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnalyses", params, optFns, c.addOperationListAnalysesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnalysesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnalysesInput struct {

	// The ID of the Amazon Web Services account that contains the analyses.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAnalysesOutput struct {

	// Metadata describing each of the analyses that are listed.
	AnalysisSummaryList []types.AnalysisSummary

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnalysesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnalyses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnalyses{}, middleware.After)
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
	if err = addOpListAnalysesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnalyses(options.Region), middleware.Before); err != nil {
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

// ListAnalysesAPIClient is a client that implements the ListAnalyses operation.
type ListAnalysesAPIClient interface {
	ListAnalyses(context.Context, *ListAnalysesInput, ...func(*Options)) (*ListAnalysesOutput, error)
}

var _ ListAnalysesAPIClient = (*Client)(nil)

// ListAnalysesPaginatorOptions is the paginator options for ListAnalyses
type ListAnalysesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnalysesPaginator is a paginator for ListAnalyses
type ListAnalysesPaginator struct {
	options   ListAnalysesPaginatorOptions
	client    ListAnalysesAPIClient
	params    *ListAnalysesInput
	nextToken *string
	firstPage bool
}

// NewListAnalysesPaginator returns a new ListAnalysesPaginator
func NewListAnalysesPaginator(client ListAnalysesAPIClient, params *ListAnalysesInput, optFns ...func(*ListAnalysesPaginatorOptions)) *ListAnalysesPaginator {
	if params == nil {
		params = &ListAnalysesInput{}
	}

	options := ListAnalysesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnalysesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnalysesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnalyses page.
func (p *ListAnalysesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnalysesOutput, error) {
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

	result, err := p.client.ListAnalyses(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAnalyses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListAnalyses",
	}
}