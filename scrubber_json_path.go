package approvals

import (
	"strings"

	"github.com/bitly/go-simplejson"
)

// converts received JSON string into a simpleJSON struct, scrubs the path with passed "scrubVal" using a separate util function, turns it into a string again
// (pretty naive implementation)
func scrubJSONPath(jsonStr string, path string, scrubVal interface{}) string {
	js, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		panic(err)
	}

	scrubJSONPathSimpleJSON(js, path, scrubVal)

	resS, err := js.EncodePretty()
	if err != nil {
		panic(err)
	}
	return string(resS)
}

// limited implementation of JSONPath for scrubbing values in simplejson structs:
//
// ScrubJSONPath(data, "foo.bar", "<bar_scrubbed>")
// { foo: { bar: "123" } }                   -> { foo: { bar: "<bar_scrubbed>" }
//
// - ScrubJSONPath(data, "foo[*].bar", "<bar_arr_scrubbed>")
// { foo: [{ bar: "123" }, { bar: "234" }] } -> { foo: [{ bar: "<bar_arr_scrubbed>" }, { bar: "<bar_arr_scrubbed>" }]) }
func scrubJSONPathSimpleJSON(data *simplejson.Json, path string, scrubVal interface{}) {
	// "foo[*].bar" -> ["foo", "[*]", "bar"]
	var elems []string
	for _, elDot := range strings.Split(path, ".") {
		if elDot == "[*]" {
			elems = append(elems, "[*]")
		} else if strings.HasSuffix(elDot, "[*]") {
			elems = append(elems, strings.TrimSuffix(elDot, "[*]"))
			elems = append(elems, "[*]")
		} else {
			elems = append(elems, elDot)
		}
	}

	// end of recursive iteration, set the value (if key is present) and return:
	if len(elems) == 1 {
		if _, found := data.CheckGet(elems[0]); found {
			data.Set(elems[0], scrubVal)
		}
		return
	}

	leftPath := path
	for _, el := range elems {
		leftPath = strings.TrimPrefix(leftPath, el)
		leftPath = strings.TrimPrefix(leftPath, ".")

		kjkj // special key for array object keys:
		if el == "[*]" {
			for i := 0; i < len(data.MustArray()); i++ {
				scrubJSONPathSimpleJSON(data.GetIndex(i), leftPath, scrubVal)
			}
		} else {
			if _, found := data.CheckGet(el); found {
				scrubJSONPathSimpleJSON(data.Get(el), leftPath, scrubVal)
			}
		}
	}
}
