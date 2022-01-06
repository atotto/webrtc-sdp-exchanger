
gen:
	protoc -I protobuf -I apis --go_out=apis --go_opt=paths=source_relative \
			--go-grpc_out=apis --go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=apis --grpc-gateway_opt=paths=source_relative \
			apis/exchange.proto

clean:
	rm -f apis/*\.pb*\.go

run:
	GOOGLE_CLOUD_PROJECT=webrtc-sdp-exchanger go run main.go

deploy:
	gcloud app deploy --project webrtc-sdp-exchanger -q
