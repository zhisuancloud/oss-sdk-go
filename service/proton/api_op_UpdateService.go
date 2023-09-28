// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Edit a service description or use a spec to add and delete service instances.
// Existing service instances and the service pipeline can't be edited using this
// API. They can only be deleted. Use the description parameter to modify the
// description. Edit the spec parameter to add or delete instances. You can't
// delete a service instance (remove it from the spec) if it has an attached
// component. For more information about components, see Proton components
// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
// Proton Administrator Guide.
func (c *Client) UpdateService(ctx context.Context, params *UpdateServiceInput, optFns ...func(*Options)) (*UpdateServiceOutput, error) {
	if params == nil {
		params = &UpdateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateService", params, optFns, c.addOperationUpdateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceInput struct {

	// The name of the service to edit.
	//
	// This member is required.
	Name *string

	// The edited service description.
	Description *string

	// Lists the service instances to add and the existing service instances to remain.
	// Omit the existing service instances to delete from the list. Don't include edits
	// to the existing service instances or pipeline. For more information, see Edit a
	// service in the Proton Administrator Guide
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-svc-update.html) or the
	// Proton User Guide
	// (https://docs.aws.amazon.com/proton/latest/userguide/ug-svc-update.html).
	//
	// This value conforms to the media type: application/yaml
	Spec *string

	noSmithyDocumentSerde
}

type UpdateServiceOutput struct {

	// The service detail data that's returned by Proton.
	//
	// This member is required.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateService{}, middleware.After)
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
	if err = addOpUpdateServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "UpdateService",
	}
}
