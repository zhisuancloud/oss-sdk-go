// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a list of model quality monitoring job definitions in your account.
func (c *Client) ListModelQualityJobDefinitions(ctx context.Context, params *ListModelQualityJobDefinitionsInput, optFns ...func(*Options)) (*ListModelQualityJobDefinitionsOutput, error) {
	if params == nil {
		params = &ListModelQualityJobDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelQualityJobDefinitions", params, optFns, c.addOperationListModelQualityJobDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelQualityJobDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelQualityJobDefinitionsInput struct {

	// A filter that returns only model quality monitoring job definitions created
	// after the specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only model quality monitoring job definitions created
	// before the specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only model quality monitoring job definitions that are
	// associated with the specified endpoint.
	EndpointName *string

	// The maximum number of results to return in a call to
	// ListModelQualityJobDefinitions.
	MaxResults *int32

	// A string in the transform job name. This filter returns only model quality
	// monitoring job definitions whose name contains the specified string.
	NameContains *string

	// If the result of the previous ListModelQualityJobDefinitions request was
	// truncated, the response includes a NextToken. To retrieve the next set of model
	// quality monitoring job definitions, use the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.MonitoringJobDefinitionSortKey

	// The sort order for results. The default is Descending.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListModelQualityJobDefinitionsOutput struct {

	// A list of summaries of model quality monitoring job definitions.
	//
	// This member is required.
	JobDefinitionSummaries []types.MonitoringJobDefinitionSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of model quality monitoring job definitions, use it in the next
	// request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelQualityJobDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModelQualityJobDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModelQualityJobDefinitions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelQualityJobDefinitions(options.Region), middleware.Before); err != nil {
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

// ListModelQualityJobDefinitionsAPIClient is a client that implements the
// ListModelQualityJobDefinitions operation.
type ListModelQualityJobDefinitionsAPIClient interface {
	ListModelQualityJobDefinitions(context.Context, *ListModelQualityJobDefinitionsInput, ...func(*Options)) (*ListModelQualityJobDefinitionsOutput, error)
}

var _ ListModelQualityJobDefinitionsAPIClient = (*Client)(nil)

// ListModelQualityJobDefinitionsPaginatorOptions is the paginator options for
// ListModelQualityJobDefinitions
type ListModelQualityJobDefinitionsPaginatorOptions struct {
	// The maximum number of results to return in a call to
	// ListModelQualityJobDefinitions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelQualityJobDefinitionsPaginator is a paginator for
// ListModelQualityJobDefinitions
type ListModelQualityJobDefinitionsPaginator struct {
	options   ListModelQualityJobDefinitionsPaginatorOptions
	client    ListModelQualityJobDefinitionsAPIClient
	params    *ListModelQualityJobDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListModelQualityJobDefinitionsPaginator returns a new
// ListModelQualityJobDefinitionsPaginator
func NewListModelQualityJobDefinitionsPaginator(client ListModelQualityJobDefinitionsAPIClient, params *ListModelQualityJobDefinitionsInput, optFns ...func(*ListModelQualityJobDefinitionsPaginatorOptions)) *ListModelQualityJobDefinitionsPaginator {
	if params == nil {
		params = &ListModelQualityJobDefinitionsInput{}
	}

	options := ListModelQualityJobDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelQualityJobDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelQualityJobDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelQualityJobDefinitions page.
func (p *ListModelQualityJobDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelQualityJobDefinitionsOutput, error) {
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

	result, err := p.client.ListModelQualityJobDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModelQualityJobDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModelQualityJobDefinitions",
	}
}
