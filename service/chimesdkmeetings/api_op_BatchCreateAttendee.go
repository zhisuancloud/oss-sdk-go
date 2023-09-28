// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/chimesdkmeetings/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates up to 100 attendees for an active Amazon Chime SDK meeting. For more
// information about the Amazon Chime SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
func (c *Client) BatchCreateAttendee(ctx context.Context, params *BatchCreateAttendeeInput, optFns ...func(*Options)) (*BatchCreateAttendeeOutput, error) {
	if params == nil {
		params = &BatchCreateAttendeeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCreateAttendee", params, optFns, c.addOperationBatchCreateAttendeeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCreateAttendeeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateAttendeeInput struct {

	// The attendee information, including attendees' IDs and join tokens.
	//
	// This member is required.
	Attendees []types.CreateAttendeeRequestItem

	// The Amazon Chime SDK ID of the meeting to which you're adding attendees.
	//
	// This member is required.
	MeetingId *string

	noSmithyDocumentSerde
}

type BatchCreateAttendeeOutput struct {

	// The attendee information, including attendees' IDs and join tokens.
	Attendees []types.Attendee

	// If the action fails for one or more of the attendees in the request, a list of
	// the attendees is returned, along with error codes and error messages.
	Errors []types.CreateAttendeeError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchCreateAttendeeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchCreateAttendee{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchCreateAttendee{}, middleware.After)
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
	if err = addOpBatchCreateAttendeeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateAttendee(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchCreateAttendee(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchCreateAttendee",
	}
}
