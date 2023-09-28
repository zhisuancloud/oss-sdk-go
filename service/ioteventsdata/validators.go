// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"context"
	"fmt"
	"oss-sdk-go/service/ioteventsdata/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchAcknowledgeAlarm struct {
}

func (*validateOpBatchAcknowledgeAlarm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchAcknowledgeAlarm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchAcknowledgeAlarmInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchAcknowledgeAlarmInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchDeleteDetector struct {
}

func (*validateOpBatchDeleteDetector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchDeleteDetector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchDeleteDetectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchDeleteDetectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchDisableAlarm struct {
}

func (*validateOpBatchDisableAlarm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchDisableAlarm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchDisableAlarmInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchDisableAlarmInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchEnableAlarm struct {
}

func (*validateOpBatchEnableAlarm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchEnableAlarm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchEnableAlarmInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchEnableAlarmInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchPutMessage struct {
}

func (*validateOpBatchPutMessage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchPutMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchPutMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchPutMessageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchResetAlarm struct {
}

func (*validateOpBatchResetAlarm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchResetAlarm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchResetAlarmInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchResetAlarmInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchSnoozeAlarm struct {
}

func (*validateOpBatchSnoozeAlarm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchSnoozeAlarm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchSnoozeAlarmInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchSnoozeAlarmInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchUpdateDetector struct {
}

func (*validateOpBatchUpdateDetector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchUpdateDetector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchUpdateDetectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchUpdateDetectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAlarm struct {
}

func (*validateOpDescribeAlarm) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAlarm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAlarmInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAlarmInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeDetector struct {
}

func (*validateOpDescribeDetector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeDetector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeDetectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeDetectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAlarms struct {
}

func (*validateOpListAlarms) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAlarms) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAlarmsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAlarmsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDetectors struct {
}

func (*validateOpListDetectors) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDetectors) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDetectorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDetectorsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchAcknowledgeAlarmValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchAcknowledgeAlarm{}, middleware.After)
}

func addOpBatchDeleteDetectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchDeleteDetector{}, middleware.After)
}

func addOpBatchDisableAlarmValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchDisableAlarm{}, middleware.After)
}

func addOpBatchEnableAlarmValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchEnableAlarm{}, middleware.After)
}

func addOpBatchPutMessageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchPutMessage{}, middleware.After)
}

func addOpBatchResetAlarmValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchResetAlarm{}, middleware.After)
}

func addOpBatchSnoozeAlarmValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchSnoozeAlarm{}, middleware.After)
}

func addOpBatchUpdateDetectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchUpdateDetector{}, middleware.After)
}

func addOpDescribeAlarmValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAlarm{}, middleware.After)
}

func addOpDescribeDetectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeDetector{}, middleware.After)
}

func addOpListAlarmsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAlarms{}, middleware.After)
}

func addOpListDetectorsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDetectors{}, middleware.After)
}

