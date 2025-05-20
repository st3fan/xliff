# XLIFF - Go Package to process XLIFF Files

[![Build Status](https://travis-ci.org/st3fan/xliff.svg?branch=master)](https://travis-ci.org/st3fan/xliff) [![Go Report Card](https://goreportcard.com/badge/github.com/st3fan/xliff)](https://goreportcard.com/report/github.com/st3fan/xliff) [![codecov](https://codecov.io/gh/st3fan/xliff/branch/master/graph/badge.svg)](https://codecov.io/gh/st3fan/xliff)


*Stefan Arentz, February 2017*

Work in progress. You can see it in action in the `validate-xliff` tool at [st3fan/validate-xliff](https://github.com/st3fan/validate-xliff)

This is a tiny package to parse `.xliff` files.

It is currently used in the following two projects:

* In [Firefox Focus](https://github.com/mozilla-mobile/focus) to import strings from the Mozilla L10N Repository into the Focus Xcode project.
* In [XLIFF Tool](https://github.com/st3fan/xlifftool), a generic XLIFF tool that will hopefully replace a number of specialized scripts that we use for [https://github.com/mozilla-mobile/firefox-ios](Firefox for iOS) and [https://github.com/mozilla-mobile/focus](Firefox Focus for iOS).

## Examples

### Loading an XLIFF Document

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/st3fan/xliff"
)

func main() {
    // Load an XLIFF document from a file
    doc, err := xliff.FromFile("translations.xliff")
    if err != nil {
        log.Fatalf("Failed to load XLIFF document: %v", err)
    }
    
    fmt.Printf("Loaded XLIFF document with %d files\n", len(doc.Files))
}
```

### Validating an XLIFF Document

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/st3fan/xliff"
)

func main() {
    // Load and validate an XLIFF document
    doc, err := xliff.FromFile("translations.xliff")
    if err != nil {
        log.Fatalf("Failed to load XLIFF document: %v", err)
    }
    
    // Check for structural issues
    errors := doc.Validate()
    if len(errors) > 0 {
        fmt.Println("Document validation failed:")
        for _, err := range errors {
            fmt.Printf("  - %s\n", err.Error())
        }
        return
    }
    
    fmt.Println("Document is valid")
}
```

### Checking for Completeness

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/st3fan/xliff"
)

func main() {
    // Load an XLIFF document
    doc, err := xliff.FromFile("translations.xliff")
    if err != nil {
        log.Fatalf("Failed to load XLIFF document: %v", err)
    }
    
    // Check if all translations are complete
    if doc.IsComplete() {
        fmt.Println("All translations are complete")
    } else {
        fmt.Println("Some translations are incomplete")
    }
}
```

### Accessing Specific Files

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/st3fan/xliff"
)

func main() {
    // Load an XLIFF document
    doc, err := xliff.FromFile("translations.xliff")
    if err != nil {
        log.Fatalf("Failed to load XLIFF document: %v", err)
    }
    
    // Access a specific file by name
    file, found := doc.File("Localizable.strings")
    if !found {
        fmt.Println("File 'Localizable.strings' not found")
        return
    }
    
    // Print information about the file
    fmt.Printf("Found file '%s' with %d translation units\n", 
        file.Original, len(file.Body.TransUnits))
    
    // Print all translation units in the file
    for _, unit := range file.Body.TransUnits {
        fmt.Printf("ID: %s\n", unit.ID)
        fmt.Printf("  Source: %s\n", unit.Source)
        fmt.Printf("  Target: %s\n", unit.Target)
        if unit.Note != "" {
            fmt.Printf("  Note: %s\n", unit.Note)
        }
    }
}
```

> Please [file a bug](https://github.com/st3fan/xliff/issues/new) if you have some specific wish or requirement.
