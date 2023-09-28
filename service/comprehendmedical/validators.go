// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	"oss-sdk-go/service/comprehendmedical/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDescribeEntitiesDetectionV2Job struct {
}

func (*validateOpDescribeEntitiesDetectionV2Job) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeEntitiesDetectionV2Job) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeEntitiesDetectionV2JobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeEntitiesDetectionV2JobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeICD10CMInferenceJob struct {
}

func (*validateOpDescribeICD10CMInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeICD10CMInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeICD10CMInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeICD10CMInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribePHIDetectionJob struct {
}

func (*validateOpDescribePHIDetectionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribePHIDetectionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribePHIDetectionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribePHIDetectionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeRxNormInferenceJob struct {
}

func (*validateOpDescribeRxNormInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeRxNormInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeRxNormInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeRxNormInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeSNOMEDCTInferenceJob struct {
}

func (*validateOpDescribeSNOMEDCTInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeSNOMEDCTInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeSNOMEDCTInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeSNOMEDCTInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDetectEntities struct {
}

func (*validateOpDetectEntities) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDetectEntities) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DetectEntitiesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDetectEntitiesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDetectEntitiesV2 struct {
}

func (*validateOpDetectEntitiesV2) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDetectEntitiesV2) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DetectEntitiesV2Input)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDetectEntitiesV2Input(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDetectPHI struct {
}

func (*validateOpDetectPHI) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDetectPHI) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DetectPHIInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDetectPHIInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInferICD10CM struct {
}

func (*validateOpInferICD10CM) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInferICD10CM) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InferICD10CMInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInferICD10CMInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInferRxNorm struct {
}

func (*validateOpInferRxNorm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInferRxNorm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InferRxNormInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInferRxNormInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInferSNOMEDCT struct {
}

func (*validateOpInferSNOMEDCT) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInferSNOMEDCT) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InferSNOMEDCTInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInferSNOMEDCTInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartEntitiesDetectionV2Job struct {
}

func (*validateOpStartEntitiesDetectionV2Job) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartEntitiesDetectionV2Job) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartEntitiesDetectionV2JobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartEntitiesDetectionV2JobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartICD10CMInferenceJob struct {
}

func (*validateOpStartICD10CMInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartICD10CMInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartICD10CMInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartICD10CMInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartPHIDetectionJob struct {
}

func (*validateOpStartPHIDetectionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartPHIDetectionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartPHIDetectionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartPHIDetectionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartRxNormInferenceJob struct {
}

func (*validateOpStartRxNormInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartRxNormInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartRxNormInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartRxNormInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartSNOMEDCTInferenceJob struct {
}

func (*validateOpStartSNOMEDCTInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartSNOMEDCTInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartSNOMEDCTInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartSNOMEDCTInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopEntitiesDetectionV2Job struct {
}

func (*validateOpStopEntitiesDetectionV2Job) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopEntitiesDetectionV2Job) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopEntitiesDetectionV2JobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopEntitiesDetectionV2JobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopICD10CMInferenceJob struct {
}

func (*validateOpStopICD10CMInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopICD10CMInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopICD10CMInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopICD10CMInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopPHIDetectionJob struct {
}

func (*validateOpStopPHIDetectionJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopPHIDetectionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopPHIDetectionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopPHIDetectionJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopRxNormInferenceJob struct {
}

func (*validateOpStopRxNormInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopRxNormInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopRxNormInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopRxNormInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopSNOMEDCTInferenceJob struct {
}

func (*validateOpStopSNOMEDCTInferenceJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopSNOMEDCTInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopSNOMEDCTInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopSNOMEDCTInferenceJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDescribeEntitiesDetectionV2JobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeEntitiesDetectionV2Job{}, middleware.After)
}

func addOpDescribeICD10CMInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeICD10CMInferenceJob{}, middleware.After)
}

func addOpDescribePHIDetectionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribePHIDetectionJob{}, middleware.After)
}

func addOpDescribeRxNormInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeRxNormInferenceJob{}, middleware.After)
}

func addOpDescribeSNOMEDCTInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeSNOMEDCTInferenceJob{}, middleware.After)
}

func addOpDetectEntitiesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDetectEntities{}, middleware.After)
}

func addOpDetectEntitiesV2ValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDetectEntitiesV2{}, middleware.After)
}

func addOpDetectPHIValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDetectPHI{}, middleware.After)
}

func addOpInferICD10CMValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInferICD10CM{}, middleware.After)
}

func addOpInferRxNormValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInferRxNorm{}, middleware.After)
}

func addOpInferSNOMEDCTValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInferSNOMEDCT{}, middleware.After)
}

func addOpStartEntitiesDetectionV2JobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartEntitiesDetectionV2Job{}, middleware.After)
}

func addOpStartICD10CMInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartICD10CMInferenceJob{}, middleware.After)
}

func addOpStartPHIDetectionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartPHIDetectionJob{}, middleware.After)
}

func addOpStartRxNormInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartRxNormInferenceJob{}, middleware.After)
}

func addOpStartSNOMEDCTInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartSNOMEDCTInferenceJob{}, middleware.After)
}

func addOpStopEntitiesDetectionV2JobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopEntitiesDetectionV2Job{}, middleware.After)
}

func addOpStopICD10CMInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopICD10CMInferenceJob{}, middleware.After)
}

func addOpStopPHIDetectionJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopPHIDetectionJob{}, middleware.After)
}

func addOpStopRxNormInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopRxNormInferenceJob{}, middleware.After)
}

func addOpStopSNOMEDCTInferenceJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopSNOMEDCTInferenceJob{}, middleware.After)
}

func validateInputDataConfig(v *types.InputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InputDataConfig"}
	if v.S3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Bucket"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOutputDataConfig(v *types.OutputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OutputDataConfig"}
	if v.S3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Bucket"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeEntitiesDetectionV2JobInput(v *DescribeEntitiesDetectionV2JobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeEntitiesDetectionV2JobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeICD10CMInferenceJobInput(v *DescribeICD10CMInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeICD10CMInferenceJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribePHIDetectionJobInput(v *DescribePHIDetectionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribePHIDetectionJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeRxNormInferenceJobInput(v *DescribeRxNormInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeRxNormInferenceJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeSNOMEDCTInferenceJobInput(v *DescribeSNOMEDCTInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeSNOMEDCTInferenceJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDetectEntitiesInput(v *DetectEntitiesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetectEntitiesInput"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDetectEntitiesV2Input(v *DetectEntitiesV2Input) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetectEntitiesV2Input"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDetectPHIInput(v *DetectPHIInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetectPHIInput"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInferICD10CMInput(v *InferICD10CMInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InferICD10CMInput"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInferRxNormInput(v *InferRxNormInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InferRxNormInput"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInferSNOMEDCTInput(v *InferSNOMEDCTInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InferSNOMEDCTInput"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartEntitiesDetectionV2JobInput(v *StartEntitiesDetectionV2JobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartEntitiesDetectionV2JobInput"}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartICD10CMInferenceJobInput(v *StartICD10CMInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartICD10CMInferenceJobInput"}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartPHIDetectionJobInput(v *StartPHIDetectionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartPHIDetectionJobInput"}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartRxNormInferenceJobInput(v *StartRxNormInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartRxNormInferenceJobInput"}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartSNOMEDCTInferenceJobInput(v *StartSNOMEDCTInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartSNOMEDCTInferenceJobInput"}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopEntitiesDetectionV2JobInput(v *StopEntitiesDetectionV2JobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopEntitiesDetectionV2JobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopICD10CMInferenceJobInput(v *StopICD10CMInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopICD10CMInferenceJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopPHIDetectionJobInput(v *StopPHIDetectionJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopPHIDetectionJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopRxNormInferenceJobInput(v *StopRxNormInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopRxNormInferenceJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopSNOMEDCTInferenceJobInput(v *StopSNOMEDCTInferenceJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopSNOMEDCTInferenceJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
