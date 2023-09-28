// Code generated by smithy-go-codegen DO NOT EDIT.

package voiceid

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/voiceid/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified speaker enrollment job.
func (c *Client) DescribeSpeakerEnrollmentJob(ctx context.Context, params *DescribeSpeakerEnrollmentJobInput, optFns ...func(*Options)) (*DescribeSpeakerEnrollmentJobOutput, error) {
	if params == nil {
		params = &DescribeSpeakerEnrollmentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpeakerEnrollmentJob", params, optFns, c.addOperationDescribeSpeakerEnrollmentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpeakerEnrollmentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSpeakerEnrollmentJobInput struct {

	// The identifier of the domain containing the speaker enrollment job.
	//
	// This member is required.
	DomainId *string

	// The identifier of the speaker enrollment job you are describing.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeSpeakerEnrollmentJobOutput struct {

	// Contains details about the specified speaker enrollment job.
	Job *types.SpeakerEnrollmentJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSpeakerEnrollmentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeSpeakerEnrollmentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeSpeakerEnrollmentJob{}, middleware.After)
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
	if err = addOpDescribeSpeakerEnrollmentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpeakerEnrollmentJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSpeakerEnrollmentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "voiceid",
		OperationName: "DescribeSpeakerEnrollmentJob",
	}
}
