package qr

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"upoader-golang/util"
)

func doTheJob(db *sql.DB, mysql *sql.DB, params requestUploadFile, values [][]interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			conn, err := db.Conn(context.Background())
			query := fmt.Sprintf("INSERT INTO transaction_qr (%s) VALUES (%s)",
				strings.Join(dataHeaders, ","),
				strings.Join(generateQuestionsMark(len(dataHeaders)), ","),
			)

			_, err = conn.ExecContext(context.Background(), query, values[0]...)
			str := make([]string, 0)
			for _, dt := range values[1] {
				str = append(str, fmt.Sprintf("%v", dt))
			}
			if err != nil {
				str := make([]string, 0)
				for _, dt := range values[1] {
					str = append(str, fmt.Sprintf("%v", dt))
				}
				util.CsvWrite(fmt.Sprintf("%s_part_%d.csv", params.Period, params.Part), str)
				util.WriteErrorLog(err.Error())
			}

			err = conn.Close()
			if err != nil {
				log.Println("Error close DB", err.Error())
			}
		}(&outerError)
		if outerError == nil {
			break
		}
	}
}
