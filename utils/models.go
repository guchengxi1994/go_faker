package utils

type FloatList []float64

type TimeStruct struct {
	From     string
	To       string
	Duration int //months
}

func (fl FloatList) Len() int {
	return len(fl)
}

func (fl FloatList) Swap(i, j int) {
	fl[i], fl[j] = fl[j], fl[i]
}

func (fl FloatList) Less(i, j int) bool {
	return fl[i] < fl[j]
}
