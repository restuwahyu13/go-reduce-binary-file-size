org:
	go build --race -v -o main .

ldf:
	go build --ldflags "-r" -o main .

mix:
	go build --race --ldflags "-r" -o main .
	upx --no-progress -9 ./main
	upx -t ./main

run:
	go run .