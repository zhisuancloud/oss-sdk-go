//go:build integration
// +build integration

package rekognition

import (
	"context"
	"testing"
	"time"

	"oss-sdk-go/service/internal/integrationtest"
	"oss-sdk-go/service/rekognition"
)

func TestInteg_00_ListCollections(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := rekognition.NewFromConfig(cfg)
	params := &rekognition.ListCollectionsInput{}
	_, err = client.ListCollections(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
