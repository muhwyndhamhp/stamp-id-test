build:
	env GOOS=linux GOARCH=amd64 go build -o ./output/linux/
	env GOOS=windows GOARCH=amd64 go build -o ./output/windows/
	env GOOS=darwin GOARCH=amd64 go build -o ./output/mac/

dep:
	go mod tidy
	GO111MODULE=on go mod vendor
