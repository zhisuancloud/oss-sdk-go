// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about the specified custom medical vocabulary. To view the
// status of the specified medical vocabulary, check the VocabularyState field. If
// the status is READY, your vocabulary is available to use. If the status is
// FAILED, FailureReason provides details on why your vocabulary failed. To get a
// list of your custom medical vocabularies, use the operation.
func (c *Client) GetMedicalVocabulary(ctx context.Context, params *GetMedicalVocabularyInput, optFns ...func(*Options)) (*GetMedicalVocabularyOutput, error) {
	if params == nil {
		params = &GetMedicalVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMedicalVocabulary", params, optFns, c.addOperationGetMedicalVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMedicalVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMedicalVocabularyInput struct {

	// The name of the custom medical vocabulary you want information about. Vocabulary
	// names are case sensitive.
	//
	// This member is required.
	VocabularyName *string

	noSmithyDocumentSerde
}

type GetMedicalVocabularyOutput struct {

	// The S3 location where the specified medical vocabulary is stored; use this URI
	// to view or download the vocabulary.
	DownloadUri *string

	// If VocabularyState is FAILED, FailureReason contains information about why the
	// medical vocabulary request failed. See also: Common Errors
	// (https://docs.aws.amazon.com/transcribe/latest/APIReference/CommonErrors.html).
	FailureReason *string

	// The language code you selected for your medical vocabulary. US English (en-US)
	// is the only language supported with Amazon Transcribe Medical.
	LanguageCode types.LanguageCode

	// The date and time the specified custom medical vocabulary was last modified.
	// Timestamps are in the format YYYY-MM-DD'T'HH:MM:SS.SSSSSS-UTC. For example,
	// 2022-05-04T12:32:58.761000-07:00 represents 12:32 PM UTC-7 on May 4, 2022.
	LastModifiedTime *time.Time

	// The name of the custom medical vocabulary you requested information about.
	VocabularyName *string

	// The processing state of your custom medical vocabulary. If the state is READY,
	// you can use the vocabulary in a StartMedicalTranscriptionJob request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMedicalVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMedicalVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMedicalVocabulary{}, middleware.After)
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
	if err = addOpGetMedicalVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMedicalVocabulary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMedicalVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetMedicalVocabulary",
	}
}