all:
	make -C protobuf
	sudo docker build -f Dockerfile -t localhost:5000/go-apm-grpc-test:v1.0.0 .
	sudo docker push localhost:5000/go-apm-grpc-test:v1.0.0