package heap

type HeapString struct {
	S string
}

func (s1 HeapString) Compare(s2 HeapString) int {

	if s1.S == s2.S {
		return 0
	} else if s1.S > s2.S {
		return 1
	}

	return -1
}

func (s HeapString) GetMin() HeapString {
	return GetHeapString("")
}

func GetHeapString(s string) HeapString {
	return HeapString{
		S: s,
	}
}
