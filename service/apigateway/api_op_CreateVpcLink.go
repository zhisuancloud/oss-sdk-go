// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a VPC link, under the caller's account in a selected region, in an
// asynchronous operation that typically takes 2-4 minutes to complete and become
// operational. The caller must have permissions to create and update VPC Endpoint
// services.
func (c *Client) CreateVpcLink(ctx context.Context, params *CreateVpcLinkInput, optFns ...func(*Options)) (*CreateVpcLinkOutput, error) {
	if params == nil {
		params = &CreateVpcLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcLink", params, optFns, c.addOperationCreateVpcLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a VPC link, under the caller's account in a selected region, in an
// asynchronous operation that typically takes 2-4 minutes to complete and become
// operational. The caller must have permissions to create and update VPC Endpoint
// services.
type CreateVpcLinkInput struct {

	// The name used to label and identify the VPC link.
	//
	// This member is required.
	Name *string

	// The ARN of the network load balancer of the VPC targeted by the VPC link. The
	// network load balancer must be owned by the same AWS account of the API owner.
	//
	// This member is required.
	TargetArns []string

	// The description of the VPC link.
	Description *string

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws:. The tag value
	// can be up to 256 characters.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An API Gateway VPC link for a RestApi to access resources in an Amazon Virtual
// Private Cloud (VPC).
type CreateVpcLinkOutput struct {

	// The description of the VPC link.
	Description *string

	// The identifier of the VpcLink. It is used in an Integration to reference this
	// VpcLink.
	Id *string

	// The name used to label and identify the VPC link.
	Name *string

	// The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING,
	// or FAILED. Deploying an API will wait if the status is PENDING and will fail if
	// the status is DELETING.
	Status types.VpcLinkStatus

	// A description about the VPC link status.
	StatusMessage *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// The ARN of the network load balancer of the VPC targeted by the VPC link. The
	// network load balancer must be owned by the same AWS account of the API owner.
	TargetArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVpcLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVpcLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVpcLink{}, middleware.After)
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
	if err = addOpCreateVpcLinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcLink(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateVpcLink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateVpcLink",
	}
}