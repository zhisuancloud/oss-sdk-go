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

// Updates a launch of a given feature. Don't use this operation to update the tags
// of an existing launch. Instead, use TagResource
// (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_TagResource.html).
func (c *Client) UpdateLaunch(ctx context.Context, params *UpdateLaunchInput, optFns ...func(*Options)) (*UpdateLaunchOutput, error) {
	if params == nil {
		params = &UpdateLaunchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLaunch", params, optFns, c.addOperationUpdateLaunchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLaunchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLaunchInput struct {

	// The name of the launch that is to be updated.
	//
	// This member is required.
	Launch *string

	// The name or ARN of the project that contains the launch that you want to update.
	//
	// This member is required.
	Project *string

	// An optional description for the launch.
	Description *string

	// An array of structures that contains the feature and variations that are to be
	// used for the launch.
	Groups []types.LaunchGroupConfig

	// An array of structures that define the metrics that will be used to monitor the
	// launch performance.
	MetricMonitors []types.MetricMonitorConfig

	// When Evidently assigns a particular user session to a launch, it must use a
	// randomization ID to determine which variation the user session is served. This
	// randomization ID is a combination of the entity ID and randomizationSalt. If you
	// omit randomizationSalt, Evidently uses the launch name as the randomizationSalt.
	RandomizationSalt *string

	// An array of structures that define the traffic allocation percentages among the
	// feature variations during each step of the launch.
	ScheduledSplitsConfig *types.ScheduledSplitsLaunchConfig

	noSmithyDocumentSerde
}

type UpdateLaunchOutput struct {

	// A structure that contains the new configuration of the launch that was updated.
	//
	// This member is required.
	Launch *types.Launch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLaunchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLaunch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLaunch{}, middleware.After)
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
	if err = addOpUpdateLaunchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLaunch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLaunch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "UpdateLaunch",
	}
}
