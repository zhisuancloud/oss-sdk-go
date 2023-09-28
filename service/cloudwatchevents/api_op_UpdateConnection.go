// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cloudwatchevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates settings for a connection.
func (c *Client) UpdateConnection(ctx context.Context, params *UpdateConnectionInput, optFns ...func(*Options)) (*UpdateConnectionOutput, error) {
	if params == nil {
		params = &UpdateConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConnection", params, optFns, c.addOperationUpdateConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConnectionInput struct {

	// The name of the connection to update.
	//
	// This member is required.
	Name *string

	// The authorization parameters to use for the connection.
	AuthParameters *types.UpdateConnectionAuthRequestParameters

	// The type of authorization to use for the connection.
	AuthorizationType types.ConnectionAuthorizationType

	// A description for the connection.
	Description *string

	noSmithyDocumentSerde
}

type UpdateConnectionOutput struct {

	// The ARN of the connection that was updated.
	ConnectionArn *string

	// The state of the connection that was updated.
	ConnectionState types.ConnectionState

	// A time stamp for the time that the connection was created.
	CreationTime *time.Time

	// A time stamp for the time that the connection was last authorized.
	LastAuthorizedTime *time.Time

	// A time stamp for the time that the connection was last modified.
	LastModifiedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateConnection{}, middleware.After)
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
	if err = addOpUpdateConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "UpdateConnection",
	}
}
