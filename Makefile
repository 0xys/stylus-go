
WABT_PATH := $(WABT_PATH)

.PHONY: build
build:
	tinygo build -o bin/wasm.wasm -target wasm ./main.go
	tinygo build -o bin/wasm_leaking.wasm -gc leaking -target wasm ./main.go
	tinygo build -o bin/wasm_cons_nosch.wasm -gc leaking -scheduler none -target wasm ./main.go
	tinygo build -o bin/wasm_cons_nosch_1.wasm -gc leaking -scheduler none -opt 1 -target wasm ./main.go
	tinygo build -o bin/wasm_cons_nosch_2.wasm -gc leaking -scheduler none -opt 2 -target wasm ./main.go
	tinygo build -o bin/main.wasm -gc leaking -scheduler none -target wasm --no-debug ./main.go

.PHONY: build/hacky
build/hacky:
	tinygo build -o bin/mainh_tmp.wasm -gc leaking -scheduler none -target stylus.json --no-debug ./main.go
	$(WABT_PATH)/bin/wasm2wat -o outputs/mainh.wat bin/mainh_tmp.wasm
	$(WABT_PATH)/bin/wat2wasm -o bin/mainh.wasm outputs/mainh.wat
	
.PHONY: check
check:
	cargo stylus check --wasm-file-path bin/main.wasm

.PHONY: check/h
check/h:
	cargo stylus check --wasm-file-path bin/mainh.wasm

.PHONY: wasm2wat
wasm2wat:
	$(WABT_PATH)/bin/wasm2wat -o outputs/main.wat bin/main.wasm

.PHONY: objdump/x
objdump/x:
	$(WABT_PATH)/bin/wasm-objdump bin/main.wasm -x

.PHONY: objdump
objdump:
	$(WABT_PATH)/bin/wasm-objdump bin/main.wasm -h

.PHONY: opt
opt:
	wasm-opt

