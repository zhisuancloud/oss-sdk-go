// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports a specified server engine attribute as a base64-encoded string. For
// example, you can export user data that you can use in EC2 to associate nodes
// with a server. This operation is synchronous. A ValidationException is raised
// when parameters of the request are not valid. A ResourceNotFoundException is
// thrown when the server does not exist. An InvalidStateException is thrown when
// the server is in any of the following states: CREATING, TERMINATED, FAILED or
// DELETING.
func (c *Client) ExportServerEngineAttribute(ctx context.Context, params *ExportServerEngineAttributeInput, optFns ...func(*Options)) (*ExportServerEngineAttributeOutput, error) {
	if params == nil {
		params = &ExportServerEngineAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportServerEngineAttribute", params, optFns, c.addOperationExportServerEngineAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportServerEngineAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportServerEngineAttributeInput struct {

	// The name of the export attribute. Currently, the supported export attribute is
	// Userdata. This exports a user data script that includes parameters and values
	// provided in the InputAttributes list.
	//
	// This member is required.
	ExportAttributeName *string

	// The name of the server from which you are exporting the attribute.
	//
	// This member is required.
	ServerName *string

	// The list of engine attributes. The list type is EngineAttribute. An
	// EngineAttribute list item is a pair that includes an attribute name and its
	// value. For the Userdata ExportAttributeName, the following are supported engine
	// attribute names.
	//
	// * RunList In Chef, a list of roles or recipes that are run in
	// the specified order. In Puppet, this parameter is ignored.
	//
	// * OrganizationName
	// In Chef, an organization name. AWS OpsWorks for Chef Automate always creates the
	// organization default. In Puppet, this parameter is ignored.
	//
	// * NodeEnvironment
	// In Chef, a node environment (for example, development, staging, or one-box). In
	// Puppet, this parameter is ignored.
	//
	// * NodeClientVersion In Chef, the version of
	// the Chef engine (three numbers separated by dots, such as 13.8.5). If this
	// attribute is empty, OpsWorks for Chef Automate uses the most current version. In
	// Puppet, this parameter is ignored.
	InputAttributes []types.EngineAttribute

	noSmithyDocumentSerde
}

type ExportServerEngineAttributeOutput struct {

	// The requested engine attribute pair with attribute name and value.
	EngineAttribute *types.EngineAttribute

	// The server name used in the request.
	ServerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportServerEngineAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExportServerEngineAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExportServerEngineAttribute{}, middleware.After)
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
	if err = addOpExportServerEngineAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportServerEngineAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportServerEngineAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "ExportServerEngineAttribute",
	}
}
