pwd
go build -o libgame.so -buildmode=c-shared pingpong.go
sudo cp libgame.so /usr/local/lib/libgame.so
sudo cp libgame.h /usr/include/libgame.h
sudo ldconfig
ls 
