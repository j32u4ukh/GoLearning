package array

type Array interface {
	Append(v interface{})
	Contains(v interface{}) bool
	Length() int
	Find(v interface{}) int
	Remove(v interface{}) interface{}
}

func Contains[T string | int](array []T, element T) bool {
	for _, a := range array {
		if a == element {
			return true
		}
	}
	return false
}

type StringArray struct {
	Elements []string
}

func (sa *StringArray) Init() {
	sa.Elements = make([]string, 10)
}

func (sa *StringArray) Append(v interface{}) {
	if v != nil {
		sa.Elements = append(sa.Elements, v.(string))
	}
}

func (sa *StringArray) Contains(v interface{}) bool {
	idx := sa.Find(v)
	return idx != -1
}

func (sa *StringArray) Length() int {
	return len(sa.Elements)
}

func (sa *StringArray) Find(v interface{}) int {
	for i, e := range sa.Elements {
		if e == v.(string) {
			return i
		}
	}

	return -1
}

func (sa *StringArray) Remove(v interface{}) interface{} {
	idx := sa.Find(v)

	if idx == -1 {
		return -1
	}

	sa.Elements = append(sa.Elements[:idx-1], sa.Elements[idx+1:]...)

	return v
}
