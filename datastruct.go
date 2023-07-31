package main

type Job interface {
	Go(Pipeline) error
	GetDataChan() DataChan
}

type DataChan chan interface{}

type Pipeline interface {
	StartPipe()
}

//--- data interfaces
type HasFieldValueString interface {
	Value() string
}
type HasIntID interface {
	ID() int
}
