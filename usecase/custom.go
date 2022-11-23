package usecase

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strconv"

	"samsul96maarif/github.com/go-api-app/request"

	"github.com/xuri/excelize/v2"
)

const (
	SHEET_NAME = "Sheet1"
	COL_DESC   = 5
)

func parsing(fi multipart.File) []request.D {
	var entities []request.D
	f, err := excelize.OpenReader(fi)
	if err != nil {
		fmt.Println("err", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("err", err)
		}
	}()
	rows, err := f.GetRows(SHEET_NAME)
	if err != nil {
		log.Print("error", err)
	}
	for i := 1; i < len(rows); i++ {
		nib, err := strconv.Atoi(rows[i][3])
		if err != nil {
			log.Println("error", err)
		}
		ket := rows[i][COL_DESC]
		entities = append(entities, request.D{Nib: uint(nib), Desc: ket})
	}
	return entities
}

func (usecase *Usecase) Im(ctx context.Context, req request.Im) (xlsx *excelize.File, err error) {
	var res []request.D
	buff := make([]byte, 512)
	_, err = req.File.Read(buff)
	if err != nil {
		fmt.Println("error", err)
	}

	_, err = req.File.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("err", err)
	}
	da := parsing(req.File)
	temp := parsing(req.Compare)
	for _, e := range da {
		var keteranga string
		for _, t := range temp {
			if e.Nib == t.Nib {
				keteranga = t.Desc
				break
			}
		}
		res = append(res, request.D{Nib: e.Nib, Desc: keteranga})
	}

	type M map[string]interface{}

	sheet1Name := "coa"
	xlsx = excelize.NewFile()
	for i, each := range res {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each.Nib)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each.Desc)
	}

	err = xlsx.SaveAs("./file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	return xlsx, err
}
