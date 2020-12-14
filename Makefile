linux:
	env GOOS=linux GOARCH=amd64 go build -i -v -o json2hcl -ldflags="-s -w"  ./main.go
mac:
	env GOOS=darwin GOARCH=amd64 go build -i -v -o json2hcl -ldflags="-s -w"  ./main.go
windows:
	env GOOS=windows GOARCH=amd64 go build -i -v -o json2hcl.exe -ldflags="-s -w"  ./main.go
clean:
	-rm json2hcl
	-rm json2hcl.exe
