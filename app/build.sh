echo --- - building --- GOPATH=$GOPATH
cat /etc/ld.so.conf.d/*.conf  
BUILD_DIR=$PWD
echo --- - building protobuffer
mkdir -p $GOPATH/src/github.com/polinomov/enserver/enbuffer/cmd 
protoc -I=/home/protodef --go_out=$GOPATH /home/protodef/*.proto

echo --- -building game
cd $GOPATH/src/game
./build.sh
cd $BUILD_DIR
echo --- - building main
cd  $GOPATH/src/libEnserver
./build.sh
cd $BUILD_DIR

