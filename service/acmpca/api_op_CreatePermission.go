// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants one or more permissions on a private CA to the Certificate Manager (ACM)
// service principal (acm.amazonaws.com). These permissions allow ACM to issue and
// renew ACM certificates that reside in the same Amazon Web Services account as
// the CA. You can list current permissions with the ListPermissions
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_ListPermissions.html)
// action and revoke them with the DeletePermission
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_DeletePermission.html)
// action. About Permissions
//
// * If the private CA and the certificates it issues
// reside in the same account, you can use CreatePermission to grant permissions
// for ACM to carry out automatic certificate renewals.
//
// * For automatic
// certificate renewal to succeed, the ACM service principal needs permissions to
// create, retrieve, and list certificates.
//
// * If the private CA and the ACM
// certificates reside in different accounts, then permissions cannot be used to
// enable automatic renewals. Instead, the ACM certificate owner must set up a
// resource-based policy to enable cross-account issuance and renewals. For more
// information, see Using a Resource Based Policy with ACM Private CA
// (https://docs.aws.amazon.com/acm-pca/latest/userguide/pca-rbp.html).
func (c *Client) CreatePermission(ctx context.Context, params *CreatePermissionInput, optFns ...func(*Options)) (*CreatePermissionOutput, error) {
	if params == nil {
		params = &CreatePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePermission", params, optFns, c.addOperationCreatePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePermissionInput struct {

	// The actions that the specified Amazon Web Services service principal can use.
	// These include IssueCertificate, GetCertificate, and ListPermissions.
	//
	// This member is required.
	Actions []types.ActionType

	// The Amazon Resource Name (ARN) of the CA that grants the permissions. You can
	// find the ARN by calling the ListCertificateAuthorities
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_ListCertificateAuthorities.html)
	// action. This must have the following form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string

	// The Amazon Web Services service or identity that receives the permission. At
	// this time, the only valid principal is acm.amazonaws.com.
	//
	// This member is required.
	Principal *string

	// The ID of the calling account.
	SourceAccount *string

	noSmithyDocumentSerde
}

type CreatePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePermission{}, middleware.After)
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
	if err = addOpCreatePermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "CreatePermission",
	}
}
