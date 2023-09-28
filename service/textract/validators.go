// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	"oss-sdk-go/service/textract/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAnalyzeDocument struct {
}

func (*validateOpAnalyzeDocument) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAnalyzeDocument) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AnalyzeDocumentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAnalyzeDocumentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAnalyzeExpense struct {
}

func (*validateOpAnalyzeExpense) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAnalyzeExpense) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AnalyzeExpenseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAnalyzeExpenseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAnalyzeID struct {
}

func (*validateOpAnalyzeID) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAnalyzeID) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AnalyzeIDInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAnalyzeIDInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDetectDocumentText struct {
}

func (*validateOpDetectDocumentText) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDetectDocumentText) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DetectDocumentTextInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDetectDocumentTextInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDocumentAnalysis struct {
}

func (*validateOpGetDocumentAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDocumentAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDocumentAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDocumentAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDocumentTextDetection struct {
}

func (*validateOpGetDocumentTextDetection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDocumentTextDetection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDocumentTextDetectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDocumentTextDetectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExpenseAnalysis struct {
}

func (*validateOpGetExpenseAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExpenseAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExpenseAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExpenseAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDocumentAnalysis struct {
}

func (*validateOpStartDocumentAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDocumentAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDocumentAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDocumentAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDocumentTextDetection struct {
}

func (*validateOpStartDocumentTextDetection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDocumentTextDetection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDocumentTextDetectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDocumentTextDetectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartExpenseAnalysis struct {
}

func (*validateOpStartExpenseAnalysis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartExpenseAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartExpenseAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartExpenseAnalysisInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAnalyzeDocumentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAnalyzeDocument{}, middleware.After)
}

func addOpAnalyzeExpenseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAnalyzeExpense{}, middleware.After)
}

func addOpAnalyzeIDValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAnalyzeID{}, middleware.After)
}

func addOpDetectDocumentTextValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDetectDocumentText{}, middleware.After)
}

func addOpGetDocumentAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDocumentAnalysis{}, middleware.After)
}

func addOpGetDocumentTextDetectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDocumentTextDetection{}, middleware.After)
}

func addOpGetExpenseAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExpenseAnalysis{}, middleware.After)
}

func addOpStartDocumentAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDocumentAnalysis{}, middleware.After)
}

func addOpStartDocumentTextDetectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDocumentTextDetection{}, middleware.After)
}

func addOpStartExpenseAnalysisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartExpenseAnalysis{}, middleware.After)
}

func validateHumanLoopConfig(v *types.HumanLoopConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HumanLoopConfig"}
	if v.HumanLoopName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HumanLoopName"))
	}
	if v.FlowDefinitionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlowDefinitionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNotificationChannel(v *types.NotificationChannel) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NotificationChannel"}
	if v.SNSTopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SNSTopicArn"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOutputConfig(v *types.OutputConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OutputConfig"}
	if v.S3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Bucket"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateQueries(v []types.Query) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Queries"}
	for i := range v {
		if err := validateQuery(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateQueriesConfig(v *types.QueriesConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueriesConfig"}
	if v.Queries == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Queries"))
	} else if v.Queries != nil {
		if err := validateQueries(v.Queries); err != nil {
			invalidParams.AddNested("Queries", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateQuery(v *types.Query) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Query"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAnalyzeDocumentInput(v *AnalyzeDocumentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AnalyzeDocumentInput"}
	if v.Document == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Document"))
	}
	if v.FeatureTypes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureTypes"))
	}
	if v.HumanLoopConfig != nil {
		if err := validateHumanLoopConfig(v.HumanLoopConfig); err != nil {
			invalidParams.AddNested("HumanLoopConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.QueriesConfig != nil {
		if err := validateQueriesConfig(v.QueriesConfig); err != nil {
			invalidParams.AddNested("QueriesConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAnalyzeExpenseInput(v *AnalyzeExpenseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AnalyzeExpenseInput"}
	if v.Document == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Document"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAnalyzeIDInput(v *AnalyzeIDInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AnalyzeIDInput"}
	if v.DocumentPages == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentPages"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDetectDocumentTextInput(v *DetectDocumentTextInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetectDocumentTextInput"}
	if v.Document == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Document"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDocumentAnalysisInput(v *GetDocumentAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDocumentAnalysisInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDocumentTextDetectionInput(v *GetDocumentTextDetectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDocumentTextDetectionInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExpenseAnalysisInput(v *GetExpenseAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExpenseAnalysisInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDocumentAnalysisInput(v *StartDocumentAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDocumentAnalysisInput"}
	if v.DocumentLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentLocation"))
	}
	if v.FeatureTypes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureTypes"))
	}
	if v.NotificationChannel != nil {
		if err := validateNotificationChannel(v.NotificationChannel); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputConfig != nil {
		if err := validateOutputConfig(v.OutputConfig); err != nil {
			invalidParams.AddNested("OutputConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.QueriesConfig != nil {
		if err := validateQueriesConfig(v.QueriesConfig); err != nil {
			invalidParams.AddNested("QueriesConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDocumentTextDetectionInput(v *StartDocumentTextDetectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDocumentTextDetectionInput"}
	if v.DocumentLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentLocation"))
	}
	if v.NotificationChannel != nil {
		if err := validateNotificationChannel(v.NotificationChannel); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputConfig != nil {
		if err := validateOutputConfig(v.OutputConfig); err != nil {
			invalidParams.AddNested("OutputConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartExpenseAnalysisInput(v *StartExpenseAnalysisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartExpenseAnalysisInput"}
	if v.DocumentLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentLocation"))
	}
	if v.NotificationChannel != nil {
		if err := validateNotificationChannel(v.NotificationChannel); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputConfig != nil {
		if err := validateOutputConfig(v.OutputConfig); err != nil {
			invalidParams.AddNested("OutputConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
