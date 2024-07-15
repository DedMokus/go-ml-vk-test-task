package processor

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/DedMokus/go-ml-vk-test-task/internal/db"
	"github.com/DedMokus/go-ml-vk-test-task/internal/document"
)

type Processor interface {
	Process(d *document.Document) (*document.Document, error)
}

type QueueProcessor struct {
	db *db.PostgreSQLProcessor
	m  *sync.Mutex
}

func CreateQueueProcessor(db *db.PostgreSQLProcessor, mu *sync.Mutex) *QueueProcessor {
	return &QueueProcessor{
		db: db,
		m:  mu,
	}
}

func (p QueueProcessor) Process(d *document.Document) (*document.Document, error) {
	var docs *sql.Rows

	//Создаем запрос на документы с нужным Url
	var query string = fmt.Sprintf("SELECT * FROM docs WHERE url = '%s';", d.Url)
	fmt.Printf("Document to process: %s\n", d.String())

	//Читаем строки из базы данных
	docs, err := p.db.Query(query)
	if err != nil {
		log.Fatalf("Error query db: %v", err)
	}
	defer docs.Close()

	var PubDateToGo, FetchTimeToGo, FirstFetchTimeToGo uint64
	var TextToGo string
	PubDateToGo = d.PubDate
	FetchTimeToGo = d.FetchTime
	FirstFetchTimeToGo = d.FirstFetchTime
	if FirstFetchTimeToGo == 0 {
		FirstFetchTimeToGo = FetchTimeToGo
	}
	for docs.Next() {
		var id int
		var Url, Text string
		var PubDate, FetchTime, FirstFetchTime uint64

		err := docs.Scan(&id, &Url, &PubDate, &FetchTime, &Text, &FirstFetchTime)
		//fmt.Printf("\nPrint next row: Url= %s, pubdate=%d", Url, PubDate)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			panic(err)
		}

		//Проверка на дубликат
		if Text == d.Text && FetchTime == d.FetchTime && FirstFetchTime == d.FirstFetchTime && PubDate == d.PubDate {
			return nil, nil
		}

		if FetchTime < FetchTimeToGo {
			FirstFetchTimeToGo = FetchTime
			PubDateToGo = PubDate
		}
		if FetchTime > FetchTimeToGo {
			FetchTimeToGo = FetchTime
			TextToGo = Text
		}
	}
	if err := docs.Err(); err != nil {
		log.Fatalf("Error after read rows: %v", err)
		panic(err)
	}

	d.Text = TextToGo
	d.FetchTime = FetchTimeToGo
	d.PubDate = PubDateToGo
	d.FirstFetchTime = FirstFetchTimeToGo

	fmt.Printf("Document after process: %s\n", d.String())

	return d, nil

}
