// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"oss-sdk-go/service/migrationhubstrategy/types"
)

func ExampleDatabaseMigrationPreference_outputUsage() {
	var union types.DatabaseMigrationPreference
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DatabaseMigrationPreferenceMemberHeterogeneous:
		_ = v.Value // Value is types.Heterogeneous

	case *types.DatabaseMigrationPreferenceMemberHomogeneous:
		_ = v.Value // Value is types.Homogeneous

	case *types.DatabaseMigrationPreferenceMemberNoPreference:
		_ = v.Value // Value is types.NoDatabaseMigrationPreference

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.Heterogeneous
var _ *types.NoDatabaseMigrationPreference
var _ *types.Homogeneous

func ExampleManagementPreference_outputUsage() {
	var union types.ManagementPreference
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ManagementPreferenceMemberAwsManagedResources:
		_ = v.Value // Value is types.AwsManagedResources

	case *types.ManagementPreferenceMemberNoPreference:
		_ = v.Value // Value is types.NoManagementPreference

	case *types.ManagementPreferenceMemberSelfManageResources:
		_ = v.Value // Value is types.SelfManageResources

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.NoManagementPreference
var _ *types.AwsManagedResources
var _ *types.SelfManageResources
