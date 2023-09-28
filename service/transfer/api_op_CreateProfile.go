// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the profile for the AS2 process. The agreement is between the partner
// and the AS2 process.
func (c *Client) CreateProfile(ctx context.Context, params *CreateProfileInput, optFns ...func(*Options)) (*CreateProfileOutput, error) {
	if params == nil {
		params = &CreateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProfile", params, optFns, c.addOperationCreateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProfileInput struct {

	// The As2Id is the AS2-name, as defined in the RFC 4130
	// (https://datatracker.ietf.org/doc/html/rfc4130). For inbound transfers, this is
	// the AS2-From header for the AS2 messages sent from the partner. For outbound
	// connectors, this is the AS2-To header for the AS2 messages sent to the partner
	// using the StartFileTransfer API operation. This ID cannot include spaces.
	//
	// This member is required.
	As2Id *string

	// Indicates whether to list only LOCAL type profiles or only PARTNER type
	// profiles. If not supplied in the request, the command lists all types of
	// profiles.
	//
	// This member is required.
	ProfileType types.ProfileType

	// An array of identifiers for the imported certificates. You use this identifier
	// for working with profiles and partner profiles.
	CertificateIds []string

	// Key-value pairs that can be used to group and search for AS2 profiles.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProfileOutput struct {

	// The unique identifier for the AS2 profile, returned after the API call succeeds.
	//
	// This member is required.
	ProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProfile{}, middleware.After)
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
	if err = addOpCreateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateProfile",
	}
}
