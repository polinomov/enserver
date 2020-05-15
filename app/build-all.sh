echo --- - building --- GOPATH=$GOPATH
echo --- - Installing Dependancies
apk add protoc
echo --- - building protobuffer
protoc --go_out ./ protodef/enginecmd.proto
echo --- - building main
$GOPATH/libEnserver/build.sh
echo --- -building game
$GOPATH/game/build.sh
echo --- -build_done  ---
