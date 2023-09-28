// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

func (c *Client) EndpointWithHostLabelHeaderOperation(ctx context.Context, params *EndpointWithHostLabelHeaderOperationInput, optFns ...func(*Options)) (*EndpointWithHostLabelHeaderOperationOutput, error) {
	if params == nil {
		params = &EndpointWithHostLabelHeaderOperationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EndpointWithHostLabelHeaderOperation", params, optFns, c.addOperationEndpointWithHostLabelHeaderOperationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EndpointWithHostLabelHeaderOperationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EndpointWithHostLabelHeaderOperationInput struct {

	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

type EndpointWithHostLabelHeaderOperationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEndpointWithHostLabelHeaderOperationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpEndpointWithHostLabelHeaderOperation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpEndpointWithHostLabelHeaderOperation{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addEndpointPrefix_opEndpointWithHostLabelHeaderOperationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpEndpointWithHostLabelHeaderOperationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEndpointWithHostLabelHeaderOperation(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opEndpointWithHostLabelHeaderOperationMiddleware struct {
}

func (*endpointPrefix_opEndpointWithHostLabelHeaderOperationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opEndpointWithHostLabelHeaderOperationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*EndpointWithHostLabelHeaderOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opEndpointWithHostLabelHeaderOperationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opEndpointWithHostLabelHeaderOperationMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opEndpointWithHostLabelHeaderOperation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EndpointWithHostLabelHeaderOperation",
	}
}
