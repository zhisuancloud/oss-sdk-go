// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables an organizational policy type in a root. A policy of a certain type can
// be attached to entities in a root only if that type is enabled in the root.
// After you perform this operation, you no longer can attach policies of the
// specified type to that root or to any organizational unit (OU) or account in
// that root. You can undo this by using the EnablePolicyType operation. This is an
// asynchronous request that Amazon Web Services performs in the background. If you
// disable a policy type for a root, it still appears enabled for the organization
// if all features
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
// are enabled for the organization. Amazon Web Services recommends that you first
// use ListRoots to see the status of policy types for a specified root, and then
// use this operation. This operation can be called only from the organization's
// management account. To view the status of available policy types in the
// organization, use DescribeOrganization.
func (c *Client) DisablePolicyType(ctx context.Context, params *DisablePolicyTypeInput, optFns ...func(*Options)) (*DisablePolicyTypeOutput, error) {
	if params == nil {
		params = &DisablePolicyTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisablePolicyType", params, optFns, c.addOperationDisablePolicyTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisablePolicyTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisablePolicyTypeInput struct {

	// The policy type that you want to disable in this root. You can specify one of
	// the following values:
	//
	// * AISERVICES_OPT_OUT_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	//
	// *
	// BACKUP_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
	//
	// *
	// SERVICE_CONTROL_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html)
	//
	// *
	// TAG_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
	//
	// This member is required.
	PolicyType types.PolicyType

	// The unique identifier (ID) of the root in which you want to disable a policy
	// type. You can get the ID from the ListRoots operation. The regex pattern
	// (http://wikipedia.org/wiki/regex) for a root ID string requires "r-" followed by
	// from 4 to 32 lowercase letters or digits.
	//
	// This member is required.
	RootId *string

	noSmithyDocumentSerde
}

type DisablePolicyTypeOutput struct {

	// A structure that shows the root with the updated list of enabled policy types.
	Root *types.Root

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisablePolicyTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisablePolicyType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisablePolicyType{}, middleware.After)
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
	if err = addOpDisablePolicyTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisablePolicyType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisablePolicyType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DisablePolicyType",
	}
}
