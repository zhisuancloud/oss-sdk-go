// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a report group. A report group contains a collection of reports.
func (c *Client) CreateReportGroup(ctx context.Context, params *CreateReportGroupInput, optFns ...func(*Options)) (*CreateReportGroupOutput, error) {
	if params == nil {
		params = &CreateReportGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReportGroup", params, optFns, c.addOperationCreateReportGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReportGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReportGroupInput struct {

	// A ReportExportConfig object that contains information about where the report
	// group test results are exported.
	//
	// This member is required.
	ExportConfig *types.ReportExportConfig

	// The name of the report group.
	//
	// This member is required.
	Name *string

	// The type of report group.
	//
	// This member is required.
	Type types.ReportType

	// A list of tag key and value pairs associated with this report group. These tags
	// are available for use by Amazon Web Services services that support CodeBuild
	// report group tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateReportGroupOutput struct {

	// Information about the report group that was created.
	ReportGroup *types.ReportGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReportGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReportGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReportGroup{}, middleware.After)
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
	if err = addOpCreateReportGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReportGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReportGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "CreateReportGroup",
	}
}
