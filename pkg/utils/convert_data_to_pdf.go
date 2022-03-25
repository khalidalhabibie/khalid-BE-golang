package utils

import (
	"fmt"
	"gokes/app/models"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func ConvertDataDataFakesToPDF(fakesM models.Fakes) bool {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)

	pdf.Text(25, 20, "Faskes (Fasilitas Kesehatan) by KEMENKES")

	code := fmt.Sprintf("Code : %s", fakesM.Code)
	name := fmt.Sprintf("Name : %s", fakesM.Name)
	status := fmt.Sprintf("Type : %s", fakesM.Type)
	description := fmt.Sprintf("Description : %s", fakesM.Description)
	nakesCount := fmt.Sprintf("Nakes Count : %d", fakesM.NakesCount)

	pdf.SetFont("Arial", "", 16)
	pdf.Text(20, 50, code)
	pdf.Text(20, 60, name)
	pdf.Text(20, 70, status)
	pdf.Text(20, 80, description)
	pdf.Text(20, 90, nakesCount)

	err := pdf.OutputFileAndClose(fmt.Sprintf("data/%v.pdf", fakesM.Code))
	if err != nil {
		log.Println("error convert data  : ", err)
		return true
	}

	return false

}
