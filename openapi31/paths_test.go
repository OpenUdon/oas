// Copyright (c) 2025 Greetingland LLC
// Created with the help of Claude Code
// MIT License - see LICENSE file for details

package openapi31

import (
	"encoding/json"
	"testing"
)

func TestPathItemRefMarshalOnly(t *testing.T) {
	item := PathItem{
		Ref:        "#/components/pathItems/Example",
		Summary:    "ignored",
		Extensions: map[string]any{"x-extra": true},
		Get:        &Operation{},
	}

	data, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(raw) != 1 {
		t.Fatalf("expected only $ref, got keys: %v", raw)
	}
	if raw["$ref"] != item.Ref {
		t.Fatalf("expected $ref %q, got %v", item.Ref, raw["$ref"])
	}
}
