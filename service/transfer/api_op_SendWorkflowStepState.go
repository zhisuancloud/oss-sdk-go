// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a callback for asynchronous custom steps. The ExecutionId, WorkflowId, and
// Token are passed to the target resource during execution of a custom step of a
// workflow. You must include those with their callback as well as providing a
// status.
func (c *Client) SendWorkflowStepState(ctx context.Context, params *SendWorkflowStepStateInput, optFns ...func(*Options)) (*SendWorkflowStepStateOutput, error) {
	if params == nil {
		params = &SendWorkflowStepStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendWorkflowStepState", params, optFns, c.addOperationSendWorkflowStepStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendWorkflowStepStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendWorkflowStepStateInput struct {

	// A unique identifier for the execution of a workflow.
	//
	// This member is required.
	ExecutionId *string

	// Indicates whether the specified step succeeded or failed.
	//
	// This member is required.
	Status types.CustomStepStatus

	// Used to distinguish between multiple callbacks for multiple Lambda steps within
	// the same execution.
	//
	// This member is required.
	Token *string

	// A unique identifier for the workflow.
	//
	// This member is required.
	WorkflowId *string

	noSmithyDocumentSerde
}

type SendWorkflowStepStateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendWorkflowStepStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendWorkflowStepState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendWorkflowStepState{}, middleware.After)
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
	if err = addOpSendWorkflowStepStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendWorkflowStepState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendWorkflowStepState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "SendWorkflowStepState",
	}
}
