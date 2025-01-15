package approvals

import (
	"regexp"
	"strconv"
)

type scrubber func(s string) string

// Deprecated: WithRegexScrubber allows you to 'scrub' dynamic data such as timestamps within your test input
// and replace it with a static placeholder
func (v verifyOptions) WithRegexScrubber(regex *regexp.Regexp, replacer string) verifyOptions {
	return v.AddScrubber(func(s string) string {
		return regex.ReplaceAllString(s, replacer)
	})
}

// CreateRegexScrubber allows you to create a scrubber that uses a regular expression to scrub data
func CreateRegexScrubber(regex *regexp.Regexp, replacer string) scrubber {
	return func(s string) string {
		return regex.ReplaceAllString(s, replacer)
	}
}

// CreateRegexScrubberWithLabeler allows you to create a scrubber that uses a regular expression to scrub data
func CreateRegexScrubberWithLabeler(regex *regexp.Regexp, replacer func(int) string) scrubber {
	return func(s string) string {
		m := map[string]int{}
		replacefn := func(s string) string {
			idx := 0
			if i, ok := m[s]; ok {
				idx = i
			} else {
				idx = len(m)
				m[s] = idx
			}

			return replacer(idx)
		}
		return regex.ReplaceAllStringFunc(s, replacefn)
	}
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
