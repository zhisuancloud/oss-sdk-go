// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the receipt rule sets that exist under your AWS account in the current AWS
// Region. If there are additional receipt rule sets to be retrieved, you will
// receive a NextToken that you can provide to the next call to ListReceiptRuleSets
// to retrieve the additional entries. For information about managing receipt rule
// sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rule-sets.html).
// You can execute this operation no more than once per second.
func (c *Client) ListReceiptRuleSets(ctx context.Context, params *ListReceiptRuleSetsInput, optFns ...func(*Options)) (*ListReceiptRuleSetsOutput, error) {
	if params == nil {
		params = &ListReceiptRuleSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReceiptRuleSets", params, optFns, c.addOperationListReceiptRuleSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReceiptRuleSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the receipt rule sets that exist under your AWS
// account. You use receipt rule sets to receive email with Amazon SES. For more
// information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type ListReceiptRuleSetsInput struct {

	// A token returned from a previous call to ListReceiptRuleSets to indicate the
	// position in the receipt rule set list.
	NextToken *string

	noSmithyDocumentSerde
}

// A list of receipt rule sets that exist under your AWS account.
type ListReceiptRuleSetsOutput struct {

	// A token indicating that there are additional receipt rule sets available to be
	// listed. Pass this token to successive calls of ListReceiptRuleSets to retrieve
	// up to 100 receipt rule sets at a time.
	NextToken *string

	// The metadata for the currently active receipt rule set. The metadata consists of
	// the rule set name and the timestamp of when the rule set was created.
	RuleSets []types.ReceiptRuleSetMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReceiptRuleSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListReceiptRuleSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListReceiptRuleSets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReceiptRuleSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListReceiptRuleSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListReceiptRuleSets",
	}
}
