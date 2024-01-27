# example
```sh
./build.sh gen ./example example.com/foo
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