module github.com/atotto/webrtc-sdp-exchanger

go 1.12

replace github.com/atotto/webrtc-sdp-exchanger => ./

require (
	cloud.google.com/go/firestore v1.2.0
	firebase.google.com/go v3.8.1+incompatible
	github.com/akutz/memconn v0.1.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.2
	github.com/rs/cors v1.6.0
	go.opencensus.io v0.22.4
	golang.org/x/net v0.0.0-20220105145211-5b0dc2dfae98 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.43.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.27.1
)
