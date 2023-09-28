// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ConfirmationState string

// Enum values for ConfirmationState
const (
	ConfirmationStateConfirmed ConfirmationState = "Confirmed"
	ConfirmationStateDenied    ConfirmationState = "Denied"
	ConfirmationStateNone      ConfirmationState = "None"
)

// Values returns all known values for ConfirmationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfirmationState) Values() []ConfirmationState {
	return []ConfirmationState{
		"Confirmed",
		"Denied",
		"None",
	}
}

type ConversationMode string

// Enum values for ConversationMode
const (
	ConversationModeAudio ConversationMode = "AUDIO"
	ConversationModeText  ConversationMode = "TEXT"
)

// Values returns all known values for ConversationMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConversationMode) Values() []ConversationMode {
	return []ConversationMode{
		"AUDIO",
		"TEXT",
	}
}

type DialogActionType string

// Enum values for DialogActionType
const (
	DialogActionTypeClose         DialogActionType = "Close"
	DialogActionTypeConfirmIntent DialogActionType = "ConfirmIntent"
	DialogActionTypeDelegate      DialogActionType = "Delegate"
	DialogActionTypeElicitIntent  DialogActionType = "ElicitIntent"
	DialogActionTypeElicitSlot    DialogActionType = "ElicitSlot"
	DialogActionTypeNone          DialogActionType = "None"
)

// Values returns all known values for DialogActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DialogActionType) Values() []DialogActionType {
	return []DialogActionType{
		"Close",
		"ConfirmIntent",
		"Delegate",
		"ElicitIntent",
		"ElicitSlot",
		"None",
	}
}

type InputMode string

// Enum values for InputMode
const (
	InputModeText   InputMode = "Text"
	InputModeSpeech InputMode = "Speech"
	InputModeDtmf   InputMode = "DTMF"
)

// Values returns all known values for InputMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (InputMode) Values() []InputMode {
	return []InputMode{
		"Text",
		"Speech",
		"DTMF",
	}
}

type IntentState string

// Enum values for IntentState
const (
	IntentStateFailed                IntentState = "Failed"
	IntentStateFulfilled             IntentState = "Fulfilled"
	IntentStateInProgress            IntentState = "InProgress"
	IntentStateReadyForFulfillment   IntentState = "ReadyForFulfillment"
	IntentStateWaiting               IntentState = "Waiting"
	IntentStateFulfillmentInProgress IntentState = "FulfillmentInProgress"
)

// Values returns all known values for IntentState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IntentState) Values() []IntentState {
	return []IntentState{
		"Failed",
		"Fulfilled",
		"InProgress",
		"ReadyForFulfillment",
		"Waiting",
		"FulfillmentInProgress",
	}
}

type MessageContentType string

// Enum values for MessageContentType
const (
	MessageContentTypeCustomPayload     MessageContentType = "CustomPayload"
	MessageContentTypeImageResponseCard MessageContentType = "ImageResponseCard"
	MessageContentTypePlainText         MessageContentType = "PlainText"
	MessageContentTypeSsml              MessageContentType = "SSML"
)

// Values returns all known values for MessageContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MessageContentType) Values() []MessageContentType {
	return []MessageContentType{
		"CustomPayload",
		"ImageResponseCard",
		"PlainText",
		"SSML",
	}
}

type PlaybackInterruptionReason string

// Enum values for PlaybackInterruptionReason
const (
	PlaybackInterruptionReasonDtmfStartDetected  PlaybackInterruptionReason = "DTMF_START_DETECTED"
	PlaybackInterruptionReasonTextDetected       PlaybackInterruptionReason = "TEXT_DETECTED"
	PlaybackInterruptionReasonVoiceStartDetected PlaybackInterruptionReason = "VOICE_START_DETECTED"
)

// Values returns all known values for PlaybackInterruptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PlaybackInterruptionReason) Values() []PlaybackInterruptionReason {
	return []PlaybackInterruptionReason{
		"DTMF_START_DETECTED",
		"TEXT_DETECTED",
		"VOICE_START_DETECTED",
	}
}

type SentimentType string

// Enum values for SentimentType
const (
	SentimentTypeMixed    SentimentType = "MIXED"
	SentimentTypeNegative SentimentType = "NEGATIVE"
	SentimentTypeNeutral  SentimentType = "NEUTRAL"
	SentimentTypePositive SentimentType = "POSITIVE"
)

// Values returns all known values for SentimentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SentimentType) Values() []SentimentType {
	return []SentimentType{
		"MIXED",
		"NEGATIVE",
		"NEUTRAL",
		"POSITIVE",
	}
}

type Shape string

// Enum values for Shape
const (
	ShapeScalar    Shape = "Scalar"
	ShapeList      Shape = "List"
	ShapeComposite Shape = "Composite"
)

// Values returns all known values for Shape. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Shape) Values() []Shape {
	return []Shape{
		"Scalar",
		"List",
		"Composite",
	}
}

type StyleType string

// Enum values for StyleType
const (
	StyleTypeDefault       StyleType = "Default"
	StyleTypeSpellByLetter StyleType = "SpellByLetter"
	StyleTypeSpellByWord   StyleType = "SpellByWord"
)

// Values returns all known values for StyleType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (StyleType) Values() []StyleType {
	return []StyleType{
		"Default",
		"SpellByLetter",
		"SpellByWord",
	}
}
