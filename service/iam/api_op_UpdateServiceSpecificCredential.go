// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the status of a service-specific credential to Active or Inactive.
// Service-specific credentials that are inactive cannot be used for authentication
// to the service. This operation can be used to disable a user's service-specific
// credential as part of a credential rotation work flow.
func (c *Client) UpdateServiceSpecificCredential(ctx context.Context, params *UpdateServiceSpecificCredentialInput, optFns ...func(*Options)) (*UpdateServiceSpecificCredentialOutput, error) {
	if params == nil {
		params = &UpdateServiceSpecificCredentialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServiceSpecificCredential", params, optFns, c.addOperationUpdateServiceSpecificCredentialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceSpecificCredentialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceSpecificCredentialInput struct {

	// The unique identifier of the service-specific credential. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters that can consist of any upper or lowercased letter or digit.
	//
	// This member is required.
	ServiceSpecificCredentialId *string

	// The status to be assigned to the service-specific credential.
	//
	// This member is required.
	Status types.StatusType

	// The name of the IAM user associated with the service-specific credential. If you
	// do not specify this value, then the operation assumes the user whose credentials
	// are used to call the operation. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	UserName *string

	noSmithyDocumentSerde
}

type UpdateServiceSpecificCredentialOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceSpecificCredentialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateServiceSpecificCredential{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateServiceSpecificCredential{}, middleware.After)
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
	if err = addOpUpdateServiceSpecificCredentialValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceSpecificCredential(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServiceSpecificCredential(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateServiceSpecificCredential",
	}
}
