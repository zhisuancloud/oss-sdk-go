// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a versioned model.
func (c *Client) UpdateModelPackage(ctx context.Context, params *UpdateModelPackageInput, optFns ...func(*Options)) (*UpdateModelPackageOutput, error) {
	if params == nil {
		params = &UpdateModelPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateModelPackage", params, optFns, c.addOperationUpdateModelPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateModelPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateModelPackageInput struct {

	// The Amazon Resource Name (ARN) of the model package.
	//
	// This member is required.
	ModelPackageArn *string

	// An array of additional Inference Specification objects to be added to the
	// existing array additional Inference Specification. Total number of additional
	// Inference Specifications can not exceed 15. Each additional Inference
	// Specification specifies artifacts based on this model package that can be used
	// on inference endpoints. Generally used with SageMaker Neo to store the compiled
	// artifacts.
	AdditionalInferenceSpecificationsToAdd []types.AdditionalInferenceSpecificationDefinition

	// A description for the approval status of the model.
	ApprovalDescription *string

	// The metadata properties associated with the model package versions.
	CustomerMetadataProperties map[string]string

	// The metadata properties associated with the model package versions to remove.
	CustomerMetadataPropertiesToRemove []string

	// The approval status of the model.
	ModelApprovalStatus types.ModelApprovalStatus

	noSmithyDocumentSerde
}

type UpdateModelPackageOutput struct {

	// The Amazon Resource Name (ARN) of the model.
	//
	// This member is required.
	ModelPackageArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateModelPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateModelPackage{}, middleware.After)
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
	if err = addOpUpdateModelPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateModelPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateModelPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateModelPackage",
	}
}