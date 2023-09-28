// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2instanceconnect

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Pushes an SSH public key to the specified EC2 instance for use by the specified
// user. The key remains for 60 seconds. For more information, see Connect to your
// Linux instance using EC2 Instance Connect
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Connect-using-EC2-Instance-Connect.html)
// in the Amazon EC2 User Guide.
func (c *Client) SendSSHPublicKey(ctx context.Context, params *SendSSHPublicKeyInput, optFns ...func(*Options)) (*SendSSHPublicKeyOutput, error) {
	if params == nil {
		params = &SendSSHPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendSSHPublicKey", params, optFns, c.addOperationSendSSHPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendSSHPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendSSHPublicKeyInput struct {

	// The ID of the EC2 instance.
	//
	// This member is required.
	InstanceId *string

	// The OS user on the EC2 instance for whom the key can be used to authenticate.
	//
	// This member is required.
	InstanceOSUser *string

	// The public key material. To use the public key, you must have the matching
	// private key.
	//
	// This member is required.
	SSHPublicKey *string

	// The Availability Zone in which the EC2 instance was launched.
	AvailabilityZone *string

	noSmithyDocumentSerde
}

type SendSSHPublicKeyOutput struct {

	// The ID of the request. Please provide this ID when contacting AWS Support for
	// assistance.
	RequestId *string

	// Is true if the request succeeds and an error otherwise.
	Success bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendSSHPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendSSHPublicKey{}, middleware.After)
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
	if err = addOpSendSSHPublicKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendSSHPublicKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendSSHPublicKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2-instance-connect",
		OperationName: "SendSSHPublicKey",
	}
}
