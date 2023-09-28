// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a preview of a lifecycle policy for the specified repository. This allows
// you to see the results before associating the lifecycle policy with the
// repository.
func (c *Client) StartLifecyclePolicyPreview(ctx context.Context, params *StartLifecyclePolicyPreviewInput, optFns ...func(*Options)) (*StartLifecyclePolicyPreviewOutput, error) {
	if params == nil {
		params = &StartLifecyclePolicyPreviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartLifecyclePolicyPreview", params, optFns, c.addOperationStartLifecyclePolicyPreviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartLifecyclePolicyPreviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartLifecyclePolicyPreviewInput struct {

	// The name of the repository to be evaluated.
	//
	// This member is required.
	RepositoryName *string

	// The policy to be evaluated against. If you do not specify a policy, the current
	// policy for the repository is used.
	LifecyclePolicyText *string

	// The Amazon Web Services account ID associated with the registry that contains
	// the repository. If you do not specify a registry, the default registry is
	// assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type StartLifecyclePolicyPreviewOutput struct {

	// The JSON repository policy text.
	LifecyclePolicyText *string

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// The status of the lifecycle policy preview request.
	Status types.LifecyclePolicyPreviewStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartLifecyclePolicyPreviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartLifecyclePolicyPreview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartLifecyclePolicyPreview{}, middleware.After)
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
	if err = addOpStartLifecyclePolicyPreviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartLifecyclePolicyPreview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartLifecyclePolicyPreview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "StartLifecyclePolicyPreview",
	}
}
