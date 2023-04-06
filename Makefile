VERSION=0.3.0

build_linux:
	@echo 'building linux binary...'
	env GOOS=linux GOARCH=amd64 go build -o nuru
	@echo 'zipping build....'
	tar -zcvf nuru_linux_amd64_v${VERSION}.tar.gz nuru
	@echo 'cleaning up...'
	rm nuru

build_windows:
	@echo 'building windows executable...'
	env GOOS=windows GOARCH=amd64 go build -o nuru_windows_amd64_v${VERSION}.exe

build_mac:
	@echo 'building mac binary...'
	env GOOS=darwin GOARCH=amd64 go build -o nuru
	@echo 'zipping build...'
	tar -zcvf nuru_mac_amd64_v${VERSION}.tar.gz nuru
	@echo 'cleaning up...'
	rm nuru

build_android:
	@echo 'building android binary'
	env GOOS=android GOARCH=arm64 go build -o nuru
	@echo 'zipping build...'
	tar -zcvf nuru_linux_amd64_v${VERSION}.tar.gz nuru
	@echo 'cleaning up...'
	rm nuru

build_test:
	go build -o test
	mv test testbinaries/

test:
	./gotest --format testname ./lexer/ ./parser/ ./ast/ ./object/ ./evaluator/

clean:
	go clean
