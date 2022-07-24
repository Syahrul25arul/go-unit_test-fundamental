# go test = untuk menjalankan semua unit test

# go test -v = untuk melihat lebih detail function test apa saja yang di running

# go test -v -run TestNamaFunction = untuk menjalankan function test tertentu

# go test ./... = untuk menjalankan semua unit test di package sekarang dan package lain yang ada didalamnya

# go test -run TestNamaFunction/NamaSubTest = untuk menjalankan salah satu sub test

# go test -run /NamaSubTest = untuk menjalankan semua sub test

# go test -v -bench=. = untuk menjalankan seluruh benchmark dan unit test di sebuah package

# go test -v -run=NamaFunctionTestYangTidakDigunakan -bench=. = untuk menjalankan seluruh benchmark di sebuah package tanpa menjalankan unit test

# go test -v -run=NamaFunctionTestYangTidakDigunakan -bench=BenchmarkName = untuk menjalankan benchmark tertentu

# go test -v -bench=. ./... = untuk menjalankan semua unit test dan benchamrk

# go test -v -run=TestFunctionTidakAda -bench=. ./... = untuk menjalankan semua benchmark di semua package tanpa menjalankan unit test
