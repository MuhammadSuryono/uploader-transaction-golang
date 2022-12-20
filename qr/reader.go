package qr

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
)

const csvFile = "QREEXXx012022.csv"

func openCsvFile(i []byte) (*csv.Reader, error) {
	log.Println("=> Open csv File")
	reader := csv.NewReader(bytes.NewBuffer(i))
	return reader, nil
}

func readCsvFilePerLineThenSendToWorker(csvReader *csv.Reader, jobs chan<- [][]interface{}, wg *sync.WaitGroup) {

	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if len(firstRow) == 0 {
			if len(DataHeaderCsvValidate) != len(row) {
				log.Println("Header format not same with data header validation")
				break
			}
			firstRow = row
			continue
		}

		rowOrdered := make([]interface{}, 0)
		dataRowI := make([]interface{}, 0)
		for index, each := range row {
			if !contains(skipIndex, index) {
				var dt interface{}
				if strings.Contains(each, ",") {
					sp := strings.Split(each, ",")
					dt, _ = strconv.Atoi(sp[0])
				} else {
					if contains(indexInteger, index) && each == "" {
						each = "0"
					}
					dt = each
				}
				rowOrdered = append(rowOrdered, dt)
			}
			dataRowI = append(dataRowI, each)
		}

		giveData := make([][]interface{}, 0)
		giveData = append(giveData, rowOrdered)
		giveData = append(giveData, dataRowI)

		wg.Add(1)
		jobs <- giveData
	}
	log.Println("Close jobs")
	close(jobs)
}
