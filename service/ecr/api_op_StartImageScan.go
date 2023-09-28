// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an image vulnerability scan. An image scan can only be started once per
// 24 hours on an individual image. This limit includes if an image was scanned on
// initial push. For more information, see Image scanning
// (https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) in
// the Amazon Elastic Container Registry User Guide.
func (c *Client) StartImageScan(ctx context.Context, params *StartImageScanInput, optFns ...func(*Options)) (*StartImageScanOutput, error) {
	if params == nil {
		params = &StartImageScanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImageScan", params, optFns, c.addOperationStartImageScanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImageScanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImageScanInput struct {

	// An object with identifying information for an image in an Amazon ECR repository.
	//
	// This member is required.
	ImageId *types.ImageIdentifier

	// The name of the repository that contains the images to scan.
	//
	// This member is required.
	RepositoryName *string

	// The Amazon Web Services account ID associated with the registry that contains
	// the repository in which to start an image scan request. If you do not specify a
	// registry, the default registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type StartImageScanOutput struct {

	// An object with identifying information for an image in an Amazon ECR repository.
	ImageId *types.ImageIdentifier

	// The current state of the scan.
	ImageScanStatus *types.ImageScanStatus

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImageScanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartImageScan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartImageScan{}, middleware.After)
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
	if err = addOpStartImageScanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImageScan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartImageScan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "StartImageScan",
	}
}