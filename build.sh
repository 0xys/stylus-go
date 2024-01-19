#!/bin/bash

rm_wasi() {
    cat outputs/mainh_tmp2.wat | grep -v preview1 > tmp.wat

    LINE_FD_WRITE=$(cat tmp.wat | grep -n fd_write | cut -d ':' -f1)
    START=$(echo "$LINE_FD_WRITE-4" | bc)
    END=$(echo "$LINE_FD_WRITE+1" | bc)

    echo "$START","$END"d

    sed "$START","$END"d tmp.wat > outputs/mainh_tmp3.wat
    diff --color outputs/mainh_tmp2.wat outputs/mainh_tmp3.wat
    rm tmp.wat
}

build() {
    local mainfile=$1
    tinygo build -o outputs/mainh_tmp1.wasm -gc leaking -scheduler none -target ./configs/stylus.json --no-debug $mainfile
	$WABT_PATH/bin/wasm2wat -o outputs/mainh_tmp2.wat outputs/mainh_tmp1.wasm
	rm_wasi
	$WABT_PATH/bin/wat2wasm -o bin/mainh.wasm outputs/mainh_tmp3.wat
    echo success
    echo outfile: ./bin/mainh.wasm
}

check() {
    local binfile=$1
    cargo stylus check --wasm-file-path $binfile
}

deploy() {
    local binfile=$1
    cargo stylus deploy --endpoint http://localhost:8547 --wasm-file-path $binfile --private-key $ETH_PRIVATE_KEY
}


main() {
    local cmd=$1
    if [ -z "$cmd" ]; then
        echo "specify command type. Usage $0 <build|check|deploy>"
        return
    fi

    if [ "$cmd" = "build" ]; then
        arg=$2
        if [ -z "$arg" ]; then
            echo "specify the go program being built. Usage $0 build <path_to_go_file>"
            return
        fi
        build $arg
        return
    fi

    if [ "$cmd" = "check" ]; then
        arg=$2
        if [ -z "$arg" ]; then
            echo "specify the wasm file to check. Usage $0 check <path_to_wasm_file>"
            return
        fi
        check $arg
        return
    fi

    if [ "$cmd" = "deploy" ]; then
        arg=$2
        if [ -z "$arg" ]; then
            echo "specify the wasm file to deploy. Usage $0 deploy <path_to_wasm_file>"
            return
        fi
        deploy $arg
        return
    fi

    echo "specify command type. Usage $0 <build|check|deploy>"
}

main $1 $2
