// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns summary information about the results of a stack set operation.
func (c *Client) ListStackSetOperationResults(ctx context.Context, params *ListStackSetOperationResultsInput, optFns ...func(*Options)) (*ListStackSetOperationResultsOutput, error) {
	if params == nil {
		params = &ListStackSetOperationResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStackSetOperationResults", params, optFns, c.addOperationListStackSetOperationResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStackSetOperationResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStackSetOperationResultsInput struct {

	// The ID of the stack set operation.
	//
	// This member is required.
	OperationId *string

	// The name or unique ID of the stack set that you want to get operation results
	// for.
	//
	// This member is required.
	StackSetName *string

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account. By default, SELF is specified. Use SELF for
	// stack sets with self-managed permissions.
	//
	// * If you are signed in to the
	// management account, specify SELF.
	//
	// * If you are signed in to a delegated
	// administrator account, specify DELEGATED_ADMIN. Your Amazon Web Services account
	// must be registered as a delegated administrator in the management account. For
	// more information, see Register a delegated administrator
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	// in the CloudFormation User Guide.
	CallAs types.CallAs

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32

	// If the previous request didn't return all the remaining results, the response
	// object's NextToken parameter value is set to a token. To retrieve the next set
	// of results, call ListStackSetOperationResults again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStackSetOperationResultsOutput struct {

	// If the request doesn't return all results, NextToken is set to a token. To
	// retrieve the next set of results, call ListOperationResults again and assign
	// that token to the request object's NextToken parameter. If there are no
	// remaining results, NextToken is set to null.
	NextToken *string

	// A list of StackSetOperationResultSummary structures that contain information
	// about the specified operation results, for accounts and Amazon Web Services
	// Regions that are included in the operation.
	Summaries []types.StackSetOperationResultSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStackSetOperationResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListStackSetOperationResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListStackSetOperationResults{}, middleware.After)
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
	if err = addOpListStackSetOperationResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStackSetOperationResults(options.Region), middleware.Before); err != nil {
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

// ListStackSetOperationResultsAPIClient is a client that implements the
// ListStackSetOperationResults operation.
type ListStackSetOperationResultsAPIClient interface {
	ListStackSetOperationResults(context.Context, *ListStackSetOperationResultsInput, ...func(*Options)) (*ListStackSetOperationResultsOutput, error)
}

var _ ListStackSetOperationResultsAPIClient = (*Client)(nil)

// ListStackSetOperationResultsPaginatorOptions is the paginator options for
// ListStackSetOperationResults
type ListStackSetOperationResultsPaginatorOptions struct {
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStackSetOperationResultsPaginator is a paginator for
// ListStackSetOperationResults
type ListStackSetOperationResultsPaginator struct {
	options   ListStackSetOperationResultsPaginatorOptions
	client    ListStackSetOperationResultsAPIClient
	params    *ListStackSetOperationResultsInput
	nextToken *string
	firstPage bool
}

// NewListStackSetOperationResultsPaginator returns a new
// ListStackSetOperationResultsPaginator
func NewListStackSetOperationResultsPaginator(client ListStackSetOperationResultsAPIClient, params *ListStackSetOperationResultsInput, optFns ...func(*ListStackSetOperationResultsPaginatorOptions)) *ListStackSetOperationResultsPaginator {
	if params == nil {
		params = &ListStackSetOperationResultsInput{}
	}

	options := ListStackSetOperationResultsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStackSetOperationResultsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStackSetOperationResultsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStackSetOperationResults page.
func (p *ListStackSetOperationResultsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStackSetOperationResultsOutput, error) {
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

	result, err := p.client.ListStackSetOperationResults(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStackSetOperationResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListStackSetOperationResults",
	}
}
