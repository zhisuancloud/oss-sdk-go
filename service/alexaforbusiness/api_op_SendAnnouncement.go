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

// Triggers an asynchronous flow to send text, SSML, or audio announcements to
// rooms that are identified by a search or filter.
func (c *Client) SendAnnouncement(ctx context.Context, params *SendAnnouncementInput, optFns ...func(*Options)) (*SendAnnouncementOutput, error) {
	if params == nil {
		params = &SendAnnouncementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendAnnouncement", params, optFns, c.addOperationSendAnnouncementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendAnnouncementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendAnnouncementInput struct {

	// The unique, user-specified identifier for the request that ensures idempotency.
	//
	// This member is required.
	ClientRequestToken *string

	// The announcement content. This can contain only one of the three possible
	// announcement types (text, SSML or audio).
	//
	// This member is required.
	Content *types.Content

	// The filters to use to send an announcement to a specified list of rooms. The
	// supported filter keys are RoomName, ProfileName, RoomArn, and ProfileArn. To
	// send to all rooms, specify an empty RoomFilters list.
	//
	// This member is required.
	RoomFilters []types.Filter

	// The time to live for an announcement. Default is 300. If delivery doesn't occur
	// within this time, the announcement is not delivered.
	TimeToLiveInSeconds *int32

	noSmithyDocumentSerde
}

type SendAnnouncementOutput struct {

	// The identifier of the announcement.
	AnnouncementArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendAnnouncementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendAnnouncement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendAnnouncement{}, middleware.After)
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
	if err = addIdempotencyToken_opSendAnnouncementMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendAnnouncementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendAnnouncement(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSendAnnouncement struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendAnnouncement) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendAnnouncement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendAnnouncementInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendAnnouncementInput ")
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
func addIdempotencyToken_opSendAnnouncementMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendAnnouncement{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendAnnouncement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SendAnnouncement",
	}
}
