package hasher

type Hasher interface {
	Compare(hashed, original string) bool
}
