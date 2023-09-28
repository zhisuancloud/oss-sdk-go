//go:build integration
// +build integration

package dynamodb

import (
	"context"
	"errors"
	"testing"
	"time"

	"oss-sdk-go/oss"
	"oss-sdk-go/service/dynamodb"
	"github.com/aws/smithy-go"

	"oss-sdk-go/service/internal/integrationtest"
)

func TestInteg_00_ListTables(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	params := &dynamodb.ListTablesInput{
		Limit: oss.Int32(1),
	}
	_, err = client.ListTables(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_DescribeTable(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	params := &dynamodb.DescribeTableInput{
		TableName: oss.String("fake-table"),
	}
	_, err = client.DescribeTable(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}

	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expect error to be API error, was not, %v", err)
	}
	if len(apiErr.ErrorCode()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if len(apiErr.ErrorMessage()) == 0 {
		t.Errorf("expect non-empty error message")
	}
}
