// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a domain. CodeArtifact domains make it easier to manage multiple
// repositories across an organization. You can use a domain to apply permissions
// across many repositories owned by different Amazon Web Services accounts. An
// asset is stored only once in a domain, even if it's in multiple repositories.
// Although you can have multiple domains, we recommend a single production domain
// that contains all published artifacts so that your development teams can find
// and share packages. You can use a second pre-production domain to test changes
// to the production domain configuration.
func (c *Client) CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) {
	if params == nil {
		params = &CreateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomain", params, optFns, c.addOperationCreateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainInput struct {

	// The name of the domain to create. All domain names in an Amazon Web Services
	// Region that are in the same Amazon Web Services account must be unique. The
	// domain name is used as the prefix in DNS hostnames. Do not use sensitive
	// information in a domain name because it is publicly discoverable.
	//
	// This member is required.
	Domain *string

	// The encryption key for the domain. This is used to encrypt content stored in a
	// domain. An encryption key can be a key ID, a key Amazon Resource Name (ARN), a
	// key alias, or a key alias ARN. To specify an encryptionKey, your IAM role must
	// have kms:DescribeKey and kms:CreateGrant permissions on the encryption key that
	// is used. For more information, see DescribeKey
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestSyntax)
	// in the Key Management Service API Reference and Key Management Service API
	// Permissions Reference
	// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
	// in the Key Management Service Developer Guide. CodeArtifact supports only
	// symmetric CMKs. Do not associate an asymmetric CMK with your domain. For more
	// information, see Using symmetric and asymmetric keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
	// in the Key Management Service Developer Guide.
	EncryptionKey *string

	// One or more tag key-value pairs for the domain.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDomainOutput struct {

	// Contains information about the created domain after processing the request.
	Domain *types.DomainDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomain{}, middleware.After)
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
	if err = addOpCreateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "CreateDomain",
	}
}
