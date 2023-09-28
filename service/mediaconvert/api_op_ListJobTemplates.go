// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a JSON array of up to twenty of your job templates. This will return
// the templates themselves, not just a list of them. To retrieve the next twenty
// templates, use the nextToken string returned with the array
func (c *Client) ListJobTemplates(ctx context.Context, params *ListJobTemplatesInput, optFns ...func(*Options)) (*ListJobTemplatesOutput, error) {
	if params == nil {
		params = &ListJobTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobTemplates", params, optFns, c.addOperationListJobTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobTemplatesInput struct {

	// Optionally, specify a job template category to limit responses to only job
	// templates from that category.
	Category *string

	// Optional. When you request a list of job templates, you can choose to list them
	// alphabetically by NAME or chronologically by CREATION_DATE. If you don't
	// specify, the service will list them by name.
	ListBy types.JobTemplateListBy

	// Optional. Number of job templates, up to twenty, that will be returned at one
	// time.
	MaxResults int32

	// Use this string, provided with the response to a previous request, to request
	// the next batch of job templates.
	NextToken *string

	// Optional. When you request lists of resources, you can specify whether they are
	// sorted in ASCENDING or DESCENDING order. Default varies by resource.
	Order types.Order

	noSmithyDocumentSerde
}

type ListJobTemplatesOutput struct {

	// List of Job templates.
	JobTemplates []types.JobTemplate

	// Use this string to request the next batch of job templates.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJobTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobTemplates(options.Region), middleware.Before); err != nil {
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

// ListJobTemplatesAPIClient is a client that implements the ListJobTemplates
// operation.
type ListJobTemplatesAPIClient interface {
	ListJobTemplates(context.Context, *ListJobTemplatesInput, ...func(*Options)) (*ListJobTemplatesOutput, error)
}

var _ ListJobTemplatesAPIClient = (*Client)(nil)

// ListJobTemplatesPaginatorOptions is the paginator options for ListJobTemplates
type ListJobTemplatesPaginatorOptions struct {
	// Optional. Number of job templates, up to twenty, that will be returned at one
	// time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobTemplatesPaginator is a paginator for ListJobTemplates
type ListJobTemplatesPaginator struct {
	options   ListJobTemplatesPaginatorOptions
	client    ListJobTemplatesAPIClient
	params    *ListJobTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListJobTemplatesPaginator returns a new ListJobTemplatesPaginator
func NewListJobTemplatesPaginator(client ListJobTemplatesAPIClient, params *ListJobTemplatesInput, optFns ...func(*ListJobTemplatesPaginatorOptions)) *ListJobTemplatesPaginator {
	if params == nil {
		params = &ListJobTemplatesInput{}
	}

	options := ListJobTemplatesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJobTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJobTemplates page.
func (p *ListJobTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobTemplatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListJobTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "ListJobTemplates",
	}
}