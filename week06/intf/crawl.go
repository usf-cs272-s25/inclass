func crawl(seedUrl string, idx Index) {
    url = ...
    stems = ...

    // crawler doesn't need to know if idx represents
    // a SQL DB impl or an inverted index impl
    idx.BuildIndex(url, stems)
}
