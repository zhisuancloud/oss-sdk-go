// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a container image that is registered to your Amazon Lightsail container
// service.
func (c *Client) DeleteContainerImage(ctx context.Context, params *DeleteContainerImageInput, optFns ...func(*Options)) (*DeleteContainerImageOutput, error) {
	if params == nil {
		params = &DeleteContainerImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteContainerImage", params, optFns, c.addOperationDeleteContainerImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteContainerImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteContainerImageInput struct {

	// The name of the container image to delete from the container service. Use the
	// GetContainerImages action to get the name of the container images that are
	// registered to a container service. Container images sourced from your Lightsail
	// container service, that are registered and stored on your service, start with a
	// colon (:). For example, :container-service-1.mystaticwebsite.1. Container images
	// sourced from a public registry like Docker Hub don't start with a colon. For
	// example, nginx:latest or nginx.
	//
	// This member is required.
	Image *string

	// The name of the container service for which to delete a registered container
	// image.
	//
	// This member is required.
	ServiceName *string

	noSmithyDocumentSerde
}

type DeleteContainerImageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteContainerImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteContainerImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteContainerImage{}, middleware.After)
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
	if err = addOpDeleteContainerImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteContainerImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteContainerImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DeleteContainerImage",
	}
}
