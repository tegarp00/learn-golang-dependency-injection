package simple

type BarReposiory struct {
}

func NewBarRepository() *BarReposiory {
	return &BarReposiory{}
}

type BarService struct {
	*BarReposiory
}

func NewBarService(barRepository *BarReposiory) *BarService {
	return &BarService{BarReposiory: barRepository}
}
