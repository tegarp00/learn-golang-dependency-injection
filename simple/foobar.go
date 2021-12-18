package simple

type FoorBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(fooService *FooService, barService *BarService) *FoorBarService {
	return &FoorBarService{FooService: fooService, BarService: barService}
}
