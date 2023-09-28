// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes all versions of a bot, including the Draft version. To delete a specific
// version, use the DeleteBotVersion operation. When you delete a bot, all of the
// resources contained in the bot are also deleted. Deleting a bot removes all
// locales, intents, slot, and slot types defined for the bot. If a bot has an
// alias, the DeleteBot operation returns a ResourceInUseException exception. If
// you want to delete the bot and the alias, set the skipResourceInUseCheck
// parameter to true.
func (c *Client) DeleteBot(ctx context.Context, params *DeleteBotInput, optFns ...func(*Options)) (*DeleteBotOutput, error) {
	if params == nil {
		params = &DeleteBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBot", params, optFns, c.addOperationDeleteBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBotInput struct {

	// The identifier of the bot to delete.
	//
	// This member is required.
	BotId *string

	// When true, Amazon Lex doesn't check to see if another resource, such as an
	// alias, is using the bot before it is deleted.
	SkipResourceInUseCheck bool

	noSmithyDocumentSerde
}

type DeleteBotOutput struct {

	// The unique identifier of the bot that Amazon Lex is deleting.
	BotId *string

	// The current status of the bot. The status is Deleting while the bot and its
	// associated resources are being deleted.
	BotStatus types.BotStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBot{}, middleware.After)
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
	if err = addOpDeleteBotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteBot",
	}
}
