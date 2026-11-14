package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"example.com/rfc-go-labs/lab"
)

func TestExample(t *testing.T) {
	input, err := os.ReadFile("../../examples/payload.hex")
	if err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	if err := run(nil, bytes.NewReader(input), &output); err != nil {
		t.Fatal(err)
	}
	var report lab.Report
	if err := json.Unmarshal(output.Bytes(), &report); err != nil {
		t.Fatal(err)
	}
	if report.Bytes != 11 || len(report.Commands) != 5 ||
		report.Commands[4].Name != "CLOSING" || report.Commands[4].Offset != 9 {
		t.Fatalf("unexpected fixture report: %+v", report)
	}
	output.Reset()
	if err := run([]string{"-html"}, bytes.NewReader(input), &output); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "<caption>5 decoded commands</caption>") {
		t.Fatal("expected the five-command report")
	}
}

func TestRejectsBadInputWithoutReport(t *testing.T) {
	for _, input := range []string{"00 12 fe", "01 12", strings.Repeat(" ", (64<<10)+1)} {
		var output bytes.Buffer
		if err := run(nil, strings.NewReader(input), &output); err == nil || output.Len() != 0 {
			t.Fatalf("invalid input of %d bytes produced %d bytes, error %v", len(input), output.Len(), err)
		}
	}
}
