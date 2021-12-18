
LINUX_OUT = build/gomail-linux-x86_64
WINDOWS_OUT = build/gomail-windows-x86_64.exe

linux:
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_OUT)

windows:
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_OUT)

clean:
	rm -f $(LINUX_OUT) $(WINDOWS_OUT)

