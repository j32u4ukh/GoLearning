package color

type Color int

const (
	Black Color = iota
	White
	Red
)

var str = []string{
	"Black",
	"White",
	"Red",
}

func (c Color) String() string {
	return str[c]
}
