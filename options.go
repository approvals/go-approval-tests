package approvals

import (
	"io"
	"strings"
)

// verifyOptions can be accessed via the approvals.Options() API enabling configuration of scrubbers
type verifyOptions struct {
	fields map[string]interface{}
}

type fileOptions struct {
	fields map[string]interface{}
}

func (v verifyOptions) ForFile() fileOptions {
	return fileOptions{fields: v.fields}
}

// Deprecated: Use `ForFile().WithExtension(extension)` instead.
func (v verifyOptions) WithExtension(extension string) verifyOptions {
	return v.ForFile().WithExtension(extension)
}

// Deprecated: Use `ForFile().GetExtension()` instead.
func (v verifyOptions) GetExtension() string {
	return v.ForFile().GetExtension()
}

func (f fileOptions) GetExtension() string {
	ext := getField(f.fields, "extWithDot", ".txt")
	return ext.(string)
}

func (f fileOptions) GetNamer(t Failable) *ApprovalName {
	ext := getField(f.fields, "namer", getApprovalName(t))
	return ext.(*ApprovalName)
}

func (v verifyOptions) getField(key string, defaultValue interface{}) interface{} {
	return getField(v.fields, key, defaultValue)
}

func (f fileOptions) getField(key string, defaultValue interface{}) interface{} {
	return getField(f.fields, key, defaultValue)
}

func getField(fields map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if fields == nil {
		return defaultValue
	}
	if value, ok := fields[key]; ok {
		return value
	}
	return defaultValue
}

// Options enables providing individual Verify functions with customisations such as scrubbers
func Options() verifyOptions {
	return verifyOptions{}
}

// WithScrubber allows you to 'scrub' data within your test input and replace it with a static placeholder
func (v verifyOptions) WithScrubber(scrub scrubber) verifyOptions {
	return NewVerifyOptions(v.fields, "scrubber", scrub)
}

// AddScrubber allows you to 'scrub' data within your test input and replace it with a static placeholder
func (v verifyOptions) AddScrubber(scrubfn scrubber) verifyOptions {
	scrub := CreateMultiScrubber(v.getField("scrubber", CreateNoopScrubber()).(scrubber), scrubfn)
	return v.WithScrubber(scrub)
}

// WithExtension overrides the default file extension (.txt) for approval files.
func (f fileOptions) WithExtension(extensionWithDot string) verifyOptions {
	if !strings.HasPrefix(extensionWithDot, ".") {
		extensionWithDot = "." + extensionWithDot
	}
	return NewVerifyOptions(f.fields, "extWithDot", extensionWithDot)
}

func (v verifyOptions) Scrub(reader io.Reader) (io.Reader, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	scrub := v.getField("scrubber", CreateNoopScrubber()).(scrubber)
	result := scrub(string(b))

	return strings.NewReader(result), nil
}

func NewVerifyOptions(fields map[string]interface{}, key string, value interface{}) verifyOptions {
	// Make a copy of the fields map, but with the new key and value
	newFields := make(map[string]interface{}, len(fields))
	for k, v := range fields {
		newFields[k] = v
	}
	newFields[key] = value
	return verifyOptions{
		fields: newFields,
	}
}

func alwaysOption(opts []verifyOptions) verifyOptions {
	var v verifyOptions
	if len(opts) == 0 {
		v = Options()
	} else {
		v = opts[0]
	}

	return v
}
