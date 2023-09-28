// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Kinesis Data Analytics application. For information about creating a
// Kinesis Data Analytics application, see Creating an Application
// (https://docs.aws.amazon.com/kinesisanalytics/latest/java/getting-started.html).
func (c *Client) CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) {
	if params == nil {
		params = &CreateApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApplication", params, optFns, c.addOperationCreateApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationInput struct {

	// The name of your application (for example, sample-app).
	//
	// This member is required.
	ApplicationName *string

	// The runtime environment for the application (SQL-1_0, FLINK-1_6, FLINK-1_8, or
	// FLINK-1_11).
	//
	// This member is required.
	RuntimeEnvironment types.RuntimeEnvironment

	// The IAM role used by the application to access Kinesis data streams, Kinesis
	// Data Firehose delivery streams, Amazon S3 objects, and other external resources.
	//
	// This member is required.
	ServiceExecutionRole *string

	// Use this parameter to configure the application.
	ApplicationConfiguration *types.ApplicationConfiguration

	// A summary description of the application.
	ApplicationDescription *string

	// Use the STREAMING mode to create a Kinesis Data Analytics Studio notebook. To
	// create a Kinesis Data Analytics Studio notebook, use the INTERACTIVE mode.
	ApplicationMode types.ApplicationMode

	// Use this parameter to configure an Amazon CloudWatch log stream to monitor
	// application configuration errors.
	CloudWatchLoggingOptions []types.CloudWatchLoggingOption

	// A list of one or more tags to assign to the application. A tag is a key-value
	// pair that identifies an application. Note that the maximum number of application
	// tags includes system tags. The maximum number of user-defined application tags
	// is 50. For more information, see Using Tagging
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/java/how-tagging.html).
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateApplicationOutput struct {

	// In response to your CreateApplication request, Kinesis Data Analytics returns a
	// response with details of the application it created.
	//
	// This member is required.
	ApplicationDetail *types.ApplicationDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateApplication{}, middleware.After)
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
	if err = addOpCreateApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "CreateApplication",
	}
}