package pagination

const (
	DefaultLimit = 10
	MaxLimit     = 50
)

type Params struct {
	LastId uint32 `json:"last_id"`
	Limit  uint64 `json:"limit"`
}

type List[T any] struct {
	Result  []T  `json:"result"`
	HasNext bool `json:"has_next"`
}

func NewList[T any](objList []T, limit uint64) List[T] {
	return List[T]{
		Result:  objList,
		HasNext: uint64(len(objList)) == limit+1,
	}
}
