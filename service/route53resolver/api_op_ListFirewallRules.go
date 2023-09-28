// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the firewall rules that you have defined for the specified firewall
// rule group. DNS Firewall uses the rules in a rule group to filter DNS network
// traffic for a VPC. A single call might return only a partial list of the rules.
// For information, see MaxResults.
func (c *Client) ListFirewallRules(ctx context.Context, params *ListFirewallRulesInput, optFns ...func(*Options)) (*ListFirewallRulesOutput, error) {
	if params == nil {
		params = &ListFirewallRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFirewallRules", params, optFns, c.addOperationListFirewallRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFirewallRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFirewallRulesInput struct {

	// The unique identifier of the firewall rule group that you want to retrieve the
	// rules for.
	//
	// This member is required.
	FirewallRuleGroupId *string

	// Optional additional filter for the rules to retrieve. The action that DNS
	// Firewall should take on a DNS query when it matches one of the domains in the
	// rule's domain list:
	//
	// * ALLOW - Permit the request to go through.
	//
	// * ALERT -
	// Permit the request to go through but send an alert to the logs.
	//
	// * BLOCK -
	// Disallow the request. If this is specified, additional handling details are
	// provided in the rule's BlockResponse setting.
	Action types.Action

	// The maximum number of objects that you want Resolver to return for this request.
	// If more objects are available, in the response, Resolver provides a NextToken
	// value that you can use in a subsequent call to get the next batch of objects. If
	// you don't specify a value for MaxResults, Resolver returns up to 100 objects.
	MaxResults *int32

	// For the first call to this list request, omit this value. When you request a
	// list of objects, Resolver returns at most the number of objects specified in
	// MaxResults. If more objects are available for retrieval, Resolver returns a
	// NextToken value in the response. To retrieve the next batch of objects, use the
	// token that was returned for the prior request in your next request.
	NextToken *string

	// Optional additional filter for the rules to retrieve. The setting that
	// determines the processing order of the rules in a rule group. DNS Firewall
	// processes the rules in a rule group by order of priority, starting from the
	// lowest setting.
	Priority *int32

	noSmithyDocumentSerde
}

type ListFirewallRulesOutput struct {

	// A list of the rules that you have defined. This might be a partial list of the
	// firewall rules that you've defined. For information, see MaxResults.
	FirewallRules []types.FirewallRule

	// If objects are still available for retrieval, Resolver returns this token in the
	// response. To retrieve the next batch of objects, provide this token in your next
	// request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFirewallRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFirewallRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFirewallRules{}, middleware.After)
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
	if err = addOpListFirewallRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFirewallRules(options.Region), middleware.Before); err != nil {
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

// ListFirewallRulesAPIClient is a client that implements the ListFirewallRules
// operation.
type ListFirewallRulesAPIClient interface {
	ListFirewallRules(context.Context, *ListFirewallRulesInput, ...func(*Options)) (*ListFirewallRulesOutput, error)
}

var _ ListFirewallRulesAPIClient = (*Client)(nil)

// ListFirewallRulesPaginatorOptions is the paginator options for ListFirewallRules
type ListFirewallRulesPaginatorOptions struct {
	// The maximum number of objects that you want Resolver to return for this request.
	// If more objects are available, in the response, Resolver provides a NextToken
	// value that you can use in a subsequent call to get the next batch of objects. If
	// you don't specify a value for MaxResults, Resolver returns up to 100 objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFirewallRulesPaginator is a paginator for ListFirewallRules
type ListFirewallRulesPaginator struct {
	options   ListFirewallRulesPaginatorOptions
	client    ListFirewallRulesAPIClient
	params    *ListFirewallRulesInput
	nextToken *string
	firstPage bool
}

// NewListFirewallRulesPaginator returns a new ListFirewallRulesPaginator
func NewListFirewallRulesPaginator(client ListFirewallRulesAPIClient, params *ListFirewallRulesInput, optFns ...func(*ListFirewallRulesPaginatorOptions)) *ListFirewallRulesPaginator {
	if params == nil {
		params = &ListFirewallRulesInput{}
	}

	options := ListFirewallRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFirewallRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFirewallRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFirewallRules page.
func (p *ListFirewallRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFirewallRulesOutput, error) {
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

	result, err := p.client.ListFirewallRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFirewallRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListFirewallRules",
	}
}
