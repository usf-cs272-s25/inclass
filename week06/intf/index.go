type Index interface {
	Search(term) Results
	BuildIndex(url, stems)
}
