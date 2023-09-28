// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Recovers the saved resource identified by an Amazon Resource Name (ARN).
func (c *Client) StartRestoreJob(ctx context.Context, params *StartRestoreJobInput, optFns ...func(*Options)) (*StartRestoreJobOutput, error) {
	if params == nil {
		params = &StartRestoreJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartRestoreJob", params, optFns, c.addOperationStartRestoreJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartRestoreJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRestoreJobInput struct {

	// A set of metadata key-value pairs. Contains information, such as a resource
	// name, required to restore a recovery point. You can get configuration metadata
	// about a resource at the time it was backed up by calling
	// GetRecoveryPointRestoreMetadata. However, values in addition to those provided
	// by GetRecoveryPointRestoreMetadata might be required to restore a resource. For
	// example, you might need to provide a new resource name if the original already
	// exists. You need to specify specific metadata to restore an Amazon Elastic File
	// System (Amazon EFS) instance:
	//
	// * file-system-id: The ID of the Amazon EFS file
	// system that is backed up by Backup. Returned in
	// GetRecoveryPointRestoreMetadata.
	//
	// * Encrypted: A Boolean value that, if true,
	// specifies that the file system is encrypted. If KmsKeyId is specified, Encrypted
	// must be set to true.
	//
	// * KmsKeyId: Specifies the Amazon Web Services KMS key that
	// is used to encrypt the restored file system. You can specify a key from another
	// Amazon Web Services account provided that key it is properly shared with your
	// account via Amazon Web Services KMS.
	//
	// * PerformanceMode: Specifies the
	// throughput mode of the file system.
	//
	// * CreationToken: A user-supplied value that
	// ensures the uniqueness (idempotency) of the request.
	//
	// * newFileSystem: A Boolean
	// value that, if true, specifies that the recovery point is restored to a new
	// Amazon EFS file system.
	//
	// * ItemsToRestore: An array of one to five strings where
	// each string is a file path. Use ItemsToRestore to restore specific files or
	// directories rather than the entire file system. This parameter is optional. For
	// example, "itemsToRestore":"[\"/my.test\"]".
	//
	// This member is required.
	Metadata map[string]string

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	//
	// This member is required.
	RecoveryPointArn *string

	// The Amazon Resource Name (ARN) of the IAM role that Backup uses to create the
	// target recovery point; for example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string

	// A customer-chosen string that you can use to distinguish between otherwise
	// identical calls to StartRestoreJob. Retrying a successful request with the same
	// idempotency token results in a success message with no action taken.
	IdempotencyToken *string

	// Starts a job to restore a recovery point for one of the following resources:
	//
	// *
	// Aurora for Amazon Aurora
	//
	// * DocumentDB for Amazon DocumentDB (with MongoDB
	// compatibility)
	//
	// * DynamoDB for Amazon DynamoDB
	//
	// * EBS for Amazon Elastic Block
	// Store
	//
	// * EC2 for Amazon Elastic Compute Cloud
	//
	// * EFS for Amazon Elastic File
	// System
	//
	// * FSx for Amazon FSx
	//
	// * Neptune for Amazon Neptune
	//
	// * RDS for Amazon
	// Relational Database Service
	//
	// * Storage Gateway for Storage Gateway
	//
	// * S3 for
	// Amazon S3
	//
	// * VirtualMachine for virtual machines
	ResourceType *string

	noSmithyDocumentSerde
}

type StartRestoreJobOutput struct {

	// Uniquely identifies the job that restores a recovery point.
	RestoreJobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartRestoreJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartRestoreJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartRestoreJob{}, middleware.After)
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
	if err = addOpStartRestoreJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartRestoreJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartRestoreJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "StartRestoreJob",
	}
}
