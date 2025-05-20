// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package xliff_test

import (
	"encoding/xml"
	"os"
	"strings"
	"testing"

	"github.com/st3fan/xliff"
)

func Test_Parse(t *testing.T) {
	if _, err := xliff.FromFile("testdata/focus-ios-ar.xliff"); err != nil {
		t.Error("Could not parse testdata/focus-ios-ar.xliff:", err)
	}
	if _, err := xliff.FromFile("testdata/focus-ios-it.xliff"); err != nil {
		t.Error("Could not parse testdata/focus-ios-it.xliff:", err)
	}
}

func Test_ParseNonExistentFile(t *testing.T) {
	if _, err := xliff.FromFile("testdata/doesnotexist.xliff"); !os.IsNotExist(err) {
		t.Error("Unexpected error when opening testdata/doesnotexist.xliff:", err)
	}
}

func Test_ParseBadXMLFile(t *testing.T) {
	_, err := xliff.FromFile("testdata/badxml.xliff")
	if _, ok := err.(*xml.SyntaxError); !ok {
		t.Error("Unexpected error when opening testdata/badxml.xliff:", err)
	}
}

func Test_ValidateGood(t *testing.T) {
	doc, err := xliff.FromFile("testdata/good.xliff")
	if err != nil {
		t.Error("Could not parse testdata/good.xliff:", err)
	}

	if errors := doc.Validate(); errors != nil {
		t.Error("Unexpected error from Validate()")
	}
}

func containsValidationError(t *testing.T, errors []xliff.ValidationError, code xliff.ValidationErrorCode) bool {
	for _, err := range errors {
		if err.Code == code {
			if strings.HasPrefix(err.Error(), "Unknown: ") {
				t.Error("Error has no good message: ", err)
			}
			return true
		}
	}
	return false
}

func Test_ValidateErrors(t *testing.T) {
	doc, err := xliff.FromFile("testdata/errors.xliff")
	if err != nil {
		t.Error("Could not parse testdata/errors.xliff:", err)
	}

	errors := doc.Validate()
	if len(errors) == 0 {
		t.Error("Expected error from Validate()")
	}

	if !containsValidationError(t, errors, xliff.UnsupportedVersion) {
		t.Error("Expected validation to fail with UnsupportedVersion")
	}

	if !containsValidationError(t, errors, xliff.MissingOriginalAttribute) {
		t.Error("Expected validation to fail with MissingOriginalAttribute")
	}

	if !containsValidationError(t, errors, xliff.MissingSourceLanguage) {
		t.Error("Expected validation to fail with MissingSourceLanguage")
	}

	if !containsValidationError(t, errors, xliff.MissingTargetLanguage) {
		t.Error("Expected validation to fail with MissingTargetLanguage")
	}

	if !containsValidationError(t, errors, xliff.UnsupportedDatatype) {
		t.Error("Expected validation to fail with UnsupportedDatatype")
	}

	if !containsValidationError(t, errors, xliff.InconsistentSourceLanguage) {
		t.Error("Expected validation to fail with InconsistentSourceLanguage")
	}

	if !containsValidationError(t, errors, xliff.InconsistentTargetLanguage) {
		t.Error("Expected validation to fail with InconsistentTargetLanguage")
	}

	if !containsValidationError(t, errors, xliff.MissingTransUnitID) {
		t.Error("Expected validation to fail with MissingTransUnitID")
	}

	if !containsValidationError(t, errors, xliff.MissingTransUnitSource) {
		t.Error("Expected validation to fail with MissingTransUnitSource")
	}

	// Removed test for MissingTransUnitTarget as per issue #5
	// We no longer consider missing target as a validation error
}

func Test_IsComplete(t *testing.T) {
	doc, err := xliff.FromFile("testdata/complete.xliff")
	if err != nil {
		t.Error("Could not parse testdata/complete.xliff:", err)
	}

	if !doc.IsComplete() {
		t.Error("Unexpected result from doc.IsComplete(). Got false, expected true")
	}
}

func Test_IsInComplete(t *testing.T) {
	doc, err := xliff.FromFile("testdata/incomplete.xliff")
	if err != nil {
		t.Error("Could not parse testdata/incomplete.xliff:", err)
	}

	if doc.IsComplete() {
		t.Error("Unexpected result from doc.IsComplete(). Got true, expected false")
	}
}

func Test_File(t *testing.T) {
	doc, err := xliff.FromFile("testdata/complete.xliff")
	if err != nil {
		t.Error("Could not parse testdata/complete.xliff:", err)
	}

	if _, found := doc.File("One.strings"); found != true {
		t.Error("Unexpected result from doc.File(One.strings)")
	}
	if _, found := doc.File("Unknown.strings"); found != false {
		t.Error("Unexpected result from doc.File(Unknown.strings)")
	}
}

func Test_ValidateIncomplete(t *testing.T) {
	doc, err := xliff.FromFile("testdata/incomplete.xliff")
	if err != nil {
		t.Error("Could not parse testdata/incomplete.xliff:", err)
	}

	// Validation should pass on incomplete but otherwise correct files
	errors := doc.Validate()
	if len(errors) > 0 {
		t.Error("Validation should not fail on incomplete files")
		for _, err := range errors {
			t.Error("  ", err)
		}
	}

	// IsComplete should still correctly identify it as incomplete
	if doc.IsComplete() {
		t.Error("Expected incomplete file to be identified as incomplete")
	}
}
