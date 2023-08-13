package heap

type HeapString struct {
	S     string
	index int
}

func (s1 *HeapString) Compare(s2 *HeapString) int {

	if s1.S == s2.S {
		return 0
	} else if s1.S > s2.S {
		return 1
	}

	return -1
}

func (s *HeapString) GetMin() *HeapString {
	return GetHeapString("")
}

func (s *HeapString) GetIndex() int {
	return s.index
}

func (s *HeapString) SetIndex(index int) {
	s.index = index
}

func GetHeapString(s string) *HeapString {
	return &HeapString{
		S:     s,
		index: -1,
	}
}

func (s *HeapString) String() string {
	return s.S
}
