// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables for the Network Manager service for an Amazon Web Services Organization.
// This can only be called by a management account within the organization.
func (c *Client) StartOrganizationServiceAccessUpdate(ctx context.Context, params *StartOrganizationServiceAccessUpdateInput, optFns ...func(*Options)) (*StartOrganizationServiceAccessUpdateOutput, error) {
	if params == nil {
		params = &StartOrganizationServiceAccessUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartOrganizationServiceAccessUpdate", params, optFns, c.addOperationStartOrganizationServiceAccessUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartOrganizationServiceAccessUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOrganizationServiceAccessUpdateInput struct {

	// The action to take for the update request. This can be either ENABLE or DISABLE.
	//
	// This member is required.
	Action *string

	noSmithyDocumentSerde
}

type StartOrganizationServiceAccessUpdateOutput struct {

	// The status of the service access update request for an Amazon Web Services
	// Organization.
	OrganizationStatus *types.OrganizationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartOrganizationServiceAccessUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartOrganizationServiceAccessUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartOrganizationServiceAccessUpdate{}, middleware.After)
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
	if err = addOpStartOrganizationServiceAccessUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartOrganizationServiceAccessUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartOrganizationServiceAccessUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "StartOrganizationServiceAccessUpdate",
	}
}
