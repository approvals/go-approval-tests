package approvals

import (
	"fmt"
	"regexp"
)

type DateScrubber struct {
	pattern     string
	replacement func(int) string
}

func NewDateScrubber(pattern string) scrubber {
	return CreateRegexScrubberWithLabeler(regexp.MustCompile(pattern), func(n int) string {
		return fmt.Sprintf(`[Date%d]`, n)
	})
}

type SupportedFormat struct {
	Regex    string
	Examples []string
}

func GetSupportedFormats() []SupportedFormat {
	return []SupportedFormat{
		{`[a-zA-Z]{3} [a-zA-Z]{3} \d{2} \d{2}:\d{2}:\d{2}`, []string{`Tue May 13 16:30:00`}},
		{`[a-zA-Z]{3} [a-zA-Z]{3} \d{2} \d{2}:\d{2}:\d{2} [a-zA-Z]{3,4} \d{4}`, []string{`Wed Nov 17 22:28:33 EET 2021`}},
		{`(Mon|Tue|Wed|Thu|Fri|Sat|Sun), \d{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) \d{4} \d{2}:\d{2}:\d{2} GMT`, []string{`Wed, 21 Oct 2015 07:28:00 GMT`}},
		{`[a-zA-Z]{3} [a-zA-Z]{3} \d{2} \d{4} \d{2}:\d{2}:\d{2}.\d{3}`, []string{`Tue May 13 2014 23:30:00.789`}},
		{`[a-zA-Z]{3} [a-zA-Z]{3} \d{2} \d{2}:\d{2}:\d{2} -\d{4} \d{4}`, []string{`Tue May 13 16:30:00 -0800 2014`}},
		{`\d{2} [a-zA-Z]{3} \d{4} \d{2}:\d{2}:\d{2},\d{3}`, []string{`13 May 2014 23:50:49,999`}},
		{`[A-Za-z]{3} \d{2} \d{2}:\d{2}`, []string{`Oct 13 15:29`}},
		{`[a-zA-Z]{3} \d{2}, \d{4} \d{2}:\d{2}:\d{2} [a-zA-Z]{2} [a-zA-Z]{3}`, []string{`May 13, 2014 11:30:00 PM PST`}},
		{`\d{2}:\d{2}:\d{2}`, []string{`23:30:00`}},
		{`\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}(\.\d{3})?`, []string{`2014/05/13 16:30:59.786`, `2014/05/13 16:30:59`}},
		{`\d{4}-\d{1,2}-\d{1,2}T\d{1,2}:\d{2}Z`, []string{`2020-9-10T08:07Z`, `2020-09-9T08:07Z`, `2020-09-10T8:07Z`, `2020-09-10T08:07Z`}},
		{`\d{4}-\d{1,2}-\d{1,2}T\d{1,2}:\d{2}:\d{2}Z`, []string{`2020-09-10T08:07:89Z`}},
		{`\d{4}-\d{1,2}-\d{1,2}T\d{1,2}:\d{2}\:\d{2}\.\d{3}Z`, []string{`2020-09-10T01:23:45.678Z`}},
		{`\d{8}T\d{6}Z`, []string{`20210505T091112Z`}},
		{`\d{4}-\d{2}-\d{2}`, []string{`2024-12-17`}},
		{`\d{4}-\d{1,2}-\d{1,2}T\d{1,2}:\d{2}:\d{2}(\.\d{1,9})?Z`, []string{`2024-12-18T14:04:46.746130Z`, `2024-12-18T14:04:46Z`, `2024-12-18T14:04:46.746130834Z`}},
		{`\d{2}[-/.]\d{2}[-/.]\d{4}\s\d{2}:\d{2}(:\d{2})?( (?:pm|am|PM|AM))?`, []string{`13/05/2014 23:50:49`, `13.05.2014 23:50:49`, `13-05-2014 23:50:49`, `13.05.2014 23:50`, `05/13/2014 11:50:49 PM`}},
	}
}

func GetSupportedFormatsRegex() []string {
	formats := GetSupportedFormats()
	regexList := make([]string, len(formats))
	for i, format := range formats {
		regexList[i] = format.Regex
	}
	return regexList
}

// public static DateScrubber getScrubberFor(String formattedExample)
//
//	{
//	  for (SupportedFormat pattern : getSupportedFormats())
//	  {
//	    DateScrubber scrubber = new DateScrubber(pattern.getRegex());
//	    if ("[Date1]".equals(scrubber.scrub(formattedExample)))
//	    { return scrubber; }
//	  }
//	  throw new FormattedException(
//	      "No match found for %s.\n Feel free to add your date at https://github.com/approvals/ApprovalTests.Java/issues/112 \n Current supported formats are: %s",
//	      formattedExample, Query.select(getSupportedFormats(), SupportedFormat::getRegex));
//	}
func GetDateScrubberFor(formattedExample string) (scrubber, error) {
	for _, pattern := range GetSupportedFormats() {
		scrubber := NewDateScrubber(pattern.Regex)
		if "[Date1]" == scrubber(formattedExample) {
			return scrubber, nil
		}
	}
	return nil, fmt.Errorf(
		"No match found for %s.\n Feel free to add your date at https://github.com/approvals/ApprovalTests.Java/issues/112 \n Current supported formats are: %v",
		formattedExample, GetSupportedFormatsRegex())
}
