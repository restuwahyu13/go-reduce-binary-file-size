# Golang Reduce Binary File Size

+ build original binary file size **9.3 MB**
+ build using -ldflags binary file size **7.3 MB**
+ build using -ldflags + upx binary file size **4.3 MB**

## Command

```makefile
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
```
