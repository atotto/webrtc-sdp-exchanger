module github.com/atotto/webrtc-sdp-exchanger

go 1.12

replace github.com/atotto/webrtc-sdp-exchanger => ./

require (
	cloud.google.com/go v0.38.0
	firebase.google.com/go v3.8.1+incompatible
	github.com/akutz/memconn v0.1.0
	github.com/golang/protobuf v1.3.1
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	go.opencensus.io v0.22.0
	google.golang.org/api v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20190502173448-54afdca5d873
	google.golang.org/grpc v1.22.0
)
