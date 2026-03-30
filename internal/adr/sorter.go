package adr

import (
	"sort"

	"github.com/aholbreich/adr-tool/internal/model"
)

// ByNumber implements sort.Interface for []model.ADR.
type ByNumber []model.ADR

func (a ByNumber) Len() int      { return len(a) }
func (a ByNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByNumber) Less(i, j int) bool {
	return a[i].Number < a[j].Number
}

// SortADRList sorts the slice of model.ADR in ascending order.
func SortADRList(adrs []model.ADR) {
	sort.Sort(ByNumber(adrs))
}

// SortADRListReverse sorts the slice of model.ADR in descending order.
func SortADRListReverse(adrs []model.ADR) {
	sort.Sort(sort.Reverse(ByNumber(adrs)))
}
