pushd plugin1
    go mod tidy
    go build -buildmode=plugin -o plugin1.so plugin1.go
popd
pushd plugin2
    go mod tidy
    go build -buildmode=plugin -o plugin2.so plugin2.go
popd

go run main.go
