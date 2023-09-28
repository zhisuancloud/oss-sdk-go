// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a new conference provider under the user's AWS account.
func (c *Client) CreateConferenceProvider(ctx context.Context, params *CreateConferenceProviderInput, optFns ...func(*Options)) (*CreateConferenceProviderOutput, error) {
	if params == nil {
		params = &CreateConferenceProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConferenceProvider", params, optFns, c.addOperationCreateConferenceProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConferenceProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConferenceProviderInput struct {

	// The name of the conference provider.
	//
	// This member is required.
	ConferenceProviderName *string

	// Represents a type within a list of predefined types.
	//
	// This member is required.
	ConferenceProviderType types.ConferenceProviderType

	// The meeting settings for the conference provider.
	//
	// This member is required.
	MeetingSetting *types.MeetingSetting

	// The request token of the client.
	ClientRequestToken *string

	// The IP endpoint and protocol for calling.
	IPDialIn *types.IPDialIn

	// The information for PSTN conferencing.
	PSTNDialIn *types.PSTNDialIn

	// The tags to be added to the specified resource. Do not provide system tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateConferenceProviderOutput struct {

	// The ARN of the newly-created conference provider.
	ConferenceProviderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConferenceProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConferenceProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConferenceProvider{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateConferenceProviderMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateConferenceProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConferenceProvider(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateConferenceProvider struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateConferenceProvider) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateConferenceProvider) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateConferenceProviderInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateConferenceProviderInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateConferenceProviderMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateConferenceProvider{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateConferenceProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "CreateConferenceProvider",
	}
}