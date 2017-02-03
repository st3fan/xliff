package xliff

import (
	"encoding/xml"
	"io/ioutil"
)

type Tool struct {
	ToolID      string `xml:"tool-id,attr"`
	ToolName    string `xml:"tool-name,attr"`
	ToolVersion string `xml:"tool-version,attr"`
	BuildNum    string `xml:"build-num,attr"`
}

type Header struct {
	Tool Tool `xml:"tool"`
}

type TransUnit struct {
	ID     string `xml:"id,attr"`
	Source string `xml:"source"`
	Target string `xml:"target"`
	Note   string `xml:"note"`
}

type Body struct {
	TransUnits []TransUnit `xml:"trans-unit"`
}

type File struct {
	Original       string `xml:"original,attr"`
	SourceLanguage string `xml:"source-language,attr"`
	Datatype       string `xml:"datatype,attr"`
	TargetLanguage string `xml:"target-language,attr"`
	Header         Header `xml:"header"`
	Body           Body   `xml:"body"`
}

type Document struct {
	Version string `xml:"version,attr"`
	Files   []File `xml:"file"`
}

func FromFile(path string) (Document, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Document{}, err
	}

	var document Document
	if err := xml.Unmarshal(data, &document); err != nil {
		return Document{}, err
	}

	return document, nil
}

func (d Document) IsValid() bool {
	// Make sure the document is a version we understand
	if d.Version != "1.2" {
		return false
	}
	// Make sure all files have the attributes we need
	for _, file := range d.Files {
		if file.Original == "" {
			return false
		}
		if file.SourceLanguage == "" {
			return false
		}
		if file.TargetLanguage == "" {
			return false
		}
		if file.Datatype != "plaintext" {
			return false
		}
	}
	// Make sure all files are consistent with source and target language
	sourceLanguage, targetLanguage := d.Files[0].SourceLanguage, d.Files[0].TargetLanguage
	for _, file := range d.Files {
		if file.SourceLanguage != sourceLanguage {
			return false
		}
		if file.TargetLanguage != targetLanguage {
			return false
		}
	}
	// Make sure all trans units have the attributes and children we expect
	for _, file := range d.Files {
		for _, transUnit := range file.Body.TransUnits {
			if transUnit.ID == "" {
				return false
			}
			// if transUnit.Source == "" {
			// 	return false
			// }
			// if transUnit.Target == "" {
			// 	return false
			// }
		}
	}
	return true
}

func (d Document) IsComplete() bool {
	for _, file := range d.Files {
		for _, transUnit := range file.Body.TransUnits {
			if transUnit.Source == "" || transUnit.Target == "" {
				return false
			}
		}
	}
	return true
}
