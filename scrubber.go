package approvals

import "regexp"

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
