# ApprovalTests Go Architecture

## Overview

This diagram shows the core architecture of the ApprovalTests Go library, focusing on the main interfaces, core methods, and key components that make up the system.

## Architecture Diagram

```mermaid
classDiagram
    class Failable {
        <<interface>>
        +Error(args)
        +Name() string
        +Helper()
    }

    class ApprovalNamer {
        <<interface>>
        +GetName() string
        +GetReceivedFile(extWithDot string) string
        +GetApprovalFile(extWithDot string) string
    }

    class Reporter {
        <<interface>>
        +Report(approved string, received string) bool
    }

    class verifyOptions {
        -fields map
        +ForFile() fileOptions
        +WithScrubber(scrub scrubber) verifyOptions
        +AddScrubber(scrubfn scrubber) verifyOptions
        +Scrub(reader io.Reader) io.Reader, error
    }

    class fileOptions {
        -fields map
        +WithExtension(extensionWithDot string) verifyOptions
        +GetExtension() string
        +WithNamer(namer ApprovalNamerCreator) verifyOptions
        +GetNamer() ApprovalNamerCreator
        +WithAdditionalInformation(info string) verifyOptions
    }

    class scrubber {
        <<function type>>
        string to string function
    }

    class FirstWorkingReporter {
        -Reporters Reporter array
        +Report(approved string, received string) bool
    }

    class MultiReporter {
        -Reporters Reporter array
        +Report(approved string, received string) bool
    }

    class templatedCustomNamer {
        -template string
        -testSourceDirectory string
        -relativeTestSourceDirectory string
        -approvalsSubdirectory string
        -testFileName string
        -testCaseName string
        -additionalInformation string
        +GetName() string
        +GetReceivedFile(extWithDot string) string
        +GetApprovalFile(extWithDot string) string
    }

    class VerifyFunctions {
        <<module>>
        +Verify(t Failable, reader io.Reader, opts)
        +VerifyString(t Failable, s string, opts)
        +VerifyJSONStruct(t Failable, obj interface, opts)
        +VerifyJSONBytes(t Failable, bs bytes, opts)
        +VerifyXMLStruct(t Failable, obj interface, opts)
        +VerifyXMLBytes(t Failable, bs bytes, opts)
        +VerifyMap(t Failable, m interface, opts)
        +VerifyArray(t Failable, array interface, opts)
        +VerifyAll(t Failable, header string, collection interface, transform func, opts)
    }

    class CombinationApprovals {
        <<module>>
        +VerifyAllCombinationsFor1(t Failable, header string, transform func, collection1 interface)
        +VerifyAllCombinationsFor2(t Failable, header string, transform func, collection1 interface, collection2 interface)
        +VerifyAllCombinationsFor3 through For9
    }

    class ScrubberFunctions {
        <<module>>
        +CreateRegexScrubber(regex regexp, replacer string) scrubber
        +CreateRegexScrubberWithLabeler(regex regexp, replacer func) scrubber
        +CreateNoopScrubber() scrubber
        +CreateMultiScrubber(scrubbers) scrubber
        +CreateGuidScrubber() scrubber
    }

    class ReporterFactory {
        <<module>>
        +NewFirstWorkingReporter(reporters) Reporter
        +NewMultiReporter(reporters) Reporter
        +NewDiffReporter() Reporter
        +NewFrontLoadedReporter() Reporter
    }

    class GlobalConfiguration {
        <<module>>
        +UseReporter(reporter Reporter) io.Closer
        +UseFrontLoadedReporter(reporter Reporter) io.Closer
        +UseFolder(f string)
    }

    %% Relationships
    VerifyFunctions ..> Failable : uses
    VerifyFunctions ..> verifyOptions : uses
    VerifyFunctions ..> ApprovalNamer : creates via
    VerifyFunctions ..> Reporter : uses

    CombinationApprovals ..> Failable : uses
    CombinationApprovals ..> VerifyFunctions : calls

    verifyOptions --> fileOptions : creates
    verifyOptions ..> scrubber : uses
    fileOptions ..> ApprovalNamer : creates

    FirstWorkingReporter ..|> Reporter : implements
    MultiReporter ..|> Reporter : implements
    templatedCustomNamer ..|> ApprovalNamer : implements

    ReporterFactory --> FirstWorkingReporter : creates
    ReporterFactory --> MultiReporter : creates

    ScrubberFunctions --> scrubber : creates

    GlobalConfiguration ..> Reporter : configures
```

## Component Descriptions

### Core Interfaces

- **Failable**: Abstraction over testing.T, provides test failure and naming capabilities
- **ApprovalNamer**: Manages the naming convention for approved and received files
- **Reporter**: Handles reporting when approval tests fail

### Configuration

- **verifyOptions**: Main configuration object for verification operations
- **fileOptions**: File-specific configuration options
- **scrubber**: Function type for data sanitization/scrubbing

### Verification

- **VerifyFunctions**: Core verification methods that form the main API
- **CombinationApprovals**: Specialized functions for testing combinations of inputs

### Reporters

- **FirstWorkingReporter**: Tries multiple reporters in sequence until one works
- **MultiReporter**: Runs multiple reporters simultaneously

### Naming

- **templatedCustomNamer**: Template-based implementation of ApprovalNamer

### Utilities

- **ScrubberFunctions**: Factory methods for creating various scrubber implementations
- **ReporterFactory**: Factory methods for creating reporter instances
- **GlobalConfiguration**: Global configuration functions for setting default reporters and folders

## Key Architectural Patterns

1. **Interface-based Design**: Core functionality is defined through interfaces (Failable, ApprovalNamer, Reporter)
2. **Options Pattern**: Flexible configuration through verifyOptions and fileOptions
3. **Strategy Pattern**: Different reporting strategies through Reporter interface
4. **Chain of Responsibility**: FirstWorkingReporter tries reporters in sequence
5. **Template Method**: templatedCustomNamer uses templates for file naming
6. **Functional Options**: scrubber as a function type allows flexible data transformation
