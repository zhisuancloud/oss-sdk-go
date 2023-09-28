// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcampaigns

import (
	"context"
	"fmt"
	"oss-sdk-go/service/connectcampaigns/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateCampaign struct {
}

func (*validateOpCreateCampaign) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCampaign) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCampaignInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCampaignInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCampaign struct {
}

func (*validateOpDeleteCampaign) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCampaign) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCampaignInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCampaignInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteConnectInstanceConfig struct {
}

func (*validateOpDeleteConnectInstanceConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteConnectInstanceConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteConnectInstanceConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteConnectInstanceConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteInstanceOnboardingJob struct {
}

func (*validateOpDeleteInstanceOnboardingJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteInstanceOnboardingJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteInstanceOnboardingJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteInstanceOnboardingJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCampaign struct {
}

func (*validateOpDescribeCampaign) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCampaign) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeCampaignInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeCampaignInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCampaignStateBatch struct {
}

func (*validateOpGetCampaignStateBatch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCampaignStateBatch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCampaignStateBatchInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCampaignStateBatchInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCampaignState struct {
}

func (*validateOpGetCampaignState) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCampaignState) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCampaignStateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCampaignStateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetConnectInstanceConfig struct {
}

func (*validateOpGetConnectInstanceConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetConnectInstanceConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetConnectInstanceConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetConnectInstanceConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInstanceOnboardingJobStatus struct {
}

func (*validateOpGetInstanceOnboardingJobStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInstanceOnboardingJobStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInstanceOnboardingJobStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInstanceOnboardingJobStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListCampaigns struct {
}

func (*validateOpListCampaigns) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListCampaigns) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListCampaignsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListCampaignsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPauseCampaign struct {
}

func (*validateOpPauseCampaign) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPauseCampaign) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PauseCampaignInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPauseCampaignInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutDialRequestBatch struct {
}

func (*validateOpPutDialRequestBatch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutDialRequestBatch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutDialRequestBatchInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutDialRequestBatchInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpResumeCampaign struct {
}

func (*validateOpResumeCampaign) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpResumeCampaign) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ResumeCampaignInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpResumeCampaignInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartCampaign struct {
}

func (*validateOpStartCampaign) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartCampaign) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartCampaignInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartCampaignInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartInstanceOnboardingJob struct {
}

func (*validateOpStartInstanceOnboardingJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartInstanceOnboardingJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartInstanceOnboardingJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartInstanceOnboardingJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopCampaign struct {
}

func (*validateOpStopCampaign) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopCampaign) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopCampaignInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopCampaignInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCampaignDialerConfig struct {
}

func (*validateOpUpdateCampaignDialerConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCampaignDialerConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCampaignDialerConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCampaignDialerConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCampaignName struct {
}

func (*validateOpUpdateCampaignName) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCampaignName) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCampaignNameInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCampaignNameInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCampaignOutboundCallConfig struct {
}

func (*validateOpUpdateCampaignOutboundCallConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCampaignOutboundCallConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCampaignOutboundCallConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCampaignOutboundCallConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateCampaignValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCampaign{}, middleware.After)
}

func addOpDeleteCampaignValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCampaign{}, middleware.After)
}

func addOpDeleteConnectInstanceConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteConnectInstanceConfig{}, middleware.After)
}

func addOpDeleteInstanceOnboardingJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteInstanceOnboardingJob{}, middleware.After)
}

func addOpDescribeCampaignValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCampaign{}, middleware.After)
}

func addOpGetCampaignStateBatchValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCampaignStateBatch{}, middleware.After)
}

func addOpGetCampaignStateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCampaignState{}, middleware.After)
}

func addOpGetConnectInstanceConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetConnectInstanceConfig{}, middleware.After)
}

func addOpGetInstanceOnboardingJobStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInstanceOnboardingJobStatus{}, middleware.After)
}

func addOpListCampaignsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListCampaigns{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPauseCampaignValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPauseCampaign{}, middleware.After)
}

func addOpPutDialRequestBatchValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutDialRequestBatch{}, middleware.After)
}

func addOpResumeCampaignValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpResumeCampaign{}, middleware.After)
}

func addOpStartCampaignValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartCampaign{}, middleware.After)
}

func addOpStartInstanceOnboardingJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartInstanceOnboardingJob{}, middleware.After)
}

func addOpStopCampaignValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopCampaign{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateCampaignDialerConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCampaignDialerConfig{}, middleware.After)
}

func addOpUpdateCampaignNameValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCampaignName{}, middleware.After)
}

func addOpUpdateCampaignOutboundCallConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCampaignOutboundCallConfig{}, middleware.After)
}

