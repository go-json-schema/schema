package common

func (l *EnumList) Append(list ...Enum) {
	*l = append(*l, list...)
}
