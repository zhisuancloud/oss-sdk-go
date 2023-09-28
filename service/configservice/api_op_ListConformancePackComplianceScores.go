// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of conformance pack compliance scores. A compliance score is the
// percentage of the number of compliant rule-resource combinations in a
// conformance pack compared to the number of total possible rule-resource
// combinations in the conformance pack. This metric provides you with a high-level
// view of the compliance state of your conformance packs, and can be used to
// identify, investigate, and understand the level of compliance in your
// conformance packs. Conformance packs with no evaluation results will have a
// compliance score of INSUFFICIENT_DATA.
func (c *Client) ListConformancePackComplianceScores(ctx context.Context, params *ListConformancePackComplianceScoresInput, optFns ...func(*Options)) (*ListConformancePackComplianceScoresOutput, error) {
	if params == nil {
		params = &ListConformancePackComplianceScoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConformancePackComplianceScores", params, optFns, c.addOperationListConformancePackComplianceScoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConformancePackComplianceScoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConformancePackComplianceScoresInput struct {

	// Filters the results based on the ConformancePackComplianceScoresFilters.
	Filters *types.ConformancePackComplianceScoresFilters

	// The maximum number of conformance pack compliance scores returned on each page.
	Limit int32

	// The nextToken string in a prior request that you can use to get the paginated
	// response for next set of conformance pack compliance scores.
	NextToken *string

	// Sorts your conformance pack compliance scores in either ascending or descending
	// order, depending on SortOrder. By default, conformance pack compliance scores
	// are sorted in ascending order by compliance score and alphabetically by name of
	// the conformance pack if there is more than one conformance pack with the same
	// compliance score.
	SortBy types.SortBy

	// Determines the order in which conformance pack compliance scores are sorted.
	// Either in ascending or descending order. Conformance packs with a compliance
	// score of INSUFFICIENT_DATA will be first when sorting by ascending order and
	// last when sorting by descending order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListConformancePackComplianceScoresOutput struct {

	// A list of ConformancePackComplianceScore objects.
	//
	// This member is required.
	ConformancePackComplianceScores []types.ConformancePackComplianceScore

	// The nextToken string that you can use to get the next page of results in a
	// paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConformancePackComplianceScoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListConformancePackComplianceScores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConformancePackComplianceScores{}, middleware.After)
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
	if err = addOpListConformancePackComplianceScoresValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConformancePackComplianceScores(options.Region), middleware.Before); err != nil {
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

// ListConformancePackComplianceScoresAPIClient is a client that implements the
// ListConformancePackComplianceScores operation.
type ListConformancePackComplianceScoresAPIClient interface {
	ListConformancePackComplianceScores(context.Context, *ListConformancePackComplianceScoresInput, ...func(*Options)) (*ListConformancePackComplianceScoresOutput, error)
}

var _ ListConformancePackComplianceScoresAPIClient = (*Client)(nil)

// ListConformancePackComplianceScoresPaginatorOptions is the paginator options for
// ListConformancePackComplianceScores
type ListConformancePackComplianceScoresPaginatorOptions struct {
	// The maximum number of conformance pack compliance scores returned on each page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConformancePackComplianceScoresPaginator is a paginator for
// ListConformancePackComplianceScores
type ListConformancePackComplianceScoresPaginator struct {
	options   ListConformancePackComplianceScoresPaginatorOptions
	client    ListConformancePackComplianceScoresAPIClient
	params    *ListConformancePackComplianceScoresInput
	nextToken *string
	firstPage bool
}

// NewListConformancePackComplianceScoresPaginator returns a new
// ListConformancePackComplianceScoresPaginator
func NewListConformancePackComplianceScoresPaginator(client ListConformancePackComplianceScoresAPIClient, params *ListConformancePackComplianceScoresInput, optFns ...func(*ListConformancePackComplianceScoresPaginatorOptions)) *ListConformancePackComplianceScoresPaginator {
	if params == nil {
		params = &ListConformancePackComplianceScoresInput{}
	}

	options := ListConformancePackComplianceScoresPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConformancePackComplianceScoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConformancePackComplianceScoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConformancePackComplianceScores page.
func (p *ListConformancePackComplianceScoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConformancePackComplianceScoresOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.ListConformancePackComplianceScores(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConformancePackComplianceScores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "ListConformancePackComplianceScores",
	}
}
