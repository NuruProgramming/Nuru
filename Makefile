VERSION=0.5.1

build_linux:
	@echo 'building linux binary...'
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o nuru
	@echo 'shrinking binary...'
	./upx --brute nuru
	@echo 'zipping build....'
	tar -zcvf nuru_linux_amd64_v${VERSION}.tar.gz nuru
	@echo 'cleaning up...'
	rm nuru

build_windows:
	@echo 'building windows executable...'
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o nuru_windows_amd64_v${VERSION}.exe
	@echo 'shrinking build...'
	./upx --brute nuru_windows_amd64_v${VERSION}.exe

build_mac:
	@echo 'building mac binary...'
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o nuru
	@echo 'shrinking binary...'
	./upx --brute nuru
	@echo 'zipping build...'
	tar -zcvf nuru_mac_amd64_v${VERSION}.tar.gz nuru
	@echo 'cleaning up...'
	rm nuru

build_android:
	@echo 'building android binary'
	env GOOS=android GOARCH=arm64 go build -ldflags="-s -w" -o nuru
	@echo 'zipping build...'
	tar -zcvf nuru_android_arm64_v${VERSION}.tar.gz nuru
	@echo 'cleaning up...'
	rm nuru

build_test:
	go build -ldflags="-s -w" -o nuru

dependencies:
	@echo 'checking dependencies...'
	go mod tidy

test:
	@echo -e '\nTesting Lexer...'
	@./gotest --format testname ./lexer/ 
	@echo -e '\nTesting Parser...'
	@./gotest --format testname ./parser/
	@echo -e '\nTesting AST...'
	@./gotest --format testname ./ast/
	@echo -e '\nTesting Object...'
	@./gotest --format testname ./object/
	@echo -e '\nTesting Evaluator...'
	@./gotest --format testname ./evaluator/

clean:
	go clean
