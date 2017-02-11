// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package xliff_test

import (
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

func Test_ValidateGood(t *testing.T) {
	doc, err := xliff.FromFile("testdata/good.xliff")
	if err != nil {
		t.Error("Could not parse testdata/good.xliff:", err)
	}

	if errors := doc.Validate(); errors != nil {
		t.Error("Unexpected error from Validate()")
	}
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
}

func Test_IsComplete(t *testing.T) {
	doc, err := xliff.FromFile("testdata/focus-ios-ar.xliff")
	if err != nil {
		t.Error("Could not parse testdata/focus-ios-ar.xliff:", err)
	}

	if doc.IsComplete() {
		t.Error("Unexpected result from doc.IsComplete(). Got true, expected false")
	}
}
