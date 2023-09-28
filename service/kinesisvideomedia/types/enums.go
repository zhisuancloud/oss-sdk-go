// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type StartSelectorType string

// Enum values for StartSelectorType
const (
	StartSelectorTypeFragmentNumber    StartSelectorType = "FRAGMENT_NUMBER"
	StartSelectorTypeServerTimestamp   StartSelectorType = "SERVER_TIMESTAMP"
	StartSelectorTypeProducerTimestamp StartSelectorType = "PRODUCER_TIMESTAMP"
	StartSelectorTypeNow               StartSelectorType = "NOW"
	StartSelectorTypeEarliest          StartSelectorType = "EARLIEST"
	StartSelectorTypeContinuationToken StartSelectorType = "CONTINUATION_TOKEN"
)

// Values returns all known values for StartSelectorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StartSelectorType) Values() []StartSelectorType {
	return []StartSelectorType{
		"FRAGMENT_NUMBER",
		"SERVER_TIMESTAMP",
		"PRODUCER_TIMESTAMP",
		"NOW",
		"EARLIEST",
		"CONTINUATION_TOKEN",
	}
}
