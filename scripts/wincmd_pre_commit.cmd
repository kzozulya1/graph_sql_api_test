cd ..

go test ./...
pause

golangci-lint.exe  run
pause