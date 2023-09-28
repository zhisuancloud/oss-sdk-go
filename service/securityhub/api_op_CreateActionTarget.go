// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a custom action target in Security Hub. You can use custom actions on
// findings and insights in Security Hub to trigger target actions in Amazon
// CloudWatch Events.
func (c *Client) CreateActionTarget(ctx context.Context, params *CreateActionTargetInput, optFns ...func(*Options)) (*CreateActionTargetOutput, error) {
	if params == nil {
		params = &CreateActionTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateActionTarget", params, optFns, c.addOperationCreateActionTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateActionTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateActionTargetInput struct {

	// The description for the custom action target.
	//
	// This member is required.
	Description *string

	// The ID for the custom action target. Can contain up to 20 alphanumeric
	// characters.
	//
	// This member is required.
	Id *string

	// The name of the custom action target. Can contain up to 20 characters.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateActionTargetOutput struct {

	// The ARN for the custom action target.
	//
	// This member is required.
	ActionTargetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateActionTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateActionTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateActionTarget{}, middleware.After)
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
	if err = addOpCreateActionTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateActionTarget(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateActionTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "CreateActionTarget",
	}
}