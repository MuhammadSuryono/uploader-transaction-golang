package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func WriteErrorLog(str string) {
	f := openFile("log_data")
	writeLog(f, str)
}

func LogConsole(filename string, msg string) {
	f := openFile(filename)
	writeLog(f, msg)
}

func writeLog(f *os.File, str string) {
	log.SetOutput(f)
	log.Println(str)
	defer f.Close()
}

func openFile(filename string) *os.File {
	if strings.Contains(filename, ".log") {
		spl := strings.Split(filename, ".")
		filename = ""

		for i, s := range spl {
			if (len(spl) - 1) != i {
				filename += s
			}
		}
	}
	f, err := os.OpenFile(filename+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("=> Error opening file: %v", err)
	}

	return f
}

func CsvWrite(fname string, column []string) {
	// read the file
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	w := csv.NewWriter(f)
	w.Write(column)
	w.Flush()
}
