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

// Associates an Amazon VPC with a specified query logging configuration. Route 53
// Resolver logs DNS queries that originate in all of the Amazon VPCs that are
// associated with a specified query logging configuration. To associate more than
// one VPC with a configuration, submit one AssociateResolverQueryLogConfig request
// for each VPC. The VPCs that you associate with a query logging configuration
// must be in the same Region as the configuration. To remove a VPC from a query
// logging configuration, see DisassociateResolverQueryLogConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverQueryLogConfig.html).
func (c *Client) AssociateResolverQueryLogConfig(ctx context.Context, params *AssociateResolverQueryLogConfigInput, optFns ...func(*Options)) (*AssociateResolverQueryLogConfigOutput, error) {
	if params == nil {
		params = &AssociateResolverQueryLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateResolverQueryLogConfig", params, optFns, c.addOperationAssociateResolverQueryLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateResolverQueryLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResolverQueryLogConfigInput struct {

	// The ID of the query logging configuration that you want to associate a VPC with.
	//
	// This member is required.
	ResolverQueryLogConfigId *string

	// The ID of an Amazon VPC that you want this query logging configuration to log
	// queries for. The VPCs and the query logging configuration must be in the same
	// Region.
	//
	// This member is required.
	ResourceId *string

	noSmithyDocumentSerde
}

type AssociateResolverQueryLogConfigOutput struct {

	// A complex type that contains settings for a specified association between an
	// Amazon VPC and a query logging configuration.
	ResolverQueryLogConfigAssociation *types.ResolverQueryLogConfigAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateResolverQueryLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateResolverQueryLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateResolverQueryLogConfig{}, middleware.After)
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
	if err = addOpAssociateResolverQueryLogConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResolverQueryLogConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateResolverQueryLogConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "AssociateResolverQueryLogConfig",
	}
}
