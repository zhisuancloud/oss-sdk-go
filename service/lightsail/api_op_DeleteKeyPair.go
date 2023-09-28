// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified key pair by removing the public key from Amazon Lightsail.
// You can delete key pairs that were created using the ImportKeyPair
// (https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ImportKeyPair.html)
// and CreateKeyPair
// (https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_CreateKeyPair.html)
// actions, as well as the Lightsail default key pair. A new default key pair will
// not be created unless you launch an instance without specifying a custom key
// pair, or you call the DownloadDefaultKeyPair
// (https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DownloadDefaultKeyPair.html)
// API. The delete key pair operation supports tag-based access control via
// resource tags applied to the resource identified by key pair name. For more
// information, see the Amazon Lightsail Developer Guide
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) DeleteKeyPair(ctx context.Context, params *DeleteKeyPairInput, optFns ...func(*Options)) (*DeleteKeyPairOutput, error) {
	if params == nil {
		params = &DeleteKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteKeyPair", params, optFns, c.addOperationDeleteKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKeyPairInput struct {

	// The name of the key pair to delete.
	//
	// This member is required.
	KeyPairName *string

	// The RSA fingerprint of the Lightsail default key pair to delete. The
	// expectedFingerprint parameter is required only when specifying to delete a
	// Lightsail default key pair.
	ExpectedFingerprint *string

	noSmithyDocumentSerde
}

type DeleteKeyPairOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteKeyPair{}, middleware.After)
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
	if err = addOpDeleteKeyPairValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKeyPair(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteKeyPair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DeleteKeyPair",
	}
}
