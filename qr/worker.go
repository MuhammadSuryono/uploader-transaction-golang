package qr

import (
	"database/sql"
	"log"
	"sync"
)

const totalWorker = 100

func dispatchWorkers(db *sql.DB, jobs <-chan [][]interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		log.Println("Dispath worker", workerIndex, "is already")
		go func(workerIndex int, db *sql.DB, jobs <-chan [][]interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				doTheJob(db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}
}
