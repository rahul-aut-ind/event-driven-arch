GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ./out/bootstrap main.go
cd ./out
chmod -R 755 bootstrap
zip function.zip bootstrap