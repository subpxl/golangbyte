proto:
	rm -f pb/*
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto


rpc:
	evans --host localhost --port 5000 -r repl



.PHONY: proto