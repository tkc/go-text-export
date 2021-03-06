package go_text_export

import (
	"fmt"
	"strings"
	"regexp"
	"unicode/utf8"
)

const moreStr = "..."

type TextExport struct {
	Str     string
	moreStr string
	moreUrl string
}

func Create(str string) *TextExport {
	return &TextExport{Str:str, moreStr:moreStr}
}

func (te *TextExport) SetMoreText(moreStr string) *TextExport {
	te.moreStr = moreStr
	return te
}

func (te *TextExport) SetMoreUrl(moreUrl string) *TextExport {
	te.moreUrl = moreUrl
	return te
}

func (te *TextExport) RemoveMoreText() *TextExport {
	te.moreStr = ""
	return te
}

func (te *TextExport) RemoveNewLine() *TextExport {
	str := strings.Replace(te.Str, "\n", "", -1)
	te.Str = str
	return te
}

func (te *TextExport) RemoveWhiteSpace() *TextExport {
	str := strings.Replace(te.Str, "[ ]+", "", -1)
	te.Str = str
	return te
}

func (te *TextExport) ConvertHtmlLink() *TextExport {
	re, _ := regexp.Compile(`https?://[-_.!~*\'()a-zA-Z0-9;/?:@&=+$,%#]+`)
	cb := func(s string) string {
		return "<a href=" + s + ">" + s + "</a>"
	}
	str := re.ReplaceAllStringFunc(te.Str, cb)
	te.Str = str
	return te
}

func (te *TextExport) ConvertHtmlBrTag() *TextExport {
	str := strings.Replace(te.Str, "\n", "<br/>", -1)
	te.Str = str
	return te
}

func (te *TextExport)Export(num int) *TextExport {
	te.Str = te.RemoveNewLine().Str // TODO
	if (utf8.RuneCountInString(te.Str) > num) {
		regStr := fmt.Sprintf(`^(.){%d}`, num)
		r := regexp.MustCompile(regStr)
		result := r.FindAllStringSubmatch(te.Str, -1)
		te.Str = result[0][0]
		return te
	} else {
		return te.RemoveMoreText()
	}
}

func (te *TextExport)Prepare(num int) *TextExport {

	//removing more text
	te.RemoveMoreText()

	var shortage int = te.Len() - num
	if shortage < 0 {
		te.Str = te.RemoveNewLine().Str // TODO
		for i := 0; i < shortage * -1; i++ {
			te.Str = te.Str + "&nbsp;"
		}
		return te
	} else {
		return te.Export(shortage)
	}
}

func (te *TextExport)String() string {

	if te.moreUrl != "" {
		return te.Str + "<a href='" + te.moreUrl + "'>" + te.moreStr + "</a>"
	} else {
		return te.Str + te.moreStr
	}
}

func (te *TextExport)Len() int {
	return len(te.Str)
}
