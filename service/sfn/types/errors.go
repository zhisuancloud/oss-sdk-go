// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified activity does not exist.
type ActivityDoesNotExist struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ActivityDoesNotExist) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActivityDoesNotExist) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActivityDoesNotExist) ErrorCode() string             { return "ActivityDoesNotExist" }
func (e *ActivityDoesNotExist) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum number of activities has been reached. Existing activities must be
// deleted before a new activity can be created.
type ActivityLimitExceeded struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ActivityLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActivityLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActivityLimitExceeded) ErrorCode() string             { return "ActivityLimitExceeded" }
func (e *ActivityLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum number of workers concurrently polling for activity tasks has been
// reached.
type ActivityWorkerLimitExceeded struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ActivityWorkerLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActivityWorkerLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActivityWorkerLimitExceeded) ErrorCode() string             { return "ActivityWorkerLimitExceeded" }
func (e *ActivityWorkerLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The execution has the same name as another execution (but a different input).
// Executions with the same name and input are considered idempotent.
type ExecutionAlreadyExists struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ExecutionAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExecutionAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExecutionAlreadyExists) ErrorCode() string             { return "ExecutionAlreadyExists" }
func (e *ExecutionAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified execution does not exist.
type ExecutionDoesNotExist struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ExecutionDoesNotExist) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExecutionDoesNotExist) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExecutionDoesNotExist) ErrorCode() string             { return "ExecutionDoesNotExist" }
func (e *ExecutionDoesNotExist) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum number of running executions has been reached. Running executions
// must end or be stopped before a new execution can be started.
type ExecutionLimitExceeded struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ExecutionLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExecutionLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExecutionLimitExceeded) ErrorCode() string             { return "ExecutionLimitExceeded" }
func (e *ExecutionLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided Amazon Resource Name (ARN) is invalid.
type InvalidArn struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidArn) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArn) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArn) ErrorCode() string             { return "InvalidArn" }
func (e *InvalidArn) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided Amazon States Language definition is invalid.
type InvalidDefinition struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDefinition) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDefinition) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDefinition) ErrorCode() string             { return "InvalidDefinition" }
func (e *InvalidDefinition) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided JSON input data is invalid.
type InvalidExecutionInput struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidExecutionInput) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidExecutionInput) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidExecutionInput) ErrorCode() string             { return "InvalidExecutionInput" }
func (e *InvalidExecutionInput) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidLoggingConfiguration struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidLoggingConfiguration) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLoggingConfiguration) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLoggingConfiguration) ErrorCode() string             { return "InvalidLoggingConfiguration" }
func (e *InvalidLoggingConfiguration) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided name is invalid.
type InvalidName struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidName) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidName) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidName) ErrorCode() string             { return "InvalidName" }
func (e *InvalidName) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided JSON output data is invalid.
type InvalidOutput struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidOutput) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOutput) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOutput) ErrorCode() string             { return "InvalidOutput" }
func (e *InvalidOutput) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided token is invalid.
type InvalidToken struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidToken) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidToken) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidToken) ErrorCode() string             { return "InvalidToken" }
func (e *InvalidToken) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your tracingConfiguration key does not match, or enabled has not been set to
// true or false.
type InvalidTracingConfiguration struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidTracingConfiguration) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTracingConfiguration) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTracingConfiguration) ErrorCode() string             { return "InvalidTracingConfiguration" }
func (e *InvalidTracingConfiguration) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request is missing a required parameter. This error occurs if both definition
// and roleArn are not specified.
type MissingRequiredParameter struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MissingRequiredParameter) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingRequiredParameter) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingRequiredParameter) ErrorCode() string             { return "MissingRequiredParameter" }
func (e *MissingRequiredParameter) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Could not find the referenced resource. Only state machine and activity ARNs are
// supported.
type ResourceNotFound struct {
	Message *string

	ResourceName *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFound) ErrorCode() string             { return "ResourceNotFound" }
func (e *ResourceNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A state machine with the same name but a different definition or role ARN
// already exists.
type StateMachineAlreadyExists struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *StateMachineAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StateMachineAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StateMachineAlreadyExists) ErrorCode() string             { return "StateMachineAlreadyExists" }
func (e *StateMachineAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified state machine is being deleted.
type StateMachineDeleting struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *StateMachineDeleting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StateMachineDeleting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StateMachineDeleting) ErrorCode() string             { return "StateMachineDeleting" }
func (e *StateMachineDeleting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified state machine does not exist.
type StateMachineDoesNotExist struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *StateMachineDoesNotExist) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StateMachineDoesNotExist) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StateMachineDoesNotExist) ErrorCode() string             { return "StateMachineDoesNotExist" }
func (e *StateMachineDoesNotExist) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum number of state machines has been reached. Existing state machines
// must be deleted before a new state machine can be created.
type StateMachineLimitExceeded struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *StateMachineLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StateMachineLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StateMachineLimitExceeded) ErrorCode() string             { return "StateMachineLimitExceeded" }
func (e *StateMachineLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type StateMachineTypeNotSupported struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *StateMachineTypeNotSupported) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StateMachineTypeNotSupported) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StateMachineTypeNotSupported) ErrorCode() string             { return "StateMachineTypeNotSupported" }
func (e *StateMachineTypeNotSupported) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type TaskDoesNotExist struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TaskDoesNotExist) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TaskDoesNotExist) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TaskDoesNotExist) ErrorCode() string             { return "TaskDoesNotExist" }
func (e *TaskDoesNotExist) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type TaskTimedOut struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TaskTimedOut) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TaskTimedOut) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TaskTimedOut) ErrorCode() string             { return "TaskTimedOut" }
func (e *TaskTimedOut) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've exceeded the number of tags allowed for a resource. See the  Limits Topic
// (https://docs.aws.amazon.com/step-functions/latest/dg/limits.html) in the AWS
// Step Functions Developer Guide.
type TooManyTags struct {
	Message *string

	ResourceName *string

	noSmithyDocumentSerde
}

func (e *TooManyTags) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTags) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTags) ErrorCode() string             { return "TooManyTags" }
func (e *TooManyTags) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }