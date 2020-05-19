echo --- - building --- GOPATH=$GOPATH
# cat /etc/ld.so.conf.d/*.conf  
BUILD_DIR=$PWD
echo --- - building protobuffer
mkdir -p $GOPATH/src/github.com/polinomov/enserver/enbuffer/cmd 
protoc -I=$BUILD_DIR/protodef/ --go_out=$GOPATH $BUILD_DIR/protodef/*.proto

echo --- -building game
cd $BUILD_DIR/game
./build.sh
cd $BUILD_DIR
echo --- - building main
cd  $BUILD_DIR/libEnserver
./build.sh
cd $BUILD_DIR