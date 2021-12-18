package simple

type FooRepostiory struct {
}

func NewFooRepository() *FooRepostiory {
	return &FooRepostiory{}
}

type FooService struct {
	*FooRepostiory
}

func NewFooService(fooRepostiory *FooRepostiory) *FooService {
	return &FooService{FooRepostiory: fooRepostiory}
}