func validateAnswerMachineDetectionConfig(v *types.AnswerMachineDetectionConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AnswerMachineDetectionConfig"}
	if v.EnableAnswerMachineDetection == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnableAnswerMachineDetection"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCampaignFilters(v *types.CampaignFilters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CampaignFilters"}
	if v.InstanceIdFilter != nil {
		if err := validateInstanceIdFilter(v.InstanceIdFilter); err != nil {
			invalidParams.AddNested("InstanceIdFilter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDialerConfig(v types.DialerConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DialerConfig"}
	switch uv := v.(type) {
	case *types.DialerConfigMemberPredictiveDialerConfig:
		if err := validatePredictiveDialerConfig(&uv.Value); err != nil {
			invalidParams.AddNested("[predictiveDialerConfig]", err.(smithy.InvalidParamsError))
		}

	case *types.DialerConfigMemberProgressiveDialerConfig:
		if err := validateProgressiveDialerConfig(&uv.Value); err != nil {
			invalidParams.AddNested("[progressiveDialerConfig]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDialRequest(v *types.DialRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DialRequest"}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.PhoneNumber == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PhoneNumber"))
	}
	if v.ExpirationTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExpirationTime"))
	}
	if v.Attributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDialRequestList(v []types.DialRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DialRequestList"}
	for i := range v {
		if err := validateDialRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEncryptionConfig(v *types.EncryptionConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EncryptionConfig"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInstanceIdFilter(v *types.InstanceIdFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InstanceIdFilter"}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if len(v.Operator) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Operator"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOutboundCallConfig(v *types.OutboundCallConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OutboundCallConfig"}
	if v.ConnectContactFlowId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectContactFlowId"))
	}
	if v.ConnectQueueId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectQueueId"))
	}
	if v.AnswerMachineDetectionConfig != nil {
		if err := validateAnswerMachineDetectionConfig(v.AnswerMachineDetectionConfig); err != nil {
			invalidParams.AddNested("AnswerMachineDetectionConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePredictiveDialerConfig(v *types.PredictiveDialerConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PredictiveDialerConfig"}
	if v.BandwidthAllocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BandwidthAllocation"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProgressiveDialerConfig(v *types.ProgressiveDialerConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProgressiveDialerConfig"}
	if v.BandwidthAllocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BandwidthAllocation"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCampaignInput(v *CreateCampaignInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCampaignInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ConnectInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectInstanceId"))
	}
	if v.DialerConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DialerConfig"))
	} else if v.DialerConfig != nil {
		if err := validateDialerConfig(v.DialerConfig); err != nil {
			invalidParams.AddNested("DialerConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutboundCallConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutboundCallConfig"))
	} else if v.OutboundCallConfig != nil {
		if err := validateOutboundCallConfig(v.OutboundCallConfig); err != nil {
			invalidParams.AddNested("OutboundCallConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCampaignInput(v *DeleteCampaignInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCampaignInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteConnectInstanceConfigInput(v *DeleteConnectInstanceConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteConnectInstanceConfigInput"}
	if v.ConnectInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteInstanceOnboardingJobInput(v *DeleteInstanceOnboardingJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteInstanceOnboardingJobInput"}
	if v.ConnectInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeCampaignInput(v *DescribeCampaignInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeCampaignInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCampaignStateBatchInput(v *GetCampaignStateBatchInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCampaignStateBatchInput"}
	if v.CampaignIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CampaignIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCampaignStateInput(v *GetCampaignStateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCampaignStateInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetConnectInstanceConfigInput(v *GetConnectInstanceConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetConnectInstanceConfigInput"}
	if v.ConnectInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInstanceOnboardingJobStatusInput(v *GetInstanceOnboardingJobStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInstanceOnboardingJobStatusInput"}
	if v.ConnectInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListCampaignsInput(v *ListCampaignsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListCampaignsInput"}
	if v.Filters != nil {
		if err := validateCampaignFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPauseCampaignInput(v *PauseCampaignInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PauseCampaignInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutDialRequestBatchInput(v *PutDialRequestBatchInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutDialRequestBatchInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.DialRequests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DialRequests"))
	} else if v.DialRequests != nil {
		if err := validateDialRequestList(v.DialRequests); err != nil {
			invalidParams.AddNested("DialRequests", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpResumeCampaignInput(v *ResumeCampaignInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResumeCampaignInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartCampaignInput(v *StartCampaignInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartCampaignInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartInstanceOnboardingJobInput(v *StartInstanceOnboardingJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartInstanceOnboardingJobInput"}
	if v.ConnectInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectInstanceId"))
	}
	if v.EncryptionConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EncryptionConfig"))
	} else if v.EncryptionConfig != nil {
		if err := validateEncryptionConfig(v.EncryptionConfig); err != nil {
			invalidParams.AddNested("EncryptionConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopCampaignInput(v *StopCampaignInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopCampaignInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCampaignDialerConfigInput(v *UpdateCampaignDialerConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCampaignDialerConfigInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.DialerConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DialerConfig"))
	} else if v.DialerConfig != nil {
		if err := validateDialerConfig(v.DialerConfig); err != nil {
			invalidParams.AddNested("DialerConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCampaignNameInput(v *UpdateCampaignNameInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCampaignNameInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCampaignOutboundCallConfigInput(v *UpdateCampaignOutboundCallConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCampaignOutboundCallConfigInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.AnswerMachineDetectionConfig != nil {
		if err := validateAnswerMachineDetectionConfig(v.AnswerMachineDetectionConfig); err != nil {
			invalidParams.AddNested("AnswerMachineDetectionConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
