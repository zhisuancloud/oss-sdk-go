// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates properties on a ledger.
func (c *Client) UpdateLedger(ctx context.Context, params *UpdateLedgerInput, optFns ...func(*Options)) (*UpdateLedgerOutput, error) {
	if params == nil {
		params = &UpdateLedgerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLedger", params, optFns, c.addOperationUpdateLedgerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLedgerInput struct {

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The flag that prevents a ledger from being deleted by any user. If not provided
	// on ledger creation, this feature is enabled (true) by default. If deletion
	// protection is enabled, you must first disable it before you can delete the
	// ledger. You can disable it by calling the UpdateLedger operation to set the flag
	// to false.
	DeletionProtection *bool

	// The key in Key Management Service (KMS) to use for encryption of data at rest in
	// the ledger. For more information, see Encryption at rest
	// (https://docs.aws.amazon.com/qldb/latest/developerguide/encryption-at-rest.html)
	// in the Amazon QLDB Developer Guide. Use one of the following options to specify
	// this parameter:
	//
	// * AWS_OWNED_KMS_KEY: Use an KMS key that is owned and managed
	// by Amazon Web Services on your behalf.
	//
	// * Undefined: Make no changes to the KMS
	// key of the ledger.
	//
	// * A valid symmetric customer managed KMS key: Use the
	// specified KMS key in your account that you create, own, and manage. Amazon QLDB
	// does not support asymmetric keys. For more information, see Using symmetric and
	// asymmetric keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
	// in the Key Management Service Developer Guide.
	//
	// To specify a customer managed
	// KMS key, you can use its key ID, Amazon Resource Name (ARN), alias name, or
	// alias ARN. When using an alias name, prefix it with "alias/". To specify a key
	// in a different Amazon Web Services account, you must use the key ARN or alias
	// ARN. For example:
	//
	// * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// *
	// Alias name: alias/ExampleAlias
	//
	// * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// For more information, see
	// Key identifiers (KeyId)
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in
	// the Key Management Service Developer Guide.
	KmsKey *string

	noSmithyDocumentSerde
}

type UpdateLedgerOutput struct {

	// The Amazon Resource Name (ARN) for the ledger.
	Arn *string

	// The date and time, in epoch time format, when the ledger was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1, 1970
	// UTC.)
	CreationDateTime *time.Time

	// The flag that prevents a ledger from being deleted by any user. If not provided
	// on ledger creation, this feature is enabled (true) by default. If deletion
	// protection is enabled, you must first disable it before you can delete the
	// ledger. You can disable it by calling the UpdateLedger operation to set the flag
	// to false.
	DeletionProtection *bool

	// Information about the encryption of data at rest in the ledger. This includes
	// the current status, the KMS key, and when the key became inaccessible (in the
	// case of an error).
	EncryptionDescription *types.LedgerEncryptionDescription

	// The name of the ledger.
	Name *string

	// The current status of the ledger.
	State types.LedgerState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLedger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLedger{}, middleware.After)
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
	if err = addOpUpdateLedgerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLedger(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "UpdateLedger",
	}
}