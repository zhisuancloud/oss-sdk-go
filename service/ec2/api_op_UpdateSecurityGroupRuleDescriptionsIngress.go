// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the description of an ingress (inbound) security group rule. You can
// replace an existing description, or add a description to a rule that did not
// have one previously. You can remove a description for a security group rule by
// omitting the description parameter in the request.
func (c *Client) UpdateSecurityGroupRuleDescriptionsIngress(ctx context.Context, params *UpdateSecurityGroupRuleDescriptionsIngressInput, optFns ...func(*Options)) (*UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	if params == nil {
		params = &UpdateSecurityGroupRuleDescriptionsIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSecurityGroupRuleDescriptionsIngress", params, optFns, c.addOperationUpdateSecurityGroupRuleDescriptionsIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSecurityGroupRuleDescriptionsIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSecurityGroupRuleDescriptionsIngressInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the security group. You must specify either the security group ID or
	// the security group name in the request. For security groups in a nondefault VPC,
	// you must specify the security group ID.
	GroupId *string

	// [EC2-Classic, default VPC] The name of the security group. You must specify
	// either the security group ID or the security group name in the request.
	GroupName *string

	// The IP permissions for the security group rule. You must specify either IP
	// permissions or a description.
	IpPermissions []types.IpPermission

	// [VPC only] The description for the ingress security group rules. You must
	// specify either a description or IP permissions.
	SecurityGroupRuleDescriptions []types.SecurityGroupRuleDescription

	noSmithyDocumentSerde
}

type UpdateSecurityGroupRuleDescriptionsIngressOutput struct {

	// Returns true if the request succeeds; otherwise, returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSecurityGroupRuleDescriptionsIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpUpdateSecurityGroupRuleDescriptionsIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpUpdateSecurityGroupRuleDescriptionsIngress{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSecurityGroupRuleDescriptionsIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSecurityGroupRuleDescriptionsIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "UpdateSecurityGroupRuleDescriptionsIngress",
	}
}
