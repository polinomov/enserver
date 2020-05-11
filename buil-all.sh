echo --- - building --- GOPATH=$GOPATH
export LD_LIBRARY_PATH=~/Projects/Proj1/src/github.com/polinomov/enserver
cd game
echo --- - building pingpong.go
go build -o ../libgame.so -buildmode=c-shared pingpong.go aux.go
cd ..
echo --- - building protobuffer
protoc --go_out ./ protodef/enginecmd.proto
echo --- - building main
go install
echo --- build_done  ---
