
linux:
	GOOS=linux GOARCH=amd64 go build -o gomail-linux-x86_64

windows:
	GOOS=windows GOARCH=amd64 go build -o gomail-windows-x86_64.exe

