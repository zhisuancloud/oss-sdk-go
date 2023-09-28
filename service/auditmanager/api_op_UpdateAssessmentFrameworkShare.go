// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a share request for a custom framework in Audit Manager.
func (c *Client) UpdateAssessmentFrameworkShare(ctx context.Context, params *UpdateAssessmentFrameworkShareInput, optFns ...func(*Options)) (*UpdateAssessmentFrameworkShareOutput, error) {
	if params == nil {
		params = &UpdateAssessmentFrameworkShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssessmentFrameworkShare", params, optFns, c.addOperationUpdateAssessmentFrameworkShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssessmentFrameworkShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssessmentFrameworkShareInput struct {

	// Specifies the update action for the share request.
	//
	// This member is required.
	Action types.ShareRequestAction

	// The unique identifier for the share request.
	//
	// This member is required.
	RequestId *string

	// Specifies whether the share request is a sent request or a received request.
	//
	// This member is required.
	RequestType types.ShareRequestType

	noSmithyDocumentSerde
}

type UpdateAssessmentFrameworkShareOutput struct {

	// The updated share request that's returned by the UpdateAssessmentFrameworkShare
	// operation.
	AssessmentFrameworkShareRequest *types.AssessmentFrameworkShareRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAssessmentFrameworkShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAssessmentFrameworkShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAssessmentFrameworkShare{}, middleware.After)
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
	if err = addOpUpdateAssessmentFrameworkShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssessmentFrameworkShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAssessmentFrameworkShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "UpdateAssessmentFrameworkShare",
	}
}
