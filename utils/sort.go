package utils

import "sort"

func Sort(values []any, comparator Comparator) {
	sort.Sort(sortable{values, comparator})

}

type sortable struct {
	values     []any
	comparator Comparator
}

func (s sortable) Len() int {
	return len(s.values)
}

func (s sortable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s sortable) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}
