package processor

import ()

type Processor interface {
	Process(d *Document) (*Document, error)
}

type QueueProcessor struct {
	CurrinChannel  chan Document
	CurrOutChannel chan Document
	CurrDocument   Document
}
