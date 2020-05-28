objs: 
	go build -o x.exe defer.go
	go tool objdump -s main x.exe

debug:
	dlv debug defer.go
