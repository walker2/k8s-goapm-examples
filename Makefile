all:
	build-grpc
	build-mux-span

build-grpc:
	make -C grpc

build-mux-span:
	make -C mux-span