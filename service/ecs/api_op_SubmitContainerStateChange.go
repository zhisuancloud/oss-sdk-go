// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action is only used by the Amazon ECS agent, and it is not intended for use
// outside of the agent. Sent to acknowledge that a container changed states.
func (c *Client) SubmitContainerStateChange(ctx context.Context, params *SubmitContainerStateChangeInput, optFns ...func(*Options)) (*SubmitContainerStateChangeOutput, error) {
	if params == nil {
		params = &SubmitContainerStateChangeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubmitContainerStateChange", params, optFns, c.addOperationSubmitContainerStateChangeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubmitContainerStateChangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitContainerStateChangeInput struct {

	// The short name or full ARN of the cluster that hosts the container.
	Cluster *string

	// The name of the container.
	ContainerName *string

	// The exit code that's returned for the state change request.
	ExitCode *int32

	// The network bindings of the container.
	NetworkBindings []types.NetworkBinding

	// The reason for the state change request.
	Reason *string

	// The ID of the Docker container.
	RuntimeId *string

	// The status of the state change request.
	Status *string

	// The task ID or full Amazon Resource Name (ARN) of the task that hosts the
	// container.
	Task *string

	noSmithyDocumentSerde
}

type SubmitContainerStateChangeOutput struct {

	// Acknowledgement of the state change.
	Acknowledgment *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSubmitContainerStateChangeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSubmitContainerStateChange{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubmitContainerStateChange{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitContainerStateChange(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubmitContainerStateChange(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "SubmitContainerStateChange",
	}
}
