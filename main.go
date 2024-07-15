package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/DedMokus/go-ml-vk-test-task/internal/db"
	"github.com/DedMokus/go-ml-vk-test-task/internal/document"
	"github.com/DedMokus/go-ml-vk-test-task/internal/processor"
)

func main() {

	//Собираем данные из env
	envnumThreads, exists := os.LookupEnv("NUM_THREADS")
	if !exists {
		envnumThreads = "1"
	}
	numThreads, err := strconv.Atoi(envnumThreads)
	if err != nil {
		log.Fatal("Error converting string:", err)
	}

	//Создаем мьютех
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	//Каналы (замена очереди kafka)
	var inChannel chan *document.Document = make(chan *document.Document, 10)
	var outChannel chan *document.Document = make(chan *document.Document, 100)

	//Подключение к базе данных
	data := new(db.PostgreSQLProcessor)
	data.Connect()

	//Заполнение бд рандомными документами для сравнения (используется один url для простоты)
	for i := 0; i < 30; i++ {
		data.QueryRow("qwet")
	}

	//Заполняем имитацию очереди рандомными документами
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go document.GenerateDocuments(inChannel, wg, mu, "qwet")

	}
	wg.Wait()

	//Обрабатываем очередь
	proc := processor.CreateQueueProcessor(data, mu)

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		//Создаем горутин по количеству выставленных потоков
		go func() {
			for {
				if len(inChannel) != 0 {
					mu.Lock()
					doc := <-inChannel
					mu.Unlock()

					doc, err = proc.Process(doc)
					if doc == nil {
						return
					}
					mu.Lock()
					outChannel <- doc
					mu.Unlock()

				} else {
					break
				}
			}
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Printf("Program end!")

}
