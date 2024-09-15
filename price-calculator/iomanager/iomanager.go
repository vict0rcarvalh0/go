package iomanager

type IOManager interface {
	ReadLinesFromFile() ([]string, error)
	WriteJson(data interface{}) error
}