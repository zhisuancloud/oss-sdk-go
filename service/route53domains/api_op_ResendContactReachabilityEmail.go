// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For operations that require confirmation that the email address for the
// registrant contact is valid, such as registering a new domain, this operation
// resends the confirmation email to the current email address for the registrant
// contact.
func (c *Client) ResendContactReachabilityEmail(ctx context.Context, params *ResendContactReachabilityEmailInput, optFns ...func(*Options)) (*ResendContactReachabilityEmailOutput, error) {
	if params == nil {
		params = &ResendContactReachabilityEmailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResendContactReachabilityEmail", params, optFns, c.addOperationResendContactReachabilityEmailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResendContactReachabilityEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResendContactReachabilityEmailInput struct {

	// The name of the domain for which you want Route 53 to resend a confirmation
	// email to the registrant contact.
	DomainName *string

	noSmithyDocumentSerde
}

type ResendContactReachabilityEmailOutput struct {

	// The domain name for which you requested a confirmation email.
	DomainName *string

	// The email address for the registrant contact at the time that we sent the
	// verification email.
	EmailAddress *string

	// True if the email address for the registrant contact has already been verified,
	// and false otherwise. If the email address has already been verified, we don't
	// send another confirmation email.
	IsAlreadyVerified *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResendContactReachabilityEmailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResendContactReachabilityEmail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResendContactReachabilityEmail{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResendContactReachabilityEmail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResendContactReachabilityEmail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "ResendContactReachabilityEmail",
	}
}
