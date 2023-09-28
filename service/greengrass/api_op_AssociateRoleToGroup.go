// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a role with a group. Your Greengrass core will use the role to access
// AWS cloud services. The role's permissions should allow Greengrass core Lambda
// functions to perform actions against the cloud.
func (c *Client) AssociateRoleToGroup(ctx context.Context, params *AssociateRoleToGroupInput, optFns ...func(*Options)) (*AssociateRoleToGroupOutput, error) {
	if params == nil {
		params = &AssociateRoleToGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateRoleToGroup", params, optFns, c.addOperationAssociateRoleToGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateRoleToGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateRoleToGroupInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string

	// The ARN of the role you wish to associate with this group. The existence of the
	// role is not validated.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

type AssociateRoleToGroupOutput struct {

	// The time, in milliseconds since the epoch, when the role ARN was associated with
	// the group.
	AssociatedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateRoleToGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateRoleToGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateRoleToGroup{}, middleware.After)
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
	if err = addOpAssociateRoleToGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateRoleToGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateRoleToGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "AssociateRoleToGroup",
	}
}
