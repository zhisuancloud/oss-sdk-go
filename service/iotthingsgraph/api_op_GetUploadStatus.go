// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the status of the specified upload.
//
// Deprecated: since: 2022-08-30
func (c *Client) GetUploadStatus(ctx context.Context, params *GetUploadStatusInput, optFns ...func(*Options)) (*GetUploadStatusOutput, error) {
	if params == nil {
		params = &GetUploadStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUploadStatus", params, optFns, c.addOperationGetUploadStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUploadStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUploadStatusInput struct {

	// The ID of the upload. This value is returned by the UploadEntityDefinitions
	// action.
	//
	// This member is required.
	UploadId *string

	noSmithyDocumentSerde
}

type GetUploadStatusOutput struct {

	// The date at which the upload was created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The ID of the upload.
	//
	// This member is required.
	UploadId *string

	// The status of the upload. The initial status is IN_PROGRESS. The response show
	// all validation failures if the upload fails.
	//
	// This member is required.
	UploadStatus types.UploadStatus

	// The reason for an upload failure.
	FailureReason []string

	// The ARN of the upload.
	NamespaceArn *string

	// The name of the upload's namespace.
	NamespaceName *string

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUploadStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUploadStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUploadStatus{}, middleware.After)
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
	if err = addOpGetUploadStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUploadStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUploadStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetUploadStatus",
	}
}
