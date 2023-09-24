env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o login.exe main.go
zip login.zip login.exe manifest/config/config.yaml manifest/db/user.db