// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Export an existing lens. Lenses are defined in JSON. For more information, see
// JSON format specification
// (https://docs.aws.amazon.com/wellarchitected/latest/userguide/lenses-format-specification.html)
// in the Well-Architected Tool User Guide. Only the owner of a lens can export it.
// Disclaimer Do not include or gather personal identifiable information (PII) of
// end users or other identifiable individuals in or via your custom lenses. If
// your custom lens or those shared with you and used in your account do include or
// collect PII you are responsible for: ensuring that the included PII is processed
// in accordance with applicable law, providing adequate privacy notices, and
// obtaining necessary consents for processing such data.
func (c *Client) ExportLens(ctx context.Context, params *ExportLensInput, optFns ...func(*Options)) (*ExportLensOutput, error) {
	if params == nil {
		params = &ExportLensInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportLens", params, optFns, c.addOperationExportLensMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportLensOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportLensInput struct {

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless, or the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2::lens/serverless. For custom lenses, this is
	// the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1:123456789012:lens/my-lens. Each lens is
	// identified by its LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The lens version to be exported.
	LensVersion *string

	noSmithyDocumentSerde
}

type ExportLensOutput struct {

	// The JSON for the lens.
	LensJSON *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportLensMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportLens{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportLens{}, middleware.After)
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
	if err = addOpExportLensValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportLens(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportLens(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ExportLens",
	}
}