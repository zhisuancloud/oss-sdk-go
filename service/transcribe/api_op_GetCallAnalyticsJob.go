// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the specified Call Analytics job. To view the job's
// status, refer to CallAnalyticsJobStatus. If the status is COMPLETED, the job is
// finished. You can find your completed transcript at the URI specified in
// TranscriptFileUri. If the status is FAILED, FailureReason provides details on
// why your transcription job failed. If you enabled personally identifiable
// information (PII) redaction, the redacted transcript appears at the location
// specified in RedactedTranscriptFileUri. If you chose to redact the audio in your
// media file, you can find your redacted media file at the location specified in
// RedactedMediaFileUri. To get a list of your Call Analytics jobs, use the
// operation.
func (c *Client) GetCallAnalyticsJob(ctx context.Context, params *GetCallAnalyticsJobInput, optFns ...func(*Options)) (*GetCallAnalyticsJobOutput, error) {
	if params == nil {
		params = &GetCallAnalyticsJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCallAnalyticsJob", params, optFns, c.addOperationGetCallAnalyticsJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCallAnalyticsJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCallAnalyticsJobInput struct {

	// The name of the Call Analytics job you want information about. Job names are
	// case sensitive.
	//
	// This member is required.
	CallAnalyticsJobName *string

	noSmithyDocumentSerde
}

type GetCallAnalyticsJobOutput struct {

	// Provides detailed information about the specified Call Analytics job, including
	// job status and, if applicable, failure reason.
	CallAnalyticsJob *types.CallAnalyticsJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCallAnalyticsJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCallAnalyticsJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCallAnalyticsJob{}, middleware.After)
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
	if err = addOpGetCallAnalyticsJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCallAnalyticsJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCallAnalyticsJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetCallAnalyticsJob",
	}
}