// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the requester's subscriptions. Each call returns a limited
// list of subscriptions, up to 100. If there are more subscriptions, a NextToken
// is also returned. Use the NextToken parameter in a new ListSubscriptions call to
// get further results. This action is throttled at 30 transactions per second
// (TPS).
func (c *Client) ListSubscriptions(ctx context.Context, params *ListSubscriptionsInput, optFns ...func(*Options)) (*ListSubscriptionsOutput, error) {
	if params == nil {
		params = &ListSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubscriptions", params, optFns, c.addOperationListSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListSubscriptions action.
type ListSubscriptionsInput struct {

	// Token returned by the previous ListSubscriptions request.
	NextToken *string

	noSmithyDocumentSerde
}

// Response for ListSubscriptions action
type ListSubscriptionsOutput struct {

	// Token to pass along to the next ListSubscriptions request. This element is
	// returned if there are more subscriptions to retrieve.
	NextToken *string

	// A list of subscriptions.
	Subscriptions []types.Subscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListSubscriptions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptions(options.Region), middleware.Before); err != nil {
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

// ListSubscriptionsAPIClient is a client that implements the ListSubscriptions
// operation.
type ListSubscriptionsAPIClient interface {
	ListSubscriptions(context.Context, *ListSubscriptionsInput, ...func(*Options)) (*ListSubscriptionsOutput, error)
}

var _ ListSubscriptionsAPIClient = (*Client)(nil)

// ListSubscriptionsPaginatorOptions is the paginator options for ListSubscriptions
type ListSubscriptionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubscriptionsPaginator is a paginator for ListSubscriptions
type ListSubscriptionsPaginator struct {
	options   ListSubscriptionsPaginatorOptions
	client    ListSubscriptionsAPIClient
	params    *ListSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewListSubscriptionsPaginator returns a new ListSubscriptionsPaginator
func NewListSubscriptionsPaginator(client ListSubscriptionsAPIClient, params *ListSubscriptionsInput, optFns ...func(*ListSubscriptionsPaginatorOptions)) *ListSubscriptionsPaginator {
	if params == nil {
		params = &ListSubscriptionsInput{}
	}

	options := ListSubscriptionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSubscriptions page.
func (p *ListSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubscriptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListSubscriptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListSubscriptions",
	}
}
