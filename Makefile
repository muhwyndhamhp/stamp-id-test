build:
	env GOOS=linux GOARCH=amd64 go build -o ./output/linux/
	env GOOS=windows GOARCH=amd64 go build -o ./output/windows/
	env GOOS=darwin GOARCH=amd64 go build -o ./output/mac/

dep:
	go mod tidy
	GO111MODULE=on go mod vendor

run-windows:
	./output/windows/stamps-id.exe

run-mac:
	./output/mac/stamps-id

run-linux:
	./output/linux/stamps-id
