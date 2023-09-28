// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of severity levels that you can assign to a support case. The
// severity level for a case is also a field in the CaseDetails data type that you
// include for a CreateCase request.
//
// * You must have a Business, Enterprise
// On-Ramp, or Enterprise Support plan to use the Amazon Web Services Support
// API.
//
// * If you call the Amazon Web Services Support API from an account that
// does not have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see Amazon Web Services Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeSeverityLevels(ctx context.Context, params *DescribeSeverityLevelsInput, optFns ...func(*Options)) (*DescribeSeverityLevelsOutput, error) {
	if params == nil {
		params = &DescribeSeverityLevelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSeverityLevels", params, optFns, c.addOperationDescribeSeverityLevelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSeverityLevelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSeverityLevelsInput struct {

	// The ISO 639-1 code for the language in which Amazon Web Services provides
	// support. Amazon Web Services Support currently supports English ("en") and
	// Japanese ("ja"). Language parameters must be passed explicitly for operations
	// that take them.
	Language *string

	noSmithyDocumentSerde
}

// The list of severity levels returned by the DescribeSeverityLevels operation.
type DescribeSeverityLevelsOutput struct {

	// The available severity levels for the support case. Available severity levels
	// are defined by your service level agreement with Amazon Web Services.
	SeverityLevels []types.SeverityLevel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSeverityLevelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSeverityLevels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSeverityLevels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSeverityLevels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSeverityLevels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeSeverityLevels",
	}
}