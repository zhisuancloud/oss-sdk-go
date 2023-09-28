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

// Updates the settings that a bot has for a specific locale.
func (c *Client) UpdateBotLocale(ctx context.Context, params *UpdateBotLocaleInput, optFns ...func(*Options)) (*UpdateBotLocaleOutput, error) {
	if params == nil {
		params = &UpdateBotLocaleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBotLocale", params, optFns, c.addOperationUpdateBotLocaleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBotLocaleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBotLocaleInput struct {

	// The unique identifier of the bot that contains the locale.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the locale to be updated. The version can
	// only be the DRAFT version.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale to update. The string must match one
	// of the supported locales. For more information, see Supported languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	// The new confidence threshold where Amazon Lex inserts the AMAZON.FallbackIntent
	// and AMAZON.KendraSearchIntent intents in the list of possible intents for an
	// utterance.
	//
	// This member is required.
	NluIntentConfidenceThreshold *float64

	// The new description of the locale.
	Description *string

	// The new Amazon Polly voice Amazon Lex should use for voice interaction with the
	// user.
	VoiceSettings *types.VoiceSettings

	noSmithyDocumentSerde
}

type UpdateBotLocaleOutput struct {

	// The identifier of the bot that contains the updated locale.
	BotId *string

	// The current status of the locale. When the bot status is Built the locale is
	// ready for use.
	BotLocaleStatus types.BotLocaleStatus

	// The version of the bot that contains the updated locale.
	BotVersion *string

	// A timestamp of the date and time that the locale was created.
	CreationDateTime *time.Time

	// The updated description of the locale.
	Description *string

	// If the botLocaleStatus is Failed, the failureReasons field lists the errors that
	// occurred while building the bot.
	FailureReasons []string

	// A timestamp of the date and time that the locale was last updated.
	LastUpdatedDateTime *time.Time

	// The language and locale of the updated bot locale.
	LocaleId *string

	// The updated locale name for the locale.
	LocaleName *string

	// The updated confidence threshold for inserting the AMAZON.FallbackIntent and
	// AMAZON.KendraSearchIntent intents in the list of possible intents for an
	// utterance.
	NluIntentConfidenceThreshold *float64

	// Recommended actions to take to resolve an error in the failureReasons field.
	RecommendedActions []string

	// The updated Amazon Polly voice to use for voice interaction with the user.
	VoiceSettings *types.VoiceSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBotLocaleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBotLocale{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBotLocale{}, middleware.After)
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
	if err = addOpUpdateBotLocaleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBotLocale(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBotLocale(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "UpdateBotLocale",
	}
}
