// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts processing a stream processor. You create a stream processor by calling
// CreateStreamProcessor. To tell StartStreamProcessor which stream processor to
// start, use the value of the Name field specified in the call to
// CreateStreamProcessor. If you are using a label detection stream processor to
// detect labels, you need to provide a Start selector and a Stop selector to
// determine the length of the stream processing time.
func (c *Client) StartStreamProcessor(ctx context.Context, params *StartStreamProcessorInput, optFns ...func(*Options)) (*StartStreamProcessorOutput, error) {
	if params == nil {
		params = &StartStreamProcessorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartStreamProcessor", params, optFns, c.addOperationStartStreamProcessorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartStreamProcessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartStreamProcessorInput struct {

	// The name of the stream processor to start processing.
	//
	// This member is required.
	Name *string

	// Specifies the starting point in the Kinesis stream to start processing. You can
	// use the producer timestamp or the fragment number. For more information, see
	// Fragment
	// (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_reader_Fragment.html).
	// This is a required parameter for label detection stream processors and should
	// not be used to start a face search stream processor.
	StartSelector *types.StreamProcessingStartSelector

	// Specifies when to stop processing the stream. You can specify a maximum amount
	// of time to process the video. This is a required parameter for label detection
	// stream processors and should not be used to start a face search stream
	// processor.
	StopSelector *types.StreamProcessingStopSelector

	noSmithyDocumentSerde
}

type StartStreamProcessorOutput struct {

	// A unique identifier for the stream processing session.
	SessionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartStreamProcessorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartStreamProcessor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartStreamProcessor{}, middleware.After)
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
	if err = addOpStartStreamProcessorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartStreamProcessor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartStreamProcessor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartStreamProcessor",
	}
}
