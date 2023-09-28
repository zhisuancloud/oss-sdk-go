// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing feature. You can't use this operation to update the tags of
// an existing feature. Instead, use TagResource
// (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_TagResource.html).
func (c *Client) UpdateFeature(ctx context.Context, params *UpdateFeatureInput, optFns ...func(*Options)) (*UpdateFeatureOutput, error) {
	if params == nil {
		params = &UpdateFeatureInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFeature", params, optFns, c.addOperationUpdateFeatureMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFeatureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFeatureInput struct {

	// The name of the feature to be updated.
	//
	// This member is required.
	Feature *string

	// The name or ARN of the project that contains the feature to be updated.
	//
	// This member is required.
	Project *string

	// To update variation configurations for this feature, or add new ones, specify
	// this structure. In this array, include any variations that you want to add or
	// update. If the array includes a variation name that already exists for this
	// feature, it is updated. If it includes a new variation name, it is added as a
	// new variation.
	AddOrUpdateVariations []types.VariationConfig

	// The name of the variation to use as the default variation. The default variation
	// is served to users who are not allocated to any ongoing launches or experiments
	// of this feature.
	DefaultVariation *string

	// An optional description of the feature.
	Description *string

	// Specified users that should always be served a specific variation of a feature.
	// Each user is specified by a key-value pair . For each key, specify a user by
	// entering their user ID, account ID, or some other identifier. For the value,
	// specify the name of the variation that they are to be served.
	EntityOverrides map[string]string

	// Specify ALL_RULES to activate the traffic allocation specified by any ongoing
	// launches or experiments. Specify DEFAULT_VARIATION to serve the default
	// variation to all users instead.
	EvaluationStrategy types.FeatureEvaluationStrategy

	// Removes a variation from the feature. If the variation you specify doesn't
	// exist, then this makes no change and does not report an error. This operation
	// fails if you try to remove a variation that is part of an ongoing launch or
	// experiment.
	RemoveVariations []string

	noSmithyDocumentSerde
}

type UpdateFeatureOutput struct {

	// A structure that contains information about the updated feature.
	//
	// This member is required.
	Feature *types.Feature

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFeatureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFeature{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFeature{}, middleware.After)
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
	if err = addOpUpdateFeatureValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFeature(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFeature(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "UpdateFeature",
	}
}