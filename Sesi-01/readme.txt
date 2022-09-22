go version

echo $GOPATH

===========
mkdir hello
cd hello

go mod init hello > create module
touch main.go > create go file
go build -o main.exe > create a executable