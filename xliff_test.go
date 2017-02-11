// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package xliff_test

import (
	"encoding/xml"
	"os"
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

func containsValidationError(errors []xliff.ValidationError, code xliff.ValidationErrorCode) bool {
	for _, err := range errors {
		if err.Code == code {
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

	if !containsValidationError(errors, xliff.UnsupportedVersion) {
		t.Error("Expected validation to fail with UnsupportedVersion")
	}

	if !containsValidationError(errors, xliff.MissingOriginalAttribute) {
		t.Error("Expected validation to fail with MissingOriginalAttribute")
	}

	if !containsValidationError(errors, xliff.MissingSourceLanguage) {
		t.Error("Expected validation to fail with MissingSourceLanguage")
	}

	if !containsValidationError(errors, xliff.MissingTargetLanguage) {
		t.Error("Expected validation to fail with MissingTargetLanguage")
	}

	if !containsValidationError(errors, xliff.UnsupportedDatatype) {
		t.Error("Expected validation to fail with UnsupportedDatatype")
	}

	if !containsValidationError(errors, xliff.InconsistentSourceLanguage) {
		t.Error("Expected validation to fail with InconsistentSourceLanguage")
	}

	if !containsValidationError(errors, xliff.InconsistentTargetLanguage) {
		t.Error("Expected validation to fail with InconsistentTargetLanguage")
	}

	if !containsValidationError(errors, xliff.MissingTransUnitID) {
		t.Error("Expected validation to fail with MissingTransUnitID")
	}

	if !containsValidationError(errors, xliff.MissingTransUnitSource) {
		t.Error("Expected validation to fail with MissingTransUnitSource")
	}

	if !containsValidationError(errors, xliff.MissingTransUnitTarget) {
		t.Error("Expected validation to fail with MissingTransUnitTarget")
	}
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
