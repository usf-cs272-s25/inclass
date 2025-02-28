

type IndexInv struct {
	// map of maps
}

func IndexInv() Index {
	return &IndexInv {
		//map stuff
	}
}

func (inv IndexInv) Search() {
	// map[term]...
}

func (inv IndexInv) BuildIndex(url, stems) {
	// map[term] Freq ++

	freq := m["romeo"]
	freq["chap1.html"]++
}
