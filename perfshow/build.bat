SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go

copy d:\kuaipan\go\code\src\perfshow\main d:\kuaipan\python\mobiledata\download\perfshow\ /Y