// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
// DO NOT EDIT: this file is automatically generated by docgen
package types

import (
	"github.com/projectdiscovery/yamldoc-go/encoder"
)

var (
	FTWOverridesDoc     encoder.Doc
	FTWOverridesMetaDoc encoder.Doc
	TestOverrideDoc     encoder.Doc
	OutputDoc           encoder.Doc
	LogDoc              encoder.Doc
)

func init() {
	FTWOverridesDoc.Type = "FTWOverrides"
	FTWOverridesDoc.Comments[encoder.LineComment] = " FTWOverrides describes platform specific overrides for tests"
	FTWOverridesDoc.Description = "FTWOverrides describes platform specific overrides for tests"
	FTWOverridesDoc.Fields = make([]encoder.Doc, 3)
	FTWOverridesDoc.Fields[0].Name = "version"
	FTWOverridesDoc.Fields[0].Type = "string"
	FTWOverridesDoc.Fields[0].Note = ""
	FTWOverridesDoc.Fields[0].Description = "The version field designates the version of the schema that validates this file"
	FTWOverridesDoc.Fields[0].Comments[encoder.LineComment] = "The version field designates the version of the schema that validates this file"

	FTWOverridesDoc.Fields[0].AddExample("", "v0.1.0")
	FTWOverridesDoc.Fields[1].Name = "meta"
	FTWOverridesDoc.Fields[1].Type = "FTWOverridesMeta"
	FTWOverridesDoc.Fields[1].Note = ""
	FTWOverridesDoc.Fields[1].Description = "Meta describes the metadata information"
	FTWOverridesDoc.Fields[1].Comments[encoder.LineComment] = "Meta describes the metadata information"

	FTWOverridesDoc.Fields[1].AddExample("", metaExample)
	FTWOverridesDoc.Fields[2].Name = "test_overrides"
	FTWOverridesDoc.Fields[2].Type = "[]TestOverride"
	FTWOverridesDoc.Fields[2].Note = ""
	FTWOverridesDoc.Fields[2].Description = "List of test override specifications"
	FTWOverridesDoc.Fields[2].Comments[encoder.LineComment] = "List of test override specifications"

	FTWOverridesDoc.Fields[2].AddExample("", testOverridesExample)

	FTWOverridesMetaDoc.Type = "FTWOverridesMeta"
	FTWOverridesMetaDoc.Comments[encoder.LineComment] = ""
	FTWOverridesMetaDoc.Description = ""

	FTWOverridesMetaDoc.AddExample("", metaExample)
	FTWOverridesMetaDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "FTWOverrides",
			FieldName: "meta",
		},
	}
	FTWOverridesMetaDoc.Fields = make([]encoder.Doc, 3)
	FTWOverridesMetaDoc.Fields[0].Name = "engine"
	FTWOverridesMetaDoc.Fields[0].Type = "string"
	FTWOverridesMetaDoc.Fields[0].Note = ""
	FTWOverridesMetaDoc.Fields[0].Description = "The name of the WAF engine the tests are expected to run against"
	FTWOverridesMetaDoc.Fields[0].Comments[encoder.LineComment] = "The name of the WAF engine the tests are expected to run against"

	FTWOverridesMetaDoc.Fields[0].AddExample("", "coraza")
	FTWOverridesMetaDoc.Fields[1].Name = "platform"
	FTWOverridesMetaDoc.Fields[1].Type = "string"
	FTWOverridesMetaDoc.Fields[1].Note = ""
	FTWOverridesMetaDoc.Fields[1].Description = "The name of the platform (e.g., web server) the tests are expected to run against"
	FTWOverridesMetaDoc.Fields[1].Comments[encoder.LineComment] = "The name of the platform (e.g., web server) the tests are expected to run against"

	FTWOverridesMetaDoc.Fields[1].AddExample("", "nginx")
	FTWOverridesMetaDoc.Fields[2].Name = "annotations"
	FTWOverridesMetaDoc.Fields[2].Type = "map[string]string"
	FTWOverridesMetaDoc.Fields[2].Note = ""
	FTWOverridesMetaDoc.Fields[2].Description = "Custom annotations; can be used to add additional meta information"
	FTWOverridesMetaDoc.Fields[2].Comments[encoder.LineComment] = "Custom annotations; can be used to add additional meta information"

	FTWOverridesMetaDoc.Fields[2].AddExample("", annotationsExample)

	TestOverrideDoc.Type = "TestOverride"
	TestOverrideDoc.Comments[encoder.LineComment] = ""
	TestOverrideDoc.Description = ""

	TestOverrideDoc.AddExample("", testOverridesExample)
	TestOverrideDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "FTWOverrides",
			FieldName: "test_overrides",
		},
	}
	TestOverrideDoc.Fields = make([]encoder.Doc, 5)
	TestOverrideDoc.Fields[0].Name = "rule_id"
	TestOverrideDoc.Fields[0].Type = "int"
	TestOverrideDoc.Fields[0].Note = ""
	TestOverrideDoc.Fields[0].Description = "ID of the rule this test targets."
	TestOverrideDoc.Fields[0].Comments[encoder.LineComment] = "ID of the rule this test targets."

	TestOverrideDoc.Fields[0].AddExample("", 920100)
	TestOverrideDoc.Fields[1].Name = "test_ids"
	TestOverrideDoc.Fields[1].Type = "[]int"
	TestOverrideDoc.Fields[1].Note = ""
	TestOverrideDoc.Fields[1].Description = "description: |\n     IDs of the tests for rule_id that overrides should be applied to.\n     If this field is not set, the overrides will be applied to all tests of rule_id.\n examples:\n     - value: [5, 6]"
	TestOverrideDoc.Fields[1].Comments[encoder.LineComment] = " description: |"

	TestOverrideDoc.Fields[2].Name = "reason"
	TestOverrideDoc.Fields[2].Type = "string"
	TestOverrideDoc.Fields[2].Note = ""
	TestOverrideDoc.Fields[2].Description = "Describes why this override is necessary."
	TestOverrideDoc.Fields[2].Comments[encoder.LineComment] = "Describes why this override is necessary."

	TestOverrideDoc.Fields[2].AddExample("", reasonExample)
	TestOverrideDoc.Fields[3].Name = "expect_failure"
	TestOverrideDoc.Fields[3].Type = "bool"
	TestOverrideDoc.Fields[3].Note = ""
	TestOverrideDoc.Fields[3].Description = "Whether this test is expected to fail for this particular configuration.\nDefault: false"
	TestOverrideDoc.Fields[3].Comments[encoder.LineComment] = "Whether this test is expected to fail for this particular configuration."

	TestOverrideDoc.Fields[3].AddExample("", true)
	TestOverrideDoc.Fields[4].Name = "output"
	TestOverrideDoc.Fields[4].Type = "Output"
	TestOverrideDoc.Fields[4].Note = ""
	TestOverrideDoc.Fields[4].Description = "Specifies overrides on the test output"
	TestOverrideDoc.Fields[4].Comments[encoder.LineComment] = "Specifies overrides on the test output"

	TestOverrideDoc.Fields[4].AddExample("", 400)

	OutputDoc.Type = "Output"
	OutputDoc.Comments[encoder.LineComment] = ""
	OutputDoc.Description = ""

	OutputDoc.AddExample("", 400)
	OutputDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "TestOverride",
			FieldName: "output",
		},
	}
	OutputDoc.Fields = make([]encoder.Doc, 6)
	OutputDoc.Fields[0].Name = "status"
	OutputDoc.Fields[0].Type = "int"
	OutputDoc.Fields[0].Note = ""
	OutputDoc.Fields[0].Description = "Status describes the HTTP status code expected in the response."
	OutputDoc.Fields[0].Comments[encoder.LineComment] = "Status describes the HTTP status code expected in the response."

	OutputDoc.Fields[0].AddExample("Status", 200)
	OutputDoc.Fields[1].Name = "response_contains"
	OutputDoc.Fields[1].Type = "string"
	OutputDoc.Fields[1].Note = ""
	OutputDoc.Fields[1].Description = "ResponseContains describes the text that should be contained in the HTTP response."
	OutputDoc.Fields[1].Comments[encoder.LineComment] = "ResponseContains describes the text that should be contained in the HTTP response."

	OutputDoc.Fields[1].AddExample("ResponseContains", "Hello, World")
	OutputDoc.Fields[2].Name = "log_contains"
	OutputDoc.Fields[2].Type = "string"
	OutputDoc.Fields[2].Note = ""
	OutputDoc.Fields[2].Description = "LogContains describes the text that should be contained in the WAF logs."
	OutputDoc.Fields[2].Comments[encoder.LineComment] = "LogContains describes the text that should be contained in the WAF logs."

	OutputDoc.Fields[2].AddExample("LogContains", "id 920100")
	OutputDoc.Fields[3].Name = "no_log_contains"
	OutputDoc.Fields[3].Type = "string"
	OutputDoc.Fields[3].Note = ""
	OutputDoc.Fields[3].Description = "NoLogContains describes the text that should not be contained in the WAF logs."
	OutputDoc.Fields[3].Comments[encoder.LineComment] = "NoLogContains describes the text that should not be contained in the WAF logs."

	OutputDoc.Fields[3].AddExample("NoLogContains", "id 920100")
	OutputDoc.Fields[4].Name = "log"
	OutputDoc.Fields[4].Type = "Log"
	OutputDoc.Fields[4].Note = ""
	OutputDoc.Fields[4].Description = "Log is used to configure expectations about the log contents."
	OutputDoc.Fields[4].Comments[encoder.LineComment] = "Log is used to configure expectations about the log contents."

	OutputDoc.Fields[4].AddExample("", exampleLog)
	OutputDoc.Fields[5].Name = "expect_error"
	OutputDoc.Fields[5].Type = "bool"
	OutputDoc.Fields[5].Note = ""
	OutputDoc.Fields[5].Description = "When `ExpectError` is true, we don't expect an answer from the WAF, just an error."
	OutputDoc.Fields[5].Comments[encoder.LineComment] = "When `ExpectError` is true, we don't expect an answer from the WAF, just an error."

	OutputDoc.Fields[5].AddExample("ExpectError", false)

	LogDoc.Type = "Log"
	LogDoc.Comments[encoder.LineComment] = ""
	LogDoc.Description = ""

	LogDoc.AddExample("", exampleLog)
	LogDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Output",
			FieldName: "log",
		},
	}
	LogDoc.Fields = make([]encoder.Doc, 4)
	LogDoc.Fields[0].Name = "expect_id"
	LogDoc.Fields[0].Type = "int"
	LogDoc.Fields[0].Note = ""
	LogDoc.Fields[0].Description = "description: |\n   Expect the given ID to be contained in the log output.\n examples:\n   - exampleLog.ExpectId"
	LogDoc.Fields[0].Comments[encoder.LineComment] = " description: |"
	LogDoc.Fields[1].Name = "expect_id"
	LogDoc.Fields[1].Type = "int"
	LogDoc.Fields[1].Note = ""
	LogDoc.Fields[1].Description = "description: |\n   Expect the given ID _not_ to be contained in the log output.\n examples:\n   - exampleLog.NoExpectId"
	LogDoc.Fields[1].Comments[encoder.LineComment] = " description: |"
	LogDoc.Fields[2].Name = "match_regex"
	LogDoc.Fields[2].Type = "string"
	LogDoc.Fields[2].Note = ""
	LogDoc.Fields[2].Description = "Expect the regular expression to match log content for the current test."
	LogDoc.Fields[2].Comments[encoder.LineComment] = "Expect the regular expression to match log content for the current test."

	LogDoc.Fields[2].AddExample("", exampleLog.MatchRegex)
	LogDoc.Fields[3].Name = "no_match_regex"
	LogDoc.Fields[3].Type = "string"
	LogDoc.Fields[3].Note = ""
	LogDoc.Fields[3].Description = "Expect the regular expression to _not_ match log content for the current test."
	LogDoc.Fields[3].Comments[encoder.LineComment] = "Expect the regular expression to _not_ match log content for the current test."

	LogDoc.Fields[3].AddExample("", exampleLog.NoMatchRegex)
}

// GetFTWOverridesDoc returns documentation for the file ./overrides_doc.go.
func GetFTWOverridesDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "FTWOverrides",
		Description: "",
		Structs: []*encoder.Doc{
			&FTWOverridesDoc,
			&FTWOverridesMetaDoc,
			&TestOverrideDoc,
			&OutputDoc,
			&LogDoc,
		},
	}
}