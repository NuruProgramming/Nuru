VERSION=0.2.0

build_linux:
	env GOOS=linux GOARCH=amd64 go build -o nuru
	tar -zcvf nuru_linux_amd64_v${VERSION}.tar.gz nuru
	rm nuru

build_windows:
	env GOOS=windows GOARCH=amd64 go build -o nuru_windows_amd64_v${VERSION}.exe

build_android:
	env GOOS=android GOARCH=arm64 go build -o nuru
	tar -zcvf nuru_linux_amd64_v${VERSION}.tar.gz nuru
	rm nuru

build_test:
	go build -o test
	mv test testbinaries/

test:
	go test ./parser/
	go test ./ast/
	go test ./evaluator/
	go test ./object/

clean:
	go clean