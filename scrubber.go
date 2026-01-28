package approvals

import (
	"fmt"
	"regexp"
	"strconv"
)

type scrubber func(s string) string

// Deprecated: WithRegexScrubber allows you to 'scrub' dynamic data such as timestamps within your test input
// and replace it with a static placeholder
func (v VerifyOptions) WithRegexScrubber(regex *regexp.Regexp, replacer string) VerifyOptions {
	return v.AddScrubber(func(s string) string {
		return regex.ReplaceAllString(s, replacer)
	})
}

// CreateRegexScrubber allows you to create a scrubber that uses a regular expression to scrub data
func CreateRegexScrubber(regex *regexp.Regexp, replacer string) scrubber {
	return CreateRegexScrubberWithLabeler(regex, func(int) string { return replacer })
}

// CreateRegexScrubberWithLabeler allows you to create a scrubber that uses a regular expression to scrub data
func CreateRegexScrubberWithLabeler(regex *regexp.Regexp, replacer func(int) string) scrubber {
	return func(s string) string {
		seen := map[string]int{}
		replacefn := func(s string) string {
			idx, ok := seen[s]
			if !ok {
				idx = len(seen)
				seen[s] = idx
			}
			return replacer(idx + 1)
		}
		return regex.ReplaceAllStringFunc(s, replacefn)
	}
}

func CreateJSONScrubber(label string, valueMatcher *regexp.Regexp) scrubber {
	return CreateRegexScrubberWithLabeler(regexp.MustCompile(fmt.Sprintf(`"%s": \d+`, label)), func(n int) string { return fmt.Sprintf(`"%s": "<%s_%d>"`, label, label, n) })
}

// NoopScrubber is a scrubber that does nothing
func CreateNoopScrubber() scrubber {
	return func(s string) string {
		return s
	}
}

// CreateMultiScrubber allows you to chain multiple scrubbers together
func CreateMultiScrubber(scrubbers ...scrubber) scrubber {
	return func(s string) string {
		for _, scrubber := range scrubbers {
			s = scrubber(s)
		}
		return s
	}
}

func CreateGuidScrubber() scrubber {
	regex := regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")
	return CreateRegexScrubberWithLabeler(regex, func(n int) string { return "guid_" + strconv.Itoa(n) })
}
