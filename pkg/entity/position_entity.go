package entity

type Position string

var (
	N Position = "N"
	E Position = "E"
	S Position = "S"
	W Position = "W"
)

func (position *Position) IsValid() bool {
	return *position == N ||
		*position == E ||
		*position == S ||
		*position == W
}

func (position *Position) String() string {
	return string(*position)
}
