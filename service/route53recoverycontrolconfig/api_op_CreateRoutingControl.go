// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/route53recoverycontrolconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new routing control. A routing control has one of two states: ON and
// OFF. You can map the routing control state to the state of an Amazon Route 53
// health check, which can be used to control traffic routing. To get or update the
// routing control state, see the Recovery Cluster (data plane) API actions for
// Amazon Route 53 Application Recovery Controller.
func (c *Client) CreateRoutingControl(ctx context.Context, params *CreateRoutingControlInput, optFns ...func(*Options)) (*CreateRoutingControlOutput, error) {
	if params == nil {
		params = &CreateRoutingControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRoutingControl", params, optFns, c.addOperationCreateRoutingControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRoutingControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The details of the routing control that you're creating.
type CreateRoutingControlInput struct {

	// The Amazon Resource Name (ARN) of the cluster that includes the routing control.
	//
	// This member is required.
	ClusterArn *string

	// The name of the routing control.
	//
	// This member is required.
	RoutingControlName *string

	// A unique, case-sensitive string of up to 64 ASCII characters. To make an
	// idempotent API request with an action, specify a client token in the request.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the control panel that includes the routing
	// control.
	ControlPanelArn *string

	noSmithyDocumentSerde
}

type CreateRoutingControlOutput struct {

	// The routing control that is created.
	RoutingControl *types.RoutingControl

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRoutingControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoutingControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoutingControl{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateRoutingControlMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRoutingControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRoutingControl(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateRoutingControl struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateRoutingControl) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateRoutingControl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateRoutingControlInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateRoutingControlInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateRoutingControlMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateRoutingControl{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateRoutingControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-control-config",
		OperationName: "CreateRoutingControl",
	}
}