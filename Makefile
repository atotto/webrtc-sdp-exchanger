
gen:
	protoc -I protobuf -I apis --go_out=plugins=grpc:apis --grpc-gateway_out=apis apis/exchange.proto

deploy:
	gcloud app deploy --project webrtc-sdp-exchanger
