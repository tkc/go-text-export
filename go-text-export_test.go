package go_text_export

import (
	"log"
	"testing"
)

func TestExport(t *testing.T) {

	str := `
     Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC,
     making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia,
     looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature,
     discovered the undoubtable source. Lorem Ipsum comes from sections
     1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC.
     This book is a treatise on the theory of ethics, very popular during the Renaissance.
     The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.
    `
	res := Create(str).Export(100).SetMoreText(" more..").RemoveWhiteSpace().String()
	log.Print(res)
}

func TestTextExport_Len(t *testing.T) {

	str := `
     Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC,
     making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia,
     looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature,
     discovered the undoubtable source. Lorem Ipsum comes from sections
     1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC.
     This book is a treatise on the theory of ethics, very popular during the Renaissance.
     The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.
    `
	res := Create(str).Export(100).Len()
	log.Print(res)
}

func TestTextExport_ConvertLinkHtml(t *testing.T) {

	str := `
	link
    	https://example.com
    `
	res := Create(str).Export(100).ConvertHtmlLink().String()
	log.Print(res)
}

func TestTextExport_Prepare(t *testing.T) {

	str := `
	text text text
    `
	short := Create(str).Prepare(30).String()
	log.Print(short)

	over := Create(str).Prepare(10).String()
	log.Print(over)
}