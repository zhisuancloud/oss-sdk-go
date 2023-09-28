// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables EBS encryption by default for your account in the current Region. After
// you enable encryption by default, the EBS volumes that you create are always
// encrypted, either using the default KMS key or the KMS key that you specified
// when you created each volume. For more information, see Amazon EBS encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide. You can specify the default KMS key for
// encryption by default using ModifyEbsDefaultKmsKeyId or ResetEbsDefaultKmsKeyId.
// Enabling encryption by default has no effect on the encryption status of your
// existing volumes. After you enable encryption by default, you can no longer
// launch instances using instance types that do not support encryption. For more
// information, see Supported instance types
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
func (c *Client) EnableEbsEncryptionByDefault(ctx context.Context, params *EnableEbsEncryptionByDefaultInput, optFns ...func(*Options)) (*EnableEbsEncryptionByDefaultOutput, error) {
	if params == nil {
		params = &EnableEbsEncryptionByDefaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableEbsEncryptionByDefault", params, optFns, c.addOperationEnableEbsEncryptionByDefaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableEbsEncryptionByDefaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableEbsEncryptionByDefaultInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type EnableEbsEncryptionByDefaultOutput struct {

	// The updated status of encryption by default.
	EbsEncryptionByDefault *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableEbsEncryptionByDefaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpEnableEbsEncryptionByDefault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpEnableEbsEncryptionByDefault{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableEbsEncryptionByDefault(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableEbsEncryptionByDefault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableEbsEncryptionByDefault",
	}
}
