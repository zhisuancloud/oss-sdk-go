// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a temporary URL to start an AppStream 2.0 streaming session for the
// specified user. A streaming URL enables application streaming to be tested
// without user setup.
func (c *Client) CreateStreamingURL(ctx context.Context, params *CreateStreamingURLInput, optFns ...func(*Options)) (*CreateStreamingURLOutput, error) {
	if params == nil {
		params = &CreateStreamingURLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStreamingURL", params, optFns, c.addOperationCreateStreamingURLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamingURLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStreamingURLInput struct {

	// The name of the fleet.
	//
	// This member is required.
	FleetName *string

	// The name of the stack.
	//
	// This member is required.
	StackName *string

	// The identifier of the user.
	//
	// This member is required.
	UserId *string

	// The name of the application to launch after the session starts. This is the name
	// that you specified as Name in the Image Assistant. If your fleet is enabled for
	// the Desktop stream view, you can also choose to launch directly to the operating
	// system desktop. To do so, specify Desktop.
	ApplicationId *string

	// The session context. For more information, see Session Context
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/managing-stacks-fleets.html#managing-stacks-fleets-parameters)
	// in the Amazon AppStream 2.0 Administration Guide.
	SessionContext *string

	// The time that the streaming URL will be valid, in seconds. Specify a value
	// between 1 and 604800 seconds. The default is 60 seconds.
	Validity *int64

	noSmithyDocumentSerde
}

type CreateStreamingURLOutput struct {

	// The elapsed time, in seconds after the Unix epoch, when this URL expires.
	Expires *time.Time

	// The URL to start the AppStream 2.0 streaming session.
	StreamingURL *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStreamingURLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStreamingURL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStreamingURL{}, middleware.After)
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
	if err = addOpCreateStreamingURLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamingURL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStreamingURL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateStreamingURL",
	}
}
