// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"oss-sdk-go/service/healthlake/types"
)

func ExampleInputDataConfig_outputUsage() {
	var union types.InputDataConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InputDataConfigMemberS3Uri:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleOutputDataConfig_outputUsage() {
	var union types.OutputDataConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OutputDataConfigMemberS3Configuration:
		_ = v.Value // Value is types.S3Configuration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3Configuration
