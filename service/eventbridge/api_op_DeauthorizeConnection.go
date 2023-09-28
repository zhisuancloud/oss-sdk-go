// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Removes all authorization parameters from the connection. This lets you remove
// the secret from the connection so you can reuse it without having to create a
// new connection.
func (c *Client) DeauthorizeConnection(ctx context.Context, params *DeauthorizeConnectionInput, optFns ...func(*Options)) (*DeauthorizeConnectionOutput, error) {
	if params == nil {
		params = &DeauthorizeConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeauthorizeConnection", params, optFns, c.addOperationDeauthorizeConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeauthorizeConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeauthorizeConnectionInput struct {

	// The name of the connection to remove authorization from.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeauthorizeConnectionOutput struct {

	// The ARN of the connection that authorization was removed from.
	ConnectionArn *string

	// The state of the connection.
	ConnectionState types.ConnectionState

	// A time stamp for the time that the connection was created.
	CreationTime *time.Time

	// A time stamp for the time that the connection was last authorized.
	LastAuthorizedTime *time.Time

	// A time stamp for the time that the connection was last updated.
	LastModifiedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeauthorizeConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeauthorizeConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeauthorizeConnection{}, middleware.After)
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
	if err = addOpDeauthorizeConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeauthorizeConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeauthorizeConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DeauthorizeConnection",
	}
}
