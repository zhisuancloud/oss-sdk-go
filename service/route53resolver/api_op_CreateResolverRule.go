// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For DNS queries that originate in your VPCs, specifies which Resolver endpoint
// the queries pass through, one domain name that you want to forward to your
// network, and the IP addresses of the DNS resolvers in your network.
func (c *Client) CreateResolverRule(ctx context.Context, params *CreateResolverRuleInput, optFns ...func(*Options)) (*CreateResolverRuleOutput, error) {
	if params == nil {
		params = &CreateResolverRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResolverRule", params, optFns, c.addOperationCreateResolverRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResolverRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResolverRuleInput struct {

	// A unique string that identifies the request and that allows failed requests to
	// be retried without the risk of running the operation twice. CreatorRequestId can
	// be any unique string, for example, a date/time stamp.
	//
	// This member is required.
	CreatorRequestId *string

	// DNS queries for this domain name are forwarded to the IP addresses that you
	// specify in TargetIps. If a query matches multiple Resolver rules (example.com
	// and www.example.com), outbound DNS queries are routed using the Resolver rule
	// that contains the most specific domain name (www.example.com).
	//
	// This member is required.
	DomainName *string

	// When you want to forward DNS queries for specified domain name to resolvers on
	// your network, specify FORWARD. When you have a forwarding rule to forward DNS
	// queries for a domain to your network and you want Resolver to process queries
	// for a subdomain of that domain, specify SYSTEM. For example, to forward DNS
	// queries for example.com to resolvers on your network, you create a rule and
	// specify FORWARD for RuleType. To then have Resolver process queries for
	// apex.example.com, you create a rule and specify SYSTEM for RuleType. Currently,
	// only Resolver can create rules that have a value of RECURSIVE for RuleType.
	//
	// This member is required.
	RuleType types.RuleTypeOption

	// A friendly name that lets you easily find a rule in the Resolver dashboard in
	// the Route 53 console.
	Name *string

	// The ID of the outbound Resolver endpoint that you want to use to route DNS
	// queries to the IP addresses that you specify in TargetIps.
	ResolverEndpointId *string

	// A list of the tag keys and values that you want to associate with the endpoint.
	Tags []types.Tag

	// The IPs that you want Resolver to forward DNS queries to. You can specify only
	// IPv4 addresses. Separate IP addresses with a space. TargetIps is available only
	// when the value of Rule type is FORWARD.
	TargetIps []types.TargetAddress

	noSmithyDocumentSerde
}

type CreateResolverRuleOutput struct {

	// Information about the CreateResolverRule request, including the status of the
	// request.
	ResolverRule *types.ResolverRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResolverRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateResolverRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateResolverRule{}, middleware.After)
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
	if err = addOpCreateResolverRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResolverRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResolverRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "CreateResolverRule",
	}
}
