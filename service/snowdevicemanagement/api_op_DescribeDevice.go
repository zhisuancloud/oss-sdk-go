// Code generated by smithy-go-codegen DO NOT EDIT.

package snowdevicemanagement

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/snowdevicemanagement/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Checks device-specific information, such as the device type, software version,
// IP addresses, and lock status.
func (c *Client) DescribeDevice(ctx context.Context, params *DescribeDeviceInput, optFns ...func(*Options)) (*DescribeDeviceOutput, error) {
	if params == nil {
		params = &DescribeDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDevice", params, optFns, c.addOperationDescribeDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDeviceInput struct {

	// The ID of the device that you are checking the information of.
	//
	// This member is required.
	ManagedDeviceId *string

	noSmithyDocumentSerde
}

type DescribeDeviceOutput struct {

	// The ID of the job used when ordering the device.
	AssociatedWithJob *string

	// The hardware specifications of the device.
	DeviceCapacities []types.Capacity

	// The current state of the device.
	DeviceState types.UnlockState

	// The type of Amazon Web Services Snow Family device.
	DeviceType *string

	// When the device last contacted the Amazon Web Services Cloud. Indicates that the
	// device is online.
	LastReachedOutAt *time.Time

	// When the device last pushed an update to the Amazon Web Services Cloud.
	// Indicates when the device cache was refreshed.
	LastUpdatedAt *time.Time

	// The Amazon Resource Name (ARN) of the device.
	ManagedDeviceArn *string

	// The ID of the device that you checked the information for.
	ManagedDeviceId *string

	// The network interfaces available on the device.
	PhysicalNetworkInterfaces []types.PhysicalNetworkInterface

	// The software installed on the device.
	Software *types.SoftwareInformation

	// Optional metadata that you assign to a resource. You can use tags to categorize
	// a resource in different ways, such as by purpose, owner, or environment.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDevice{}, middleware.After)
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
	if err = addOpDescribeDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snow-device-management",
		OperationName: "DescribeDevice",
	}
}
