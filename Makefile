all:
	go generate -v
	gopherjs build -v -m
