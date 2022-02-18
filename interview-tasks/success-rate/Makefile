.PHONY: bench bench-mem bench-profile

INFO_DIR = $(PWD)/bench-info

bench:
	go test -bench=. scanner.go scanner_test.go

bench-mem:
	go test -benchmem -bench=. scanner.go scanner_test.go > $(INFO_DIR)/new.txt

bench-profile:
	go test -bench=. -benchmem -cpuprofile=$(INFO_DIR)/cpu.out -memprofile=$(INFO_DIR)/mem.out scanner.go scanner_test.go
