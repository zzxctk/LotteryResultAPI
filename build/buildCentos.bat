set GOPATH=%GOPATH%;J:\Svn_Code\bnx
set GOARCH=amd64
set GOOS=linux
go build ..\main.go
del LotteryResultAPI
rename main LotteryResultAPI
go build -gcflags "all=-N -l" -o .\csMain ..\csMain.go