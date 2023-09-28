// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a specified resiliency policy for an AWS Resilience Hub application.
// The returned policy object includes creation time, data location constraints,
// the Amazon Resource Name (ARN) for the policy, tags, tier, and more.
func (c *Client) DescribeResiliencyPolicy(ctx context.Context, params *DescribeResiliencyPolicyInput, optFns ...func(*Options)) (*DescribeResiliencyPolicyOutput, error) {
	if params == nil {
		params = &DescribeResiliencyPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeResiliencyPolicy", params, optFns, c.addOperationDescribeResiliencyPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeResiliencyPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeResiliencyPolicyInput struct {

	// The Amazon Resource Name (ARN) of the resiliency policy. The format for this ARN
	// is: arn:partition:resiliencehub:region:account:resiliency-policy/policy-id. For
	// more information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	PolicyArn *string

	noSmithyDocumentSerde
}

type DescribeResiliencyPolicyOutput struct {

	// Information about the specific resiliency policy, returned as an object. This
	// object includes creation time, data location constraints, its name, description,
	// tags, the recovery time objective (RTO) and recovery point objective (RPO) in
	// seconds, and more.
	//
	// This member is required.
	Policy *types.ResiliencyPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeResiliencyPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeResiliencyPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeResiliencyPolicy{}, middleware.After)
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
	if err = addOpDescribeResiliencyPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeResiliencyPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeResiliencyPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "DescribeResiliencyPolicy",
	}
}
