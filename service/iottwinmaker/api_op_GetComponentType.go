// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a component type.
func (c *Client) GetComponentType(ctx context.Context, params *GetComponentTypeInput, optFns ...func(*Options)) (*GetComponentTypeOutput, error) {
	if params == nil {
		params = &GetComponentTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComponentType", params, optFns, c.addOperationGetComponentTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComponentTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComponentTypeInput struct {

	// The ID of the component type.
	//
	// This member is required.
	ComponentTypeId *string

	// The ID of the workspace that contains the component type.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

type GetComponentTypeOutput struct {

	// The ARN of the component type.
	//
	// This member is required.
	Arn *string

	// The ID of the component type.
	//
	// This member is required.
	ComponentTypeId *string

	// The date and time when the component type was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The date and time when the component was last updated.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// The ID of the workspace that contains the component type.
	//
	// This member is required.
	WorkspaceId *string

	// The description of the component type.
	Description *string

	// The name of the parent component type that this component type extends.
	ExtendsFrom []string

	// An object that maps strings to the functions in the component type. Each string
	// in the mapping must be unique to this object.
	Functions map[string]types.FunctionResponse

	// A Boolean value that specifies whether the component type is abstract.
	IsAbstract *bool

	// A Boolean value that specifies whether the component type has a schema
	// initializer and that the schema initializer has run.
	IsSchemaInitialized *bool

	// A Boolean value that specifies whether an entity can have more than one
	// component of this type.
	IsSingleton *bool

	// An object that maps strings to the property definitions in the component type.
	// Each string in the mapping must be unique to this object.
	PropertyDefinitions map[string]types.PropertyDefinitionResponse

	// The current status of the component type.
	Status *types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComponentTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetComponentType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetComponentType{}, middleware.After)
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
	if err = addEndpointPrefix_opGetComponentTypeMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetComponentTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComponentType(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetComponentTypeMiddleware struct {
}

func (*endpointPrefix_opGetComponentTypeMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetComponentTypeMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetComponentTypeMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetComponentTypeMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetComponentType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "GetComponentType",
	}
}
