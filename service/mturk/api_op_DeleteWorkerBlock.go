// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The DeleteWorkerBlock operation allows you to reinstate a blocked Worker to work
// on your HITs. This operation reverses the effects of the CreateWorkerBlock
// operation. You need the Worker ID to use this operation. If the Worker ID is
// missing or invalid, this operation fails and returns the message “WorkerId is
// invalid.” If the specified Worker is not blocked, this operation returns
// successfully.
func (c *Client) DeleteWorkerBlock(ctx context.Context, params *DeleteWorkerBlockInput, optFns ...func(*Options)) (*DeleteWorkerBlockOutput, error) {
	if params == nil {
		params = &DeleteWorkerBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteWorkerBlock", params, optFns, c.addOperationDeleteWorkerBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteWorkerBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWorkerBlockInput struct {

	// The ID of the Worker to unblock.
	//
	// This member is required.
	WorkerId *string

	// A message that explains the reason for unblocking the Worker. The Worker does
	// not see this message.
	Reason *string

	noSmithyDocumentSerde
}

type DeleteWorkerBlockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteWorkerBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteWorkerBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteWorkerBlock{}, middleware.After)
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
	if err = addOpDeleteWorkerBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWorkerBlock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteWorkerBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "DeleteWorkerBlock",
	}
}
