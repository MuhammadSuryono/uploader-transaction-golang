package qr

import (
	"bytes"
	"context"
	"fmt"
	"github.com/MuhammadSuryono/module-golang-server/http/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"math"
	"path/filepath"
	"sync"
	"time"
	"upoader-golang/database"
	"upoader-golang/util"
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

	mysql, err := database.DBOpenConnectionMysql()
	if err != nil {
		log.Println("Error DB Connection MYSQL", err.Error())
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

	util.CsvWrite(fmt.Sprintf("%s_part_%d.csv", params.Period, params.Part), DataHeaderCsvValidate)

	jobs := make(chan [][]interface{}, 0)
	wg := new(sync.WaitGroup)

	go dispatchWorkers(db, mysql, params, jobs, wg)

	conn, err := mysql.Conn(context.Background())
	if err != nil {
		log.Fatalln("Error", err.Error())
	}
	id := uuid.New().String()

	query := fmt.Sprintf("INSERT INTO logs (transaction_id,periode,part,start_at,status) VALUES ('%s','%s','%d','%v','Start')",
		id,
		params.Period,
		params.Part,
		start,
	)

	_, err = conn.ExecContext(context.Background(), query)
	if err != nil {
		log.Println("Error")
	}

	err = conn.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		//util.LogConsole(nameConsole, "=> Read csv FilePerLine")
		log.Println("=> Waiting upload.....")
		readCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)
		wg.Wait()

		duration := time.Since(start)
		log.Println(fmt.Sprintf("=> Done in %d Minute", int(math.Ceil(duration.Seconds()))/60))
		queryUpdate := fmt.Sprintf("UPDATE logs set finish_at = '%s', status = 'Finish'",
			fmt.Sprintf("%v", duration),
		)
		conn, err := mysql.Conn(context.Background())
		_, err = conn.ExecContext(context.Background(), queryUpdate)
		if err != nil {
			log.Println("Error")
		}

		err = conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	c.JSON(response.SUCCESS_CODE, response.FailureResponse(
		response.SUCCESS_STATUS,
		"Data QR still uploading",
		nil))
	return
}
