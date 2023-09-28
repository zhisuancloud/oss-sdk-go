// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// This examples serializes a streaming blob shape with a required content length
// in the request body. In this example, no JSON document is synthesized because
// the payload is not a structure or a union type.
func (c *Client) StreamingTraitsRequireLength(ctx context.Context, params *StreamingTraitsRequireLengthInput, optFns ...func(*Options)) (*StreamingTraitsRequireLengthOutput, error) {
	if params == nil {
		params = &StreamingTraitsRequireLengthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StreamingTraitsRequireLength", params, optFns, c.addOperationStreamingTraitsRequireLengthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StreamingTraitsRequireLengthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StreamingTraitsRequireLengthInput struct {
	Blob io.Reader

	Foo *string

	noSmithyDocumentSerde
}

type StreamingTraitsRequireLengthOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStreamingTraitsRequireLengthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStreamingTraitsRequireLength{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStreamingTraitsRequireLength{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStreamingTraitsRequireLength(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStreamingTraitsRequireLength(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StreamingTraitsRequireLength",
	}
}
