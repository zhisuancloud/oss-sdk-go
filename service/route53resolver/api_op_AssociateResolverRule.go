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

// Associates a Resolver rule with a VPC. When you associate a rule with a VPC,
// Resolver forwards all DNS queries for the domain name that is specified in the
// rule and that originate in the VPC. The queries are forwarded to the IP
// addresses for the DNS resolvers that are specified in the rule. For more
// information about rules, see CreateResolverRule
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html).
func (c *Client) AssociateResolverRule(ctx context.Context, params *AssociateResolverRuleInput, optFns ...func(*Options)) (*AssociateResolverRuleOutput, error) {
	if params == nil {
		params = &AssociateResolverRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateResolverRule", params, optFns, c.addOperationAssociateResolverRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateResolverRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResolverRuleInput struct {

	// The ID of the Resolver rule that you want to associate with the VPC. To list the
	// existing Resolver rules, use ListResolverRules
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRules.html).
	//
	// This member is required.
	ResolverRuleId *string

	// The ID of the VPC that you want to associate the Resolver rule with.
	//
	// This member is required.
	VPCId *string

	// A name for the association that you're creating between a Resolver rule and a
	// VPC.
	Name *string

	noSmithyDocumentSerde
}

type AssociateResolverRuleOutput struct {

	// Information about the AssociateResolverRule request, including the status of the
	// request.
	ResolverRuleAssociation *types.ResolverRuleAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateResolverRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateResolverRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateResolverRule{}, middleware.After)
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
	if err = addOpAssociateResolverRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResolverRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateResolverRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "AssociateResolverRule",
	}
}
