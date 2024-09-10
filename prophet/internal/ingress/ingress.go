package ingress

type Future struct {
	ID      uint64
	Handler string
	Input   []byte
}

func (r Future) GetID() uint64 { return r.ID }

type FutureSource interface {
	Fetch() <-chan Future
}

func Futures(s FutureSource) (<-chan Future, error) {
	reqs, err := dedup(s.Fetch())
	if err != nil {
		return nil, err
	}

	return reqs, nil
}

type FutureResult struct {
	Future
	Output []byte
}

func (r FutureResult) GetID() uint64 { return r.ID }

type FutureResultSource interface {
	Fetch() <-chan FutureResult
}

func FutureResults(s FutureResultSource) (<-chan FutureResult, error) {
	reqs, err := dedup(s.Fetch())
	if err != nil {
		return nil, err
	}

	return reqs, nil
}
