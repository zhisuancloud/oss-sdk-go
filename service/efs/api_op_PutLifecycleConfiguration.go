// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this action to manage EFS lifecycle management and intelligent tiering. A
// LifecycleConfiguration consists of one or more LifecyclePolicy objects that
// define the following:
//
// * EFS Lifecycle management - When Amazon EFS
// automatically transitions files in a file system into the lower-cost Infrequent
// Access (IA) storage class. To enable EFS Lifecycle management, set the value of
// TransitionToIA to one of the available options.
//
// * EFS Intelligent tiering -
// When Amazon EFS automatically transitions files from IA back into the file
// system's primary storage class (Standard or One Zone Standard. To enable EFS
// Intelligent Tiering, set the value of TransitionToPrimaryStorageClass to
// AFTER_1_ACCESS.
//
// For more information, see EFS Lifecycle Management
// (https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html).
//
// Each
// Amazon EFS file system supports one lifecycle configuration, which applies to
// all files in the file system. If a LifecycleConfiguration object already exists
// for the specified file system, a PutLifecycleConfiguration call modifies the
// existing configuration. A PutLifecycleConfiguration call with an empty
// LifecyclePolicies array in the request body deletes any existing
// LifecycleConfiguration and turns off lifecycle management and intelligent
// tiering for the file system.
//
// In the request, specify the following:
//
// * The ID
// for the file system for which you are enabling, disabling, or modifying
// lifecycle management and intelligent tiering.
//
// * A LifecyclePolicies array of
// LifecyclePolicy objects that define when files are moved into IA storage, and
// when they are moved back to Standard storage. Amazon EFS requires that each
// LifecyclePolicy object have only have a single transition, so the
// LifecyclePolicies array needs to be structured with separate LifecyclePolicy
// objects. See the example requests in the following section for more
// information.
//
// This operation requires permissions for the
// elasticfilesystem:PutLifecycleConfiguration operation. To apply a
// LifecycleConfiguration object to an encrypted file system, you need the same Key
// Management Service permissions as when you created the encrypted file system.
func (c *Client) PutLifecycleConfiguration(ctx context.Context, params *PutLifecycleConfigurationInput, optFns ...func(*Options)) (*PutLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &PutLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecycleConfiguration", params, optFns, c.addOperationPutLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleConfigurationInput struct {

	// The ID of the file system for which you are creating the LifecycleConfiguration
	// object (String).
	//
	// This member is required.
	FileSystemId *string

	// An array of LifecyclePolicy objects that define the file system's
	// LifecycleConfiguration object. A LifecycleConfiguration object informs EFS
	// lifecycle management and EFS Intelligent-Tiering of the following:
	//
	// * When to
	// move files in the file system from primary storage to the IA storage class.
	//
	// *
	// When to move files that are in IA storage to primary storage.
	//
	// When using the
	// put-lifecycle-configuration CLI command or the PutLifecycleConfiguration API
	// action, Amazon EFS requires that each LifecyclePolicy object have only a single
	// transition. This means that in a request body, LifecyclePolicies must be
	// structured as an array of LifecyclePolicy objects, one object for each
	// transition, TransitionToIA, TransitionToPrimaryStorageClass. See the example
	// requests in the following section for more information.
	//
	// This member is required.
	LifecyclePolicies []types.LifecyclePolicy

	noSmithyDocumentSerde
}

type PutLifecycleConfigurationOutput struct {

	// An array of lifecycle management policies. EFS supports a maximum of one policy
	// per file system.
	LifecyclePolicies []types.LifecyclePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLifecycleConfiguration{}, middleware.After)
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
	if err = addOpPutLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "PutLifecycleConfiguration",
	}
}