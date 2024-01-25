# example
```sh
./build.sh build ./templates
./build.sh check ./bin/mainh.wasm
./build.sh deploy ./bin/mainh.wasm
```

# project structure
```
- entrypoint.go // generated
- contract
  + contract.go // contain storage struct and methods
```