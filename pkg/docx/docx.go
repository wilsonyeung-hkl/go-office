package docx

import (
	"fmt"
	"io/ioutil"
	"log"

	"baliance.com/gooxml/measurement"
	"github.com/carmel/gooxml/common"
	"github.com/carmel/gooxml/document"
)

func ImageToDocx(imagePath string, savePath string) {

	fmt.Println("ImageToDocx...ing")
	doc := document.New()
	// Adding a Paragraph

	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText("This is a paragraph of text.")

	// Adding a Heading

	heading := doc.AddParagraph()
	heading.SetStyle("Heading1")
	run = heading.AddRun()
	run.AddText("This is a Heading 1")

	// Adding an Image
	img2data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatalf("unable to read file: %s", err)
	}
	img2, err := common.ImageFromBytes(img2data)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	img2ref, err := doc.AddImage(img2)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}
	para = doc.AddParagraph()

	run = para.AddRun()
	inl, err := run.AddDrawingInline(img2ref)
	if err != nil {
		log.Fatalf("unable to add inline image: %s", err)
	}
	inl.SetSize(1*measurement.Inch, 1*measurement.Inch)
	doc.SaveToFile(savePath)
}
