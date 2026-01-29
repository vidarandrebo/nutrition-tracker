package meal

type Entry[T any] struct {
	Definition     T
	Amount         float64
	SequenceNumber int
}
