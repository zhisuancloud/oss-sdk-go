// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new channel and an associated stream key to start streaming.
func (c *Client) CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) {
	if params == nil {
		params = &CreateChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannel", params, optFns, c.addOperationCreateChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChannelInput struct {

	// Whether the channel is private (enabled for playback authorization). Default:
	// false.
	Authorized bool

	// Channel latency mode. Use NORMAL to broadcast and deliver live video up to Full
	// HD. Use LOW for near-real-time interaction with viewers. (Note: In the Amazon
	// IVS console, LOW and NORMAL correspond to Ultra-low and Standard, respectively.)
	// Default: LOW.
	LatencyMode types.ChannelLatencyMode

	// Channel name.
	Name *string

	// Recording-configuration ARN. Default: "" (empty string, recording is disabled).
	RecordingConfigurationArn *string

	// Array of 1-50 maps, each of the form string:string (key:value). See Tagging
	// Amazon Web Services Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) for more
	// information, including restrictions that apply to tags and "Tag naming limits
	// and requirements"; Amazon IVS has no service-specific constraints beyond what is
	// documented there.
	Tags map[string]string

	// Channel type, which determines the allowable resolution and bitrate. If you
	// exceed the allowable resolution or bitrate, the stream probably will disconnect
	// immediately. Default: STANDARD. Valid values:
	//
	// * STANDARD: Video is transcoded:
	// multiple qualities are generated from the original input, to automatically give
	// viewers the best experience for their devices and network conditions.
	// Transcoding allows higher playback quality across a range of download speeds.
	// Resolution can be up to 1080p and bitrate can be up to 8.5 Mbps. Audio is
	// transcoded only for renditions 360p and below; above that, audio is passed
	// through. This is the default.
	//
	// * BASIC: Video is transmuxed: Amazon IVS delivers
	// the original input to viewers. The viewer’s video-quality choice is limited to
	// the original input. Resolution can be up to 1080p and bitrate can be up to 1.5
	// Mbps for 480p and up to 3.5 Mbps for resolutions between 480p and 1080p.
	Type types.ChannelType

	noSmithyDocumentSerde
}

type CreateChannelOutput struct {

	//
	Channel *types.Channel

	//
	StreamKey *types.StreamKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannel{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "CreateChannel",
	}
}
