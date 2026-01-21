.PHONY: test program.go

test: program.go
	cd runner && go run program.go

