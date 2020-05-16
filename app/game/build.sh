pwd
go build -o libgame.so -buildmode=c-shared pingpong.go
cp libgame.so /usr/local/lib/libgame.so
cp libgame.h /usr/include/libgame.h
ldconfig
ls 
