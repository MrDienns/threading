package input

type Input interface {
	Read() ([]string, error)
}
