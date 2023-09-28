// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdeviceadvisor

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a report download link for a successful Device Advisor qualifying test
// suite run. Requires permission to access the GetSuiteRunReport
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) GetSuiteRunReport(ctx context.Context, params *GetSuiteRunReportInput, optFns ...func(*Options)) (*GetSuiteRunReportOutput, error) {
	if params == nil {
		params = &GetSuiteRunReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSuiteRunReport", params, optFns, c.addOperationGetSuiteRunReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSuiteRunReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSuiteRunReportInput struct {

	// Suite definition ID of the test suite.
	//
	// This member is required.
	SuiteDefinitionId *string

	// Suite run ID of the test suite run.
	//
	// This member is required.
	SuiteRunId *string

	noSmithyDocumentSerde
}

type GetSuiteRunReportOutput struct {

	// Download URL of the qualification report.
	QualificationReportDownloadUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSuiteRunReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSuiteRunReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSuiteRunReport{}, middleware.After)
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
	if err = addOpGetSuiteRunReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSuiteRunReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSuiteRunReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdeviceadvisor",
		OperationName: "GetSuiteRunReport",
	}
}
