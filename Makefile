
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
	tinygo build -o outputs/mainh_tmp1.wasm -gc leaking -scheduler none -target stylus.json --no-debug ./main.go
	$(WABT_PATH)/bin/wasm2wat -o outputs/mainh_tmp2.wat outputs/mainh_tmp1.wasm
	./rm_wasi.sh
	$(WABT_PATH)/bin/wat2wasm -o bin/mainh.wasm outputs/mainh_tmp3.wat

.PHONY: wat2wasm
wat2wasm:
	$(WABT_PATH)/bin/wat2wasm -o bin/mainh.wasm outputs/mainh.wat

.PHONY: check
check:
	cargo stylus check --wasm-file-path bin/main.wasm

.PHONY: check/hacky
check/hacky:
	cargo stylus check --wasm-file-path bin/mainh.wasm

.PHONY: deploy
deploy:
	cargo stylus deploy --wasm-file-path bin/mainh.wasm --private-key $(ETH_PRIVATE_KEY)

.PHONY: deploy/local
deploy/local:
	cargo stylus deploy --endpoint http://localhost:8547 --wasm-file-path bin/mainh.wasm --private-key $(ETH_PRIVATE_KEY)

.PHONY: wasm2wat
wasm2wat:
	$(WABT_PATH)/bin/wasm2wat -o outputs/main.wat bin/main.wasm

.PHONY: objdump/x
objdump/x:
	$(WABT_PATH)/bin/wasm-objdump bin/main.wasm -x

.PHONY: objdump
objdump:
	$(WABT_PATH)/bin/wasm-objdump bin/main.wasm -h

.PHONY: objdump/import
objdump/import:
	$(WABT_PATH)/bin/wasm-objdump --section Import -x bin/mainh.wasm

.PHONY: opt
opt:
	wasm-opt

