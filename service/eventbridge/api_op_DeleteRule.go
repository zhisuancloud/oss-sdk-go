// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified rule. Before you can delete the rule, you must remove all
// targets, using RemoveTargets
// (https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_RemoveTargets.html).
// When you delete a rule, incoming events might continue to match to the deleted
// rule. Allow a short period of time for changes to take effect. If you call
// delete rule multiple times for the same rule, all calls will succeed. When you
// call delete rule for a non-existent custom eventbus, ResourceNotFoundException
// is returned. Managed rules are rules created and managed by another Amazon Web
// Services service on your behalf. These rules are created by those other Amazon
// Web Services services to support functionality in those services. You can delete
// these rules using the Force option, but you should do so only if you are sure
// the other service is not still using that rule.
func (c *Client) DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error) {
	if params == nil {
		params = &DeleteRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRule", params, optFns, c.addOperationDeleteRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRuleInput struct {

	// The name of the rule.
	//
	// This member is required.
	Name *string

	// The name or ARN of the event bus associated with the rule. If you omit this, the
	// default event bus is used.
	EventBusName *string

	// If this is a managed rule, created by an Amazon Web Services service on your
	// behalf, you must specify Force as True to delete the rule. This parameter is
	// ignored for rules that are not managed rules. You can check whether a rule is a
	// managed rule by using DescribeRule or ListRules and checking the ManagedBy field
	// of the response.
	Force bool

	noSmithyDocumentSerde
}

type DeleteRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRule{}, middleware.After)
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
	if err = addOpDeleteRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DeleteRule",
	}
}
