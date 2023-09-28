// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts an export of a snapshot to Amazon S3. The provided IAM role must have
// access to the S3 bucket. This command doesn't apply to RDS Custom.
func (c *Client) StartExportTask(ctx context.Context, params *StartExportTaskInput, optFns ...func(*Options)) (*StartExportTaskOutput, error) {
	if params == nil {
		params = &StartExportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartExportTask", params, optFns, c.addOperationStartExportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartExportTaskInput struct {

	// A unique identifier for the snapshot export task. This ID isn't an identifier
	// for the Amazon S3 bucket where the snapshot is to be exported to.
	//
	// This member is required.
	ExportTaskIdentifier *string

	// The name of the IAM role to use for writing to the Amazon S3 bucket when
	// exporting a snapshot.
	//
	// This member is required.
	IamRoleArn *string

	// The ID of the Amazon Web Services KMS key to use to encrypt the snapshot
	// exported to Amazon S3. The Amazon Web Services KMS key identifier is the key
	// ARN, key ID, alias ARN, or alias name for the KMS key. The caller of this
	// operation must be authorized to execute the following operations. These can be
	// set in the Amazon Web Services KMS key policy:
	//
	// * GrantOperation.Encrypt
	//
	// *
	// GrantOperation.Decrypt
	//
	// * GrantOperation.GenerateDataKey
	//
	// *
	// GrantOperation.GenerateDataKeyWithoutPlaintext
	//
	// *
	// GrantOperation.ReEncryptFrom
	//
	// * GrantOperation.ReEncryptTo
	//
	// *
	// GrantOperation.CreateGrant
	//
	// * GrantOperation.DescribeKey
	//
	// *
	// GrantOperation.RetireGrant
	//
	// This member is required.
	KmsKeyId *string

	// The name of the Amazon S3 bucket to export the snapshot to.
	//
	// This member is required.
	S3BucketName *string

	// The Amazon Resource Name (ARN) of the snapshot to export to Amazon S3.
	//
	// This member is required.
	SourceArn *string

	// The data to be exported from the snapshot. If this parameter is not provided,
	// all the snapshot data is exported. Valid values are the following:
	//
	// * database -
	// Export all the data from a specified database.
	//
	// * database.table table-name -
	// Export a table of the snapshot. This format is valid only for RDS for MySQL, RDS
	// for MariaDB, and Aurora MySQL.
	//
	// * database.schema schema-name - Export a
	// database schema of the snapshot. This format is valid only for RDS for
	// PostgreSQL and Aurora PostgreSQL.
	//
	// * database.schema.table table-name - Export a
	// table of the database schema. This format is valid only for RDS for PostgreSQL
	// and Aurora PostgreSQL.
	ExportOnly []string

	// The Amazon S3 bucket prefix to use as the file name and path of the exported
	// snapshot.
	S3Prefix *string

	noSmithyDocumentSerde
}

// Contains the details of a snapshot export to Amazon S3. This data type is used
// as a response element in the DescribeExportTasks action.
type StartExportTaskOutput struct {

	// The data exported from the snapshot. Valid values are the following:
	//
	// * database
	// - Export all the data from a specified database.
	//
	// * database.table table-name -
	// Export a table of the snapshot. This format is valid only for RDS for MySQL, RDS
	// for MariaDB, and Aurora MySQL.
	//
	// * database.schema schema-name - Export a
	// database schema of the snapshot. This format is valid only for RDS for
	// PostgreSQL and Aurora PostgreSQL.
	//
	// * database.schema.table table-name - Export a
	// table of the database schema. This format is valid only for RDS for PostgreSQL
	// and Aurora PostgreSQL.
	ExportOnly []string

	// A unique identifier for the snapshot export task. This ID isn't an identifier
	// for the Amazon S3 bucket where the snapshot is exported to.
	ExportTaskIdentifier *string

	// The reason the export failed, if it failed.
	FailureCause *string

	// The name of the IAM role that is used to write to Amazon S3 when exporting a
	// snapshot.
	IamRoleArn *string

	// The key identifier of the Amazon Web Services KMS key that is used to encrypt
	// the snapshot when it's exported to Amazon S3. The KMS key identifier is its key
	// ARN, key ID, alias ARN, or alias name. The IAM role used for the snapshot export
	// must have encryption and decryption permissions to use this KMS key.
	KmsKeyId *string

	// The progress of the snapshot export task as a percentage.
	PercentProgress int32

	// The Amazon S3 bucket that the snapshot is exported to.
	S3Bucket *string

	// The Amazon S3 bucket prefix that is the file name and path of the exported
	// snapshot.
	S3Prefix *string

	// The time that the snapshot was created.
	SnapshotTime *time.Time

	// The Amazon Resource Name (ARN) of the snapshot exported to Amazon S3.
	SourceArn *string

	// The progress status of the export task.
	Status *string

	// The time that the snapshot export task completed.
	TaskEndTime *time.Time

	// The time that the snapshot export task started.
	TaskStartTime *time.Time

	// The total amount of data exported, in gigabytes.
	TotalExtractedDataInGB int32

	// A warning about the snapshot export task.
	WarningMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartExportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStartExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStartExportTask{}, middleware.After)
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
	if err = addOpStartExportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartExportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartExportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StartExportTask",
	}
}
