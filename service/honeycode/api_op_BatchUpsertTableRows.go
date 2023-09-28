// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/honeycode/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The BatchUpsertTableRows API allows you to upsert one or more rows in a table.
// The upsert operation takes a filter expression as input and evaluates it to find
// matching rows on the destination table. If matching rows are found, it will
// update the cells in the matching rows to new values specified in the request. If
// no matching rows are found, a new row is added at the end of the table and the
// cells in that row are set to the new values specified in the request. You can
// specify the values to set in some or all of the columns in the table for the
// matching or newly appended rows. If a column is not explicitly specified for a
// particular row, then that column will not be updated for that row. To clear out
// the data in a specific cell, you need to set the value as an empty string ("").
func (c *Client) BatchUpsertTableRows(ctx context.Context, params *BatchUpsertTableRowsInput, optFns ...func(*Options)) (*BatchUpsertTableRowsOutput, error) {
	if params == nil {
		params = &BatchUpsertTableRowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpsertTableRows", params, optFns, c.addOperationBatchUpsertTableRowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpsertTableRowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpsertTableRowsInput struct {

	// The list of rows to upsert in the table. Each item in this list needs to have a
	// batch item id to uniquely identify the element in the request, a filter
	// expression to find the rows to update for that element and the cell values to
	// set for each column in the upserted rows. You need to specify at least one item
	// in this list. Note that if one of the filter formulas in the request fails to
	// evaluate because of an error or one of the column ids in any of the rows does
	// not exist in the table, then the request fails and no updates are made to the
	// table.
	//
	// This member is required.
	RowsToUpsert []types.UpsertRowData

	// The ID of the table where the rows are being upserted. If a table with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	TableId *string

	// The ID of the workbook where the rows are being upserted. If a workbook with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	// The request token for performing the update action. Request tokens help to
	// identify duplicate requests. If a call times out or fails due to a transient
	// error like a failed network connection, you can retry the call with the same
	// request token. The service ensures that if the first call using that request
	// token is successfully performed, the second call will not perform the action
	// again. Note that request tokens are valid only for a few minutes. You cannot use
	// request tokens to dedupe requests spanning hours or days.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type BatchUpsertTableRowsOutput struct {

	// A map with the batch item id as the key and the result of the upsert operation
	// as the value. The result of the upsert operation specifies whether existing rows
	// were updated or a new row was appended, along with the list of row ids that were
	// affected.
	//
	// This member is required.
	Rows map[string]types.UpsertRowsResult

	// The updated workbook cursor after updating or appending rows in the table.
	//
	// This member is required.
	WorkbookCursor int64

	// The list of batch items in the request that could not be updated or appended in
	// the table. Each element in this list contains one item from the request that
	// could not be updated in the table along with the reason why that item could not
	// be updated or appended.
	FailedBatchItems []types.FailedBatchItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpsertTableRowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpsertTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpsertTableRows{}, middleware.After)
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
	if err = addOpBatchUpsertTableRowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpsertTableRows(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpsertTableRows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "honeycode",
		OperationName: "BatchUpsertTableRows",
	}
}
