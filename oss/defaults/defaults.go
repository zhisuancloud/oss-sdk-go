// Code generated by oss-sdk-go/internal/codegen/cmd/defaultsconfig. DO NOT EDIT.

package defaults

import (
	"fmt"
	"oss-sdk-go/oss"
	"time"
)

// GetModeConfiguration returns the default Configuration descriptor for the given mode.
//
// Supports the following modes: cross-region, in-region, mobile, standard
func GetModeConfiguration(mode oss.DefaultsMode) (Configuration, error) {
	var mv oss.DefaultsMode
	mv.SetFromString(string(mode))

	switch mv {
	case oss.DefaultsModeCrossRegion:
		settings := Configuration{
			ConnectTimeout:        oss.Duration(3100 * time.Millisecond),
			RetryMode:             oss.RetryMode("standard"),
			TLSNegotiationTimeout: oss.Duration(3100 * time.Millisecond),
		}
		return settings, nil
	case oss.DefaultsModeInRegion:
		settings := Configuration{
			ConnectTimeout:        oss.Duration(1100 * time.Millisecond),
			RetryMode:             oss.RetryMode("standard"),
			TLSNegotiationTimeout: oss.Duration(1100 * time.Millisecond),
		}
		return settings, nil
	case oss.DefaultsModeMobile:
		settings := Configuration{
			ConnectTimeout:        oss.Duration(30000 * time.Millisecond),
			RetryMode:             oss.RetryMode("standard"),
			TLSNegotiationTimeout: oss.Duration(30000 * time.Millisecond),
		}
		return settings, nil
	case oss.DefaultsModeStandard:
		settings := Configuration{
			ConnectTimeout:        oss.Duration(3100 * time.Millisecond),
			RetryMode:             oss.RetryMode("standard"),
			TLSNegotiationTimeout: oss.Duration(3100 * time.Millisecond),
		}
		return settings, nil
	default:
		return Configuration{}, fmt.Errorf("unsupported defaults mode: %v", mode)
	}
}