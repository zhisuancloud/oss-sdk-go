// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an App Runner observability configuration resource. App Runner requires
// this resource when you create or update App Runner services and you want to
// enable non-default observability features. You can share an observability
// configuration across multiple services. Create multiple revisions of a
// configuration by calling this action multiple times using the same
// ObservabilityConfigurationName. The call returns incremental
// ObservabilityConfigurationRevision values. When you create a service and
// configure an observability configuration resource, the service uses the latest
// active revision of the observability configuration by default. You can
// optionally configure the service to use a specific revision. The observability
// configuration resource is designed to configure multiple features (currently one
// feature, tracing). This action takes optional parameters that describe the
// configuration of these features (currently one parameter, TraceConfiguration).
// If you don't specify a feature parameter, App Runner doesn't enable the feature.
func (c *Client) CreateObservabilityConfiguration(ctx context.Context, params *CreateObservabilityConfigurationInput, optFns ...func(*Options)) (*CreateObservabilityConfigurationOutput, error) {
	if params == nil {
		params = &CreateObservabilityConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateObservabilityConfiguration", params, optFns, c.addOperationCreateObservabilityConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateObservabilityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateObservabilityConfigurationInput struct {

	// A name for the observability configuration. When you use it for the first time
	// in an Amazon Web Services Region, App Runner creates revision number 1 of this
	// name. When you use the same name in subsequent calls, App Runner creates
	// incremental revisions of the configuration. The name DefaultConfiguration is
	// reserved. You can't use it to create a new observability configuration, and you
	// can't create a revision of it. When you want to use your own observability
	// configuration for your App Runner service, create a configuration with a
	// different name, and then provide it when you create or update your service.
	//
	// This member is required.
	ObservabilityConfigurationName *string

	// A list of metadata items that you can associate with your observability
	// configuration resource. A tag is a key-value pair.
	Tags []types.Tag

	// The configuration of the tracing feature within this observability
	// configuration. If you don't specify it, App Runner doesn't enable tracing.
	TraceConfiguration *types.TraceConfiguration

	noSmithyDocumentSerde
}

type CreateObservabilityConfigurationOutput struct {

	// A description of the App Runner observability configuration that's created by
	// this request.
	//
	// This member is required.
	ObservabilityConfiguration *types.ObservabilityConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateObservabilityConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateObservabilityConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateObservabilityConfiguration{}, middleware.After)
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
	if err = addOpCreateObservabilityConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateObservabilityConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateObservabilityConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "CreateObservabilityConfiguration",
	}
}