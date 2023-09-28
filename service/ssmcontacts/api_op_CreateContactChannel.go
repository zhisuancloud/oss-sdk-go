// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A contact channel is the method that Incident Manager uses to engage your
// contact.
func (c *Client) CreateContactChannel(ctx context.Context, params *CreateContactChannelInput, optFns ...func(*Options)) (*CreateContactChannelOutput, error) {
	if params == nil {
		params = &CreateContactChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContactChannel", params, optFns, c.addOperationCreateContactChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContactChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContactChannelInput struct {

	// The Amazon Resource Name (ARN) of the contact you are adding the contact channel
	// to.
	//
	// This member is required.
	ContactId *string

	// The details that Incident Manager uses when trying to engage the contact
	// channel. The format is dependent on the type of the contact channel. The
	// following are the expected formats:
	//
	// * SMS - '+' followed by the country code
	// and phone number
	//
	// * VOICE - '+' followed by the country code and phone number
	//
	// *
	// EMAIL - any standard email format
	//
	// This member is required.
	DeliveryAddress *types.ContactChannelAddress

	// The name of the contact channel.
	//
	// This member is required.
	Name *string

	// Incident Manager supports three types of contact channels:
	//
	// * SMS
	//
	// * VOICE
	//
	// *
	// EMAIL
	//
	// This member is required.
	Type types.ChannelType

	// If you want to activate the channel at a later time, you can choose to defer
	// activation. Incident Manager can't engage your contact channel until it has been
	// activated.
	DeferActivation *bool

	// A token ensuring that the operation is called only once with the specified
	// details.
	IdempotencyToken *string

	noSmithyDocumentSerde
}

type CreateContactChannelOutput struct {

	// The Amazon Resource Name (ARN) of the contact channel.
	//
	// This member is required.
	ContactChannelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContactChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContactChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContactChannel{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateContactChannelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateContactChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContactChannel(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateContactChannel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateContactChannel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateContactChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateContactChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateContactChannelInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateContactChannelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateContactChannel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateContactChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-contacts",
		OperationName: "CreateContactChannel",
	}
}
