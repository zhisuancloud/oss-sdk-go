// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the current snapshot for the patch baseline the managed node uses.
// This API is primarily used by the AWS-RunPatchBaseline Systems Manager document
// (SSM document). If you run the command locally, such as with the Command Line
// Interface (CLI), the system attempts to use your local Amazon Web Services
// credentials and the operation fails. To avoid this, you can run the command in
// the Amazon Web Services Systems Manager console. Use Run Command, a capability
// of Amazon Web Services Systems Manager, with an SSM document that enables you to
// target a managed node with a script or command. For example, run the command
// using the AWS-RunShellScript document or the AWS-RunPowerShellScript document.
func (c *Client) GetDeployablePatchSnapshotForInstance(ctx context.Context, params *GetDeployablePatchSnapshotForInstanceInput, optFns ...func(*Options)) (*GetDeployablePatchSnapshotForInstanceOutput, error) {
	if params == nil {
		params = &GetDeployablePatchSnapshotForInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeployablePatchSnapshotForInstance", params, optFns, c.addOperationGetDeployablePatchSnapshotForInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeployablePatchSnapshotForInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeployablePatchSnapshotForInstanceInput struct {

	// The ID of the managed node for which the appropriate patch snapshot should be
	// retrieved.
	//
	// This member is required.
	InstanceId *string

	// The snapshot ID provided by the user when running AWS-RunPatchBaseline.
	//
	// This member is required.
	SnapshotId *string

	// Defines the basic information about a patch baseline override.
	BaselineOverride *types.BaselineOverride

	noSmithyDocumentSerde
}

type GetDeployablePatchSnapshotForInstanceOutput struct {

	// The managed node ID.
	InstanceId *string

	// Returns the specific operating system (for example Windows Server 2012 or Amazon
	// Linux 2015.09) on the managed node for the specified patch snapshot.
	Product *string

	// A pre-signed Amazon Simple Storage Service (Amazon S3) URL that can be used to
	// download the patch snapshot.
	SnapshotDownloadUrl *string

	// The user-defined snapshot ID.
	SnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeployablePatchSnapshotForInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeployablePatchSnapshotForInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeployablePatchSnapshotForInstance{}, middleware.After)
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
	if err = addOpGetDeployablePatchSnapshotForInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeployablePatchSnapshotForInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeployablePatchSnapshotForInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetDeployablePatchSnapshotForInstance",
	}
}
