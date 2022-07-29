package cmd

import (
	"reflect"
	"testing"

	"github.com/razonyang/opencc"
	"github.com/urfave/cli/v2"
)

func TestGetConverter(t *testing.T) {
	conversions := map[string]bool{
		"s2t":     false,
		"t2s":     false,
		"s2tw":    false,
		"tw2s":    false,
		"s2hk":    false,
		"hk2s":    false,
		"s2twp":   false,
		"tw2sp":   false,
		"t2tw":    false,
		"hk2t":    false,
		"t2hk":    false,
		"t2jp":    false,
		"jp2t":    false,
		"tw2t":    false,
		"invalid": true,
	}
	for conversion, hasErr := range conversions {
		app := cli.NewApp()
		app.Metadata = map[string]interface{}{}
		converter, err := getConverter(app, conversion)
		if hasErr && err == nil {
			t.Errorf("expected error: %s", conversion)
		}
		expected, _ := opencc.New(conversion)
		if !hasErr && !reflect.DeepEqual(converter, expected) {
			t.Errorf("expected %v, got %v", expected, converter)
		}
	}
}
