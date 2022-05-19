package designpattern

type Operation interface {
	SetA(int)
	SetB(int)
	Result() int
}

type BaseOperation struct {
	a, b int
}

func (t *BaseOperation) SetA(a int) {
	t.a = a
}

func (t *BaseOperation) SetB(b int) {
	t.b = b
}

type AddOperation struct {
	*BaseOperation
}

func (t *AddOperation) Result() int {
	return t.a + t.b
}

type SubOperation struct {
	*BaseOperation
}

func (t *SubOperation) Result() int {
	return t.a - t.b
}

type OperationFactory interface {
	Create() Operation
}

type AddOperationFactory struct{}

func (t *AddOperationFactory) Create() Operation {
	return &AddOperation{}
}

type SubOperationFactory struct{}

func (t *SubOperationFactory) Create() Operation {
	return &SubOperation{}
}
