package domain

type Scraper interface {
	GetNewItem() (code int, data []*House, err error)
	GetSourceName() (name string)
}
