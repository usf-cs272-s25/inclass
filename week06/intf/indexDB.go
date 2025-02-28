
type IndexDB struct {
	db *DB // connection
}

New IndexDB() Index {
	return &IndexDB {
		db: //...
	}
}

func (idb IndexDB) Search() {
	// SELECT ...
}

func (idb IndexDB) BuildIndex(url, stems) {
	// INSERT ...
}
