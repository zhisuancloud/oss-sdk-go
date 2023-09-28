// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about the specified license type conversion task.
func (c *Client) GetLicenseConversionTask(ctx context.Context, params *GetLicenseConversionTaskInput, optFns ...func(*Options)) (*GetLicenseConversionTaskOutput, error) {
	if params == nil {
		params = &GetLicenseConversionTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLicenseConversionTask", params, optFns, c.addOperationGetLicenseConversionTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLicenseConversionTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLicenseConversionTaskInput struct {

	// ID of the license type conversion task to retrieve information on.
	//
	// This member is required.
	LicenseConversionTaskId *string

	noSmithyDocumentSerde
}

type GetLicenseConversionTaskOutput struct {

	// Information about the license type converted to.
	DestinationLicenseContext *types.LicenseConversionContext

	// Time at which the license type conversion task was completed.
	EndTime *time.Time

	// ID of the license type conversion task.
	LicenseConversionTaskId *string

	// Amount of time to complete the license type conversion.
	LicenseConversionTime *time.Time

	// Amazon Resource Names (ARN) of the resources the license conversion task is
	// associated with.
	ResourceArn *string

	// Information about the license type converted from.
	SourceLicenseContext *types.LicenseConversionContext

	// Time at which the license type conversion task was started .
	StartTime *time.Time

	// Status of the license type conversion task.
	Status types.LicenseConversionTaskStatus

	// The status message for the conversion task.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLicenseConversionTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLicenseConversionTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLicenseConversionTask{}, middleware.After)
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
	if err = addOpGetLicenseConversionTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLicenseConversionTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLicenseConversionTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "GetLicenseConversionTask",
	}
}
