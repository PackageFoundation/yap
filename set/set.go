package set

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Iter() <-chan interface{} {
	iter := make(chan interface{})
	go func() {
		for key := range s.m {
			iter <- key
		}
		close(iter)
	}()
	return iter
}
