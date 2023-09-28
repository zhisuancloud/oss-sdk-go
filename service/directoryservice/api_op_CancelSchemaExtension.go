// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels an in-progress schema extension to a Microsoft AD directory. Once a
// schema extension has started replicating to all domain controllers, the task can
// no longer be canceled. A schema extension can be canceled during any of the
// following states; Initializing, CreatingSnapshot, and UpdatingSchema.
func (c *Client) CancelSchemaExtension(ctx context.Context, params *CancelSchemaExtensionInput, optFns ...func(*Options)) (*CancelSchemaExtensionOutput, error) {
	if params == nil {
		params = &CancelSchemaExtensionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelSchemaExtension", params, optFns, c.addOperationCancelSchemaExtensionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelSchemaExtensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelSchemaExtensionInput struct {

	// The identifier of the directory whose schema extension will be canceled.
	//
	// This member is required.
	DirectoryId *string

	// The identifier of the schema extension that will be canceled.
	//
	// This member is required.
	SchemaExtensionId *string

	noSmithyDocumentSerde
}

type CancelSchemaExtensionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelSchemaExtensionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelSchemaExtension{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelSchemaExtension{}, middleware.After)
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
	if err = addOpCancelSchemaExtensionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelSchemaExtension(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelSchemaExtension(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CancelSchemaExtension",
	}
}