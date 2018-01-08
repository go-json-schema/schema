package common

func (l *EnumList) Contains(v interface{}) bool {
	for x := range l.Iterator() {
		if x == v {
			return true
		}
	}
	return false
}

func (l *EnumList) Append(list ...Enum) {
	*l = append(*l, list...)
}

func (l EnumList) Iterator() <-chan Enum {
	ch := make(chan Enum, len(l))
	go func() {
		defer close(ch)
		for _, e := range l {
			ch <- e
		}
	}()
	return ch
}
