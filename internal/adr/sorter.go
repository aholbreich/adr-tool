package adr

import (
	"sort"

	"github.com/aholbreich/adr-tool/internal/model"
)

// ByNumber implements sort.Interface for []model.
type ByNumber []model.Adr

func (a ByNumber) Len() int           { return len(a) }
func (a ByNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNumber) Less(i, j int) bool { return a[i].Number < a[j].Number } // Reversed order for descending sort

// ReverseSortAdrList sorts the slice of model.Adr in reverse order
func SortAdrList(adrs []model.Adr) {
	sort.Sort(ByNumber(adrs))
}

func SortAdrListReverse(adrs []model.Adr) {
	sort.Sort(sort.Reverse(ByNumber(adrs)))
}
