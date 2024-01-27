# example
```sh
./build.sh gen
./build.sh build ./example
./build.sh check ./bin/mainh.wasm
./build.sh deploy ./bin/mainh.wasm
```

# project structure
```
- entrypoint.go // generated
- contract
  + contract.go // contain storage struct and methods
```