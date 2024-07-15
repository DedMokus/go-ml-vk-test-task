package document

import (
	"fmt"
	"math/rand"
	"sync"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateDocuments(in chan *Document, wg *sync.WaitGroup, mu *sync.Mutex, q string) {
	defer wg.Done()
	var tUrl, TText string
	var TPubDate, TFetchTime, TFirstFetchTime uint64
	tUrl = q
	TText = RandStringRunes(10)
	TPubDate = uint64(rand.Int63n(1000))
	TFetchTime = uint64(rand.Int63n(1000))
	TFirstFetchTime = 0

	doc := &Document{
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

func (d *Document) String() string {
	var out string = fmt.Sprintf("Url: %s\t PubDate: %d\t FetchTime: %d\t FirstFetchTime: %d, Text: %s", d.Url, d.PubDate, d.FetchTime, d.FirstFetchTime, d.Text)
	return out
}

func GenerateRandomDocument(u string) Document {
	var tUrl, TText string
	var TPubDate, TFetchTime, TFirstFetchTime uint64
	tUrl = u
	TText = RandStringRunes(10)
	TPubDate = uint64(rand.Int63n(1000))
	TFetchTime = uint64(rand.Int63n(1000))
	TFirstFetchTime = 0

	doc := Document{
		Url:            tUrl,
		PubDate:        TPubDate,
		FetchTime:      TFetchTime,
		Text:           TText,
		FirstFetchTime: TFirstFetchTime,
	}
	return doc
}
