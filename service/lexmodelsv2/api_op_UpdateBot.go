// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the configuration of an existing bot.
func (c *Client) UpdateBot(ctx context.Context, params *UpdateBotInput, optFns ...func(*Options)) (*UpdateBotOutput, error) {
	if params == nil {
		params = &UpdateBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBot", params, optFns, c.addOperationUpdateBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBotInput struct {

	// The unique identifier of the bot to update. This identifier is returned by the
	// CreateBot (https://docs.aws.amazon.com/lexv2/latest/dg/API_CreateBot.html)
	// operation.
	//
	// This member is required.
	BotId *string

	// The new name of the bot. The name must be unique in the account that creates the
	// bot.
	//
	// This member is required.
	BotName *string

	// Provides information on additional privacy protections Amazon Lex should use
	// with the bot's data.
	//
	// This member is required.
	DataPrivacy *types.DataPrivacy

	// The time, in seconds, that Amazon Lex should keep information about a user's
	// conversation with the bot. A user interaction remains active for the amount of
	// time specified. If no conversation occurs during this time, the session expires
	// and Amazon Lex deletes any data provided before the timeout. You can specify
	// between 60 (1 minute) and 86,400 (24 hours) seconds.
	//
	// This member is required.
	IdleSessionTTLInSeconds *int32

	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access the
	// bot.
	//
	// This member is required.
	RoleArn *string

	// A description of the bot.
	Description *string

	noSmithyDocumentSerde
}

type UpdateBotOutput struct {

	// The unique identifier of the bot that was updated.
	BotId *string

	// The name of the bot after the update.
	BotName *string

	// Shows the current status of the bot. The bot is first in the Creating status.
	// Once the bot is read for use, it changes to the Available status. After the bot
	// is created, you can use the DRAFT version of the bot.
	BotStatus types.BotStatus

	// A timestamp of the date and time that the bot was created.
	CreationDateTime *time.Time

	// The data privacy settings for the bot after the update.
	DataPrivacy *types.DataPrivacy

	// The description of the bot after the update.
	Description *string

	// The session timeout, in seconds, for the bot after the update.
	IdleSessionTTLInSeconds *int32

	// A timestamp of the date and time that the bot was last updated.
	LastUpdatedDateTime *time.Time

	// The Amazon Resource Name (ARN) of the IAM role used by the bot after the update.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBot{}, middleware.After)
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
	if err = addOpUpdateBotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "UpdateBot",
	}
}
