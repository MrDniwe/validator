package main

type LenSortRow struct {
	Id    int
	Value string
}

type lenSortJob struct {
	p    Pipeline
	deps []string
}

func (j *lenSortJob) Go(p Pipeline) error {
	j.p = p
	return nil
}

func (j *lenSortJob) GetDataChan() DataChan {
	c := make(DataChan)
	return c
}

func NewIntSorterJob() Job {
	return &lenSortJob{
		deps: []string{"source"},
	}
}
