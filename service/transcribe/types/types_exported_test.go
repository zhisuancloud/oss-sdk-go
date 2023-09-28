// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"oss-sdk-go/service/transcribe/types"
)

func ExampleRule_outputUsage() {
	var union types.Rule
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RuleMemberInterruptionFilter:
		_ = v.Value // Value is types.InterruptionFilter

	case *types.RuleMemberNonTalkTimeFilter:
		_ = v.Value // Value is types.NonTalkTimeFilter

	case *types.RuleMemberSentimentFilter:
		_ = v.Value // Value is types.SentimentFilter

	case *types.RuleMemberTranscriptFilter:
		_ = v.Value // Value is types.TranscriptFilter

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.NonTalkTimeFilter
var _ *types.SentimentFilter
var _ *types.TranscriptFilter
var _ *types.InterruptionFilter
