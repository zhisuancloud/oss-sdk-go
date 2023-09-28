// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new domain association for an Amplify app.
func (c *Client) UpdateDomainAssociation(ctx context.Context, params *UpdateDomainAssociationInput, optFns ...func(*Options)) (*UpdateDomainAssociationOutput, error) {
	if params == nil {
		params = &UpdateDomainAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainAssociation", params, optFns, c.addOperationUpdateDomainAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the update domain association request.
type UpdateDomainAssociationInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// Sets the branch patterns for automatic subdomain creation.
	AutoSubDomainCreationPatterns []string

	// The required AWS Identity and Access Management (IAM) service role for the
	// Amazon Resource Name (ARN) for automatically creating subdomains.
	AutoSubDomainIAMRole *string

	// Enables the automated creation of subdomains for branches.
	EnableAutoSubDomain *bool

	// Describes the settings for the subdomain.
	SubDomainSettings []types.SubDomainSetting

	noSmithyDocumentSerde
}

// The result structure for the update domain association request.
type UpdateDomainAssociationOutput struct {

	// Describes a domain association, which associates a custom domain with an Amplify
	// app.
	//
	// This member is required.
	DomainAssociation *types.DomainAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDomainAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDomainAssociation{}, middleware.After)
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
	if err = addOpUpdateDomainAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomainAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "UpdateDomainAssociation",
	}
}
