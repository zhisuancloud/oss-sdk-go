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

// Returns information about all available Trusted Advisor checks, including the
// name, ID, category, description, and metadata. You must specify a language code.
// The Amazon Web Services Support API currently supports English ("en") and
// Japanese ("ja"). The response contains a TrustedAdvisorCheckDescription object
// for each check. You must set the Amazon Web Services Region to us-east-1.
//
// * You
// must have a Business, Enterprise On-Ramp, or Enterprise Support plan to use the
// Amazon Web Services Support API.
//
// * If you call the Amazon Web Services Support
// API from an account that does not have a Business, Enterprise On-Ramp, or
// Enterprise Support plan, the SubscriptionRequiredException error message
// appears. For information about changing your support plan, see Amazon Web
// Services Support (http://aws.amazon.com/premiumsupport/).
//
// * The names and
// descriptions for Trusted Advisor checks are subject to change. We recommend that
// you specify the check ID in your code to uniquely identify a check.
func (c *Client) DescribeTrustedAdvisorChecks(ctx context.Context, params *DescribeTrustedAdvisorChecksInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorChecksOutput, error) {
	if params == nil {
		params = &DescribeTrustedAdvisorChecksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrustedAdvisorChecks", params, optFns, c.addOperationDescribeTrustedAdvisorChecksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrustedAdvisorChecksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrustedAdvisorChecksInput struct {

	// The ISO 639-1 code for the language in which Amazon Web Services provides
	// support. Amazon Web Services Support currently supports English ("en") and
	// Japanese ("ja"). Language parameters must be passed explicitly for operations
	// that take them.
	//
	// This member is required.
	Language *string

	noSmithyDocumentSerde
}

// Information about the Trusted Advisor checks returned by the
// DescribeTrustedAdvisorChecks operation.
type DescribeTrustedAdvisorChecksOutput struct {

	// Information about all available Trusted Advisor checks.
	//
	// This member is required.
	Checks []types.TrustedAdvisorCheckDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrustedAdvisorChecksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrustedAdvisorChecks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrustedAdvisorChecks{}, middleware.After)
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
	if err = addOpDescribeTrustedAdvisorChecksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrustedAdvisorChecks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTrustedAdvisorChecks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeTrustedAdvisorChecks",
	}
}
