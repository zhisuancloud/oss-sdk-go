// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/lookoutvision/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the model packaging jobs created for an Amazon Lookout for Vision project.
// This operation requires permissions to perform the
// lookoutvision:ListModelPackagingJobs operation. For more information, see Using
// your Amazon Lookout for Vision model on an edge device in the Amazon Lookout for
// Vision Developer Guide.
func (c *Client) ListModelPackagingJobs(ctx context.Context, params *ListModelPackagingJobsInput, optFns ...func(*Options)) (*ListModelPackagingJobsOutput, error) {
	if params == nil {
		params = &ListModelPackagingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelPackagingJobs", params, optFns, c.addOperationListModelPackagingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelPackagingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelPackagingJobsInput struct {

	// The name of the project for which you want to list the model packaging jobs.
	//
	// This member is required.
	ProjectName *string

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	MaxResults *int32

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Lookout for Vision returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListModelPackagingJobsOutput struct {

	// A list of the model packaging jobs created for the specified Amazon Lookout for
	// Vision project.
	ModelPackagingJobs []types.ModelPackagingJobMetadata

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Lookout for Vision returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelPackagingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListModelPackagingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListModelPackagingJobs{}, middleware.After)
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
	if err = addOpListModelPackagingJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelPackagingJobs(options.Region), middleware.Before); err != nil {
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

// ListModelPackagingJobsAPIClient is a client that implements the
// ListModelPackagingJobs operation.
type ListModelPackagingJobsAPIClient interface {
	ListModelPackagingJobs(context.Context, *ListModelPackagingJobsInput, ...func(*Options)) (*ListModelPackagingJobsOutput, error)
}

var _ ListModelPackagingJobsAPIClient = (*Client)(nil)

// ListModelPackagingJobsPaginatorOptions is the paginator options for
// ListModelPackagingJobs
type ListModelPackagingJobsPaginatorOptions struct {
	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelPackagingJobsPaginator is a paginator for ListModelPackagingJobs
type ListModelPackagingJobsPaginator struct {
	options   ListModelPackagingJobsPaginatorOptions
	client    ListModelPackagingJobsAPIClient
	params    *ListModelPackagingJobsInput
	nextToken *string
	firstPage bool
}

// NewListModelPackagingJobsPaginator returns a new ListModelPackagingJobsPaginator
func NewListModelPackagingJobsPaginator(client ListModelPackagingJobsAPIClient, params *ListModelPackagingJobsInput, optFns ...func(*ListModelPackagingJobsPaginatorOptions)) *ListModelPackagingJobsPaginator {
	if params == nil {
		params = &ListModelPackagingJobsInput{}
	}

	options := ListModelPackagingJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelPackagingJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelPackagingJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelPackagingJobs page.
func (p *ListModelPackagingJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelPackagingJobsOutput, error) {
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

	result, err := p.client.ListModelPackagingJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModelPackagingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutvision",
		OperationName: "ListModelPackagingJobs",
	}
}