func validateAcknowledgeAlarmActionRequest(v *types.AcknowledgeAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AcknowledgeAlarmActionRequest"}
	if v.RequestId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequestId"))
	}
	if v.AlarmModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlarmModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAcknowledgeAlarmActionRequests(v []types.AcknowledgeAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AcknowledgeAlarmActionRequests"}
	for i := range v {
		if err := validateAcknowledgeAlarmActionRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeleteDetectorRequest(v *types.DeleteDetectorRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDetectorRequest"}
	if v.MessageId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MessageId"))
	}
	if v.DetectorModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DetectorModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeleteDetectorRequests(v []types.DeleteDetectorRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDetectorRequests"}
	for i := range v {
		if err := validateDeleteDetectorRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDetectorStateDefinition(v *types.DetectorStateDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetectorStateDefinition"}
	if v.StateName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateName"))
	}
	if v.Variables == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Variables"))
	} else if v.Variables != nil {
		if err := validateVariableDefinitions(v.Variables); err != nil {
			invalidParams.AddNested("Variables", err.(smithy.InvalidParamsError))
		}
	}
	if v.Timers == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timers"))
	} else if v.Timers != nil {
		if err := validateTimerDefinitions(v.Timers); err != nil {
			invalidParams.AddNested("Timers", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDisableAlarmActionRequest(v *types.DisableAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisableAlarmActionRequest"}
	if v.RequestId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequestId"))
	}
	if v.AlarmModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlarmModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDisableAlarmActionRequests(v []types.DisableAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisableAlarmActionRequests"}
	for i := range v {
		if err := validateDisableAlarmActionRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEnableAlarmActionRequest(v *types.EnableAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EnableAlarmActionRequest"}
	if v.RequestId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequestId"))
	}
	if v.AlarmModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlarmModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEnableAlarmActionRequests(v []types.EnableAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EnableAlarmActionRequests"}
	for i := range v {
		if err := validateEnableAlarmActionRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessage(v *types.Message) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Message"}
	if v.MessageId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MessageId"))
	}
	if v.InputName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputName"))
	}
	if v.Payload == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Payload"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessages(v []types.Message) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Messages"}
	for i := range v {
		if err := validateMessage(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResetAlarmActionRequest(v *types.ResetAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResetAlarmActionRequest"}
	if v.RequestId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequestId"))
	}
	if v.AlarmModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlarmModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResetAlarmActionRequests(v []types.ResetAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResetAlarmActionRequests"}
	for i := range v {
		if err := validateResetAlarmActionRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSnoozeAlarmActionRequest(v *types.SnoozeAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SnoozeAlarmActionRequest"}
	if v.RequestId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequestId"))
	}
	if v.AlarmModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlarmModelName"))
	}
	if v.SnoozeDuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnoozeDuration"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSnoozeAlarmActionRequests(v []types.SnoozeAlarmActionRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SnoozeAlarmActionRequests"}
	for i := range v {
		if err := validateSnoozeAlarmActionRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTimerDefinition(v *types.TimerDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TimerDefinition"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Seconds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Seconds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTimerDefinitions(v []types.TimerDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TimerDefinitions"}
	for i := range v {
		if err := validateTimerDefinition(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateDetectorRequest(v *types.UpdateDetectorRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDetectorRequest"}
	if v.MessageId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MessageId"))
	}
	if v.DetectorModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DetectorModelName"))
	}
	if v.State == nil {
		invalidParams.Add(smithy.NewErrParamRequired("State"))
	} else if v.State != nil {
		if err := validateDetectorStateDefinition(v.State); err != nil {
			invalidParams.AddNested("State", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateDetectorRequests(v []types.UpdateDetectorRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDetectorRequests"}
	for i := range v {
		if err := validateUpdateDetectorRequest(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVariableDefinition(v *types.VariableDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableDefinition"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVariableDefinitions(v []types.VariableDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableDefinitions"}
	for i := range v {
		if err := validateVariableDefinition(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchAcknowledgeAlarmInput(v *BatchAcknowledgeAlarmInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchAcknowledgeAlarmInput"}
	if v.AcknowledgeActionRequests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcknowledgeActionRequests"))
	} else if v.AcknowledgeActionRequests != nil {
		if err := validateAcknowledgeAlarmActionRequests(v.AcknowledgeActionRequests); err != nil {
			invalidParams.AddNested("AcknowledgeActionRequests", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchDeleteDetectorInput(v *BatchDeleteDetectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchDeleteDetectorInput"}
	if v.Detectors == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Detectors"))
	} else if v.Detectors != nil {
		if err := validateDeleteDetectorRequests(v.Detectors); err != nil {
			invalidParams.AddNested("Detectors", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchDisableAlarmInput(v *BatchDisableAlarmInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchDisableAlarmInput"}
	if v.DisableActionRequests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DisableActionRequests"))
	} else if v.DisableActionRequests != nil {
		if err := validateDisableAlarmActionRequests(v.DisableActionRequests); err != nil {
			invalidParams.AddNested("DisableActionRequests", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchEnableAlarmInput(v *BatchEnableAlarmInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchEnableAlarmInput"}
	if v.EnableActionRequests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnableActionRequests"))
	} else if v.EnableActionRequests != nil {
		if err := validateEnableAlarmActionRequests(v.EnableActionRequests); err != nil {
			invalidParams.AddNested("EnableActionRequests", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchPutMessageInput(v *BatchPutMessageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchPutMessageInput"}
	if v.Messages == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Messages"))
	} else if v.Messages != nil {
		if err := validateMessages(v.Messages); err != nil {
			invalidParams.AddNested("Messages", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchResetAlarmInput(v *BatchResetAlarmInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchResetAlarmInput"}
	if v.ResetActionRequests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResetActionRequests"))
	} else if v.ResetActionRequests != nil {
		if err := validateResetAlarmActionRequests(v.ResetActionRequests); err != nil {
			invalidParams.AddNested("ResetActionRequests", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchSnoozeAlarmInput(v *BatchSnoozeAlarmInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchSnoozeAlarmInput"}
	if v.SnoozeActionRequests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnoozeActionRequests"))
	} else if v.SnoozeActionRequests != nil {
		if err := validateSnoozeAlarmActionRequests(v.SnoozeActionRequests); err != nil {
			invalidParams.AddNested("SnoozeActionRequests", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchUpdateDetectorInput(v *BatchUpdateDetectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchUpdateDetectorInput"}
	if v.Detectors == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Detectors"))
	} else if v.Detectors != nil {
		if err := validateUpdateDetectorRequests(v.Detectors); err != nil {
			invalidParams.AddNested("Detectors", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAlarmInput(v *DescribeAlarmInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAlarmInput"}
	if v.AlarmModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlarmModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeDetectorInput(v *DescribeDetectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeDetectorInput"}
	if v.DetectorModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DetectorModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAlarmsInput(v *ListAlarmsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAlarmsInput"}
	if v.AlarmModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlarmModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDetectorsInput(v *ListDetectorsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDetectorsInput"}
	if v.DetectorModelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DetectorModelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
