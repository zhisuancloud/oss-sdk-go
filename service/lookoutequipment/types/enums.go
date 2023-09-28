// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DatasetStatus string

// Enum values for DatasetStatus
const (
	DatasetStatusCreated             DatasetStatus = "CREATED"
	DatasetStatusIngestionInProgress DatasetStatus = "INGESTION_IN_PROGRESS"
	DatasetStatusActive              DatasetStatus = "ACTIVE"
)

// Values returns all known values for DatasetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DatasetStatus) Values() []DatasetStatus {
	return []DatasetStatus{
		"CREATED",
		"INGESTION_IN_PROGRESS",
		"ACTIVE",
	}
}

type DataUploadFrequency string

// Enum values for DataUploadFrequency
const (
	DataUploadFrequencyPt5m  DataUploadFrequency = "PT5M"
	DataUploadFrequencyPt10m DataUploadFrequency = "PT10M"
	DataUploadFrequencyPt15m DataUploadFrequency = "PT15M"
	DataUploadFrequencyPt30m DataUploadFrequency = "PT30M"
	DataUploadFrequencyPt1h  DataUploadFrequency = "PT1H"
)

// Values returns all known values for DataUploadFrequency. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataUploadFrequency) Values() []DataUploadFrequency {
	return []DataUploadFrequency{
		"PT5M",
		"PT10M",
		"PT15M",
		"PT30M",
		"PT1H",
	}
}

type InferenceExecutionStatus string

// Enum values for InferenceExecutionStatus
const (
	InferenceExecutionStatusInProgress InferenceExecutionStatus = "IN_PROGRESS"
	InferenceExecutionStatusSuccess    InferenceExecutionStatus = "SUCCESS"
	InferenceExecutionStatusFailed     InferenceExecutionStatus = "FAILED"
)

// Values returns all known values for InferenceExecutionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InferenceExecutionStatus) Values() []InferenceExecutionStatus {
	return []InferenceExecutionStatus{
		"IN_PROGRESS",
		"SUCCESS",
		"FAILED",
	}
}

type InferenceSchedulerStatus string

// Enum values for InferenceSchedulerStatus
const (
	InferenceSchedulerStatusPending  InferenceSchedulerStatus = "PENDING"
	InferenceSchedulerStatusRunning  InferenceSchedulerStatus = "RUNNING"
	InferenceSchedulerStatusStopping InferenceSchedulerStatus = "STOPPING"
	InferenceSchedulerStatusStopped  InferenceSchedulerStatus = "STOPPED"
)

// Values returns all known values for InferenceSchedulerStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InferenceSchedulerStatus) Values() []InferenceSchedulerStatus {
	return []InferenceSchedulerStatus{
		"PENDING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
	}
}

type IngestionJobStatus string

// Enum values for IngestionJobStatus
const (
	IngestionJobStatusInProgress IngestionJobStatus = "IN_PROGRESS"
	IngestionJobStatusSuccess    IngestionJobStatus = "SUCCESS"
	IngestionJobStatusFailed     IngestionJobStatus = "FAILED"
)

// Values returns all known values for IngestionJobStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobStatus) Values() []IngestionJobStatus {
	return []IngestionJobStatus{
		"IN_PROGRESS",
		"SUCCESS",
		"FAILED",
	}
}

type LabelRating string

// Enum values for LabelRating
const (
	LabelRatingAnomaly   LabelRating = "ANOMALY"
	LabelRatingNoAnomaly LabelRating = "NO_ANOMALY"
	LabelRatingNeutral   LabelRating = "NEUTRAL"
)

// Values returns all known values for LabelRating. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LabelRating) Values() []LabelRating {
	return []LabelRating{
		"ANOMALY",
		"NO_ANOMALY",
		"NEUTRAL",
	}
}

type LatestInferenceResult string

// Enum values for LatestInferenceResult
const (
	LatestInferenceResultAnomalous LatestInferenceResult = "ANOMALOUS"
	LatestInferenceResultNormal    LatestInferenceResult = "NORMAL"
)

// Values returns all known values for LatestInferenceResult. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LatestInferenceResult) Values() []LatestInferenceResult {
	return []LatestInferenceResult{
		"ANOMALOUS",
		"NORMAL",
	}
}

type ModelStatus string

// Enum values for ModelStatus
const (
	ModelStatusInProgress ModelStatus = "IN_PROGRESS"
	ModelStatusSuccess    ModelStatus = "SUCCESS"
	ModelStatusFailed     ModelStatus = "FAILED"
)

// Values returns all known values for ModelStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ModelStatus) Values() []ModelStatus {
	return []ModelStatus{
		"IN_PROGRESS",
		"SUCCESS",
		"FAILED",
	}
}

type Monotonicity string

// Enum values for Monotonicity
const (
	MonotonicityDecreasing Monotonicity = "DECREASING"
	MonotonicityIncreasing Monotonicity = "INCREASING"
	MonotonicityStatic     Monotonicity = "STATIC"
)

// Values returns all known values for Monotonicity. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Monotonicity) Values() []Monotonicity {
	return []Monotonicity{
		"DECREASING",
		"INCREASING",
		"STATIC",
	}
}

type StatisticalIssueStatus string

// Enum values for StatisticalIssueStatus
const (
	StatisticalIssueStatusPotentialIssueDetected StatisticalIssueStatus = "POTENTIAL_ISSUE_DETECTED"
	StatisticalIssueStatusNoIssueDetected        StatisticalIssueStatus = "NO_ISSUE_DETECTED"
)

// Values returns all known values for StatisticalIssueStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StatisticalIssueStatus) Values() []StatisticalIssueStatus {
	return []StatisticalIssueStatus{
		"POTENTIAL_ISSUE_DETECTED",
		"NO_ISSUE_DETECTED",
	}
}

type TargetSamplingRate string

// Enum values for TargetSamplingRate
const (
	TargetSamplingRatePt1s  TargetSamplingRate = "PT1S"
	TargetSamplingRatePt5s  TargetSamplingRate = "PT5S"
	TargetSamplingRatePt10s TargetSamplingRate = "PT10S"
	TargetSamplingRatePt15s TargetSamplingRate = "PT15S"
	TargetSamplingRatePt30s TargetSamplingRate = "PT30S"
	TargetSamplingRatePt1m  TargetSamplingRate = "PT1M"
	TargetSamplingRatePt5m  TargetSamplingRate = "PT5M"
	TargetSamplingRatePt10m TargetSamplingRate = "PT10M"
	TargetSamplingRatePt15m TargetSamplingRate = "PT15M"
	TargetSamplingRatePt30m TargetSamplingRate = "PT30M"
	TargetSamplingRatePt1h  TargetSamplingRate = "PT1H"
)

// Values returns all known values for TargetSamplingRate. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TargetSamplingRate) Values() []TargetSamplingRate {
	return []TargetSamplingRate{
		"PT1S",
		"PT5S",
		"PT10S",
		"PT15S",
		"PT30S",
		"PT1M",
		"PT5M",
		"PT10M",
		"PT15M",
		"PT30M",
		"PT1H",
	}
}
