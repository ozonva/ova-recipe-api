package recipe

type Action interface {
	String() string
	DoAction() error
}
