module github.com/atotto/webrtc-sdp-exchanger

go 1.12

replace github.com/atotto/webrtc-sdp-exchanger => ./

require (
	cloud.google.com/go v0.58.0 // indirect
	cloud.google.com/go/firestore v1.2.0
	firebase.google.com/go v3.8.1+incompatible
	github.com/akutz/memconn v0.1.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rs/cors v1.6.0
	go.opencensus.io v0.22.3
	golang.org/x/net v0.0.0-20220105145211-5b0dc2dfae98 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.0.0-20200612220849-54c614fe050c // indirect
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.43.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.27.1
)
