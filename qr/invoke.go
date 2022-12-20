package qr

import (
	"bytes"
	"fmt"
	"github.com/MuhammadSuryono/module-golang-server/http/response"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math"
	"path/filepath"
	"sync"
	"time"
	"upoader-golang/database"
)

type requestUploadFile struct {
	Period string `json:"period" form:"period" validate:"required"`
	Part   int    `json:"part" form:"part" validate:"required"`
}

var nameConsole string

func InvokeRequest(c *gin.Context) {
	var params requestUploadFile
	errBind := c.ShouldBind(&params)
	if errBind != nil {
		c.JSON(response.BAD_REQUEST_CODE, response.FailureResponse(
			response.BAD_REQUEST_STATUS,
			"Please check data request must be not null",
			params))
		return
	}

	f, errFile := c.FormFile("file")
	if errFile != nil {
		c.JSON(response.BAD_REQUEST_CODE, response.FailureResponse(
			response.BAD_REQUEST_STATUS,
			"File not found",
			nil))
		return
	}
	filename := filepath.Base(f.Filename)
	nameConsole = "QR_CONSOLE_" + filename

	start := time.Now()
	//util.LogConsole(nameConsole, fmt.Sprintf("=> Started at %v", start))

	buf := bytes.NewBuffer(nil)
	mFile, _ := f.Open()
	_, err := io.Copy(buf, mFile)

	if err != nil {
		c.JSON(response.BAD_REQUEST_CODE, response.FailureResponse(
			response.BAD_REQUEST_STATUS,
			"Error copy file",
			err))
		return
	}

	db, errDb := database.DBOpenConnection()
	if errDb != nil {
		//util.LogConsole(nameConsole, fmt.Sprintf("=> Error connection DB %v", errDb.Error()))
		return
	}

	csvReader, err := openCsvFile(buf.Bytes())
	if err != nil {
		//util.LogConsole(nameConsole, fmt.Sprintf("=> Error open file CSV %v", errDb.Error()))
		return
	}

	defer mFile.Close()
	r, _ := csvReader.Read()
	errValidate := validateHeader(r)
	if errValidate != nil {
		return
	}

	jobs := make(chan [][]interface{}, 0)
	wg := new(sync.WaitGroup)

	go dispatchWorkers(db, jobs, wg)

	go func() {
		//util.LogConsole(nameConsole, "=> Read csv FilePerLine")
		log.Println("=> Waiting upload.....")
		readCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)
		wg.Wait()

		duration := time.Since(start)
		log.Println(fmt.Sprintf("=> Done in %d Minute", int(math.Ceil(duration.Seconds()))/60))
	}()
	c.JSON(response.SUCCESS_CODE, response.FailureResponse(
		response.SUCCESS_STATUS,
		"Data QR still uploading",
		nil))
	return
}
