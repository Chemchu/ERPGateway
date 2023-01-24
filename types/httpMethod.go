package types

type HttpMethod int

const (
	GET HttpMethod = iota + 1
	POST
	PUT
	DELETE
)

// String - Creating common behavior - give the type a String function
func (d HttpMethod) String() string {
	return [...]string{"GET", "POST", "PUT", "DELETE"}[d-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (d HttpMethod) EnumIndex() int {
	return int(d)
}
