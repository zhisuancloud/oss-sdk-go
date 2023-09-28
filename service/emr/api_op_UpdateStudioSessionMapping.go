// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the session policy attached to the user or group for the specified
// Amazon EMR Studio.
func (c *Client) UpdateStudioSessionMapping(ctx context.Context, params *UpdateStudioSessionMappingInput, optFns ...func(*Options)) (*UpdateStudioSessionMappingOutput, error) {
	if params == nil {
		params = &UpdateStudioSessionMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStudioSessionMapping", params, optFns, c.addOperationUpdateStudioSessionMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStudioSessionMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStudioSessionMappingInput struct {

	// Specifies whether the identity to update is a user or a group.
	//
	// This member is required.
	IdentityType types.IdentityType

	// The Amazon Resource Name (ARN) of the session policy to associate with the
	// specified user or group.
	//
	// This member is required.
	SessionPolicyArn *string

	// The ID of the Amazon EMR Studio.
	//
	// This member is required.
	StudioId *string

	// The globally unique identifier (GUID) of the user or group. For more
	// information, see UserId
	// (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserId)
	// and GroupId
	// (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-GroupId)
	// in the Amazon Web Services SSO Identity Store API Reference. Either IdentityName
	// or IdentityId must be specified.
	IdentityId *string

	// The name of the user or group to update. For more information, see UserName
	// (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserName)
	// and DisplayName
	// (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-DisplayName)
	// in the Amazon Web Services SSO Identity Store API Reference. Either IdentityName
	// or IdentityId must be specified.
	IdentityName *string

	noSmithyDocumentSerde
}

type UpdateStudioSessionMappingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStudioSessionMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateStudioSessionMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateStudioSessionMapping{}, middleware.After)
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
	if err = addOpUpdateStudioSessionMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStudioSessionMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStudioSessionMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "UpdateStudioSessionMapping",
	}
}
