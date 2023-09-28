// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/privatenetworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates the specified device identifier.
func (c *Client) ActivateDeviceIdentifier(ctx context.Context, params *ActivateDeviceIdentifierInput, optFns ...func(*Options)) (*ActivateDeviceIdentifierOutput, error) {
	if params == nil {
		params = &ActivateDeviceIdentifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ActivateDeviceIdentifier", params, optFns, c.addOperationActivateDeviceIdentifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ActivateDeviceIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ActivateDeviceIdentifierInput struct {

	// The Amazon Resource Name (ARN) of the device identifier.
	//
	// This member is required.
	DeviceIdentifierArn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string

	noSmithyDocumentSerde
}

type ActivateDeviceIdentifierOutput struct {

	// Information about the device identifier.
	//
	// This member is required.
	DeviceIdentifier *types.DeviceIdentifier

	// The tags on the device identifier.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationActivateDeviceIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpActivateDeviceIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpActivateDeviceIdentifier{}, middleware.After)
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
	if err = addOpActivateDeviceIdentifierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opActivateDeviceIdentifier(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opActivateDeviceIdentifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "private-networks",
		OperationName: "ActivateDeviceIdentifier",
	}
}
