package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func init() {
	//loads values from .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file!")
	}

}

func main() {
	numThreads, exists := os.LookupEnv("NUM_THREADS")
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	inChannel := make(chan Document, 10)
	outChannel := make(chan Document, 100)

	//Заполняем имитацию очереди рандомными документами
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go generateDocument(inChannel, wg, mu)

	}
	wg.Wait()

	//Обрабатываем очередь

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateDocument(in chan Document, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	var tUrl, TText string
	var TPubDate, TFetchTime, TFirstFetchTime uint64
	tUrl = string(rand.Int63n(10))
	TText = RandStringRunes(10)
	TPubDate = uint64(rand.Int63n(1000))
	TFetchTime = uint64(rand.Int63n(1000))
	TFetchTime = 0

	doc := Document{
		Url:            tUrl,
		PubDate:        TPubDate,
		FetchTime:      TFetchTime,
		Text:           TText,
		FirstFetchTime: TFirstFetchTime,
	}
	mu.Lock()
	in <- doc
	mu.Unlock()
}

type Document struct {
	Url            string
	PubDate        uint64
	FetchTime      uint64
	Text           string
	FirstFetchTime uint64
}

type Database interface {
	ProcessCommand()
}

type DBAccess struct{}

func (db DBAccess) ProcessCommand(command string)
