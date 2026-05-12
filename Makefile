
win:
	@GOOS=windows GOARCH=amd64 go build -o build/windows/Xion.exe
	@build/windows/Xion.exe

lin:
	@GOOS=linux GOARCH=amd64 go build -o build/linux/Xion
	@build/linux/Xion

clean:
	@rm -rf build/
	@echo "\033[32mBuild Cleaned\033[0m"