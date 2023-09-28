// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing custom verification email template. For more information
// about custom verification email templates, see Using Custom Verification Email
// Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-verify-address-custom.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) DeleteCustomVerificationEmailTemplate(ctx context.Context, params *DeleteCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*DeleteCustomVerificationEmailTemplateOutput, error) {
	if params == nil {
		params = &DeleteCustomVerificationEmailTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomVerificationEmailTemplate", params, optFns, c.addOperationDeleteCustomVerificationEmailTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomVerificationEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete an existing custom verification email template.
type DeleteCustomVerificationEmailTemplateInput struct {

	// The name of the custom verification email template that you want to delete.
	//
	// This member is required.
	TemplateName *string

	noSmithyDocumentSerde
}

// If the action is successful, the service sends back an HTTP 200 response with an
// empty HTTP body.
type DeleteCustomVerificationEmailTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCustomVerificationEmailTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteCustomVerificationEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteCustomVerificationEmailTemplate{}, middleware.After)
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
	if err = addOpDeleteCustomVerificationEmailTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomVerificationEmailTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCustomVerificationEmailTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteCustomVerificationEmailTemplate",
	}
}
