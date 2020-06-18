module github.com/atotto/webrtc-sdp-exchanger

go 1.12

replace github.com/atotto/webrtc-sdp-exchanger => ./

require (
	cloud.google.com/go v0.58.0 // indirect
	cloud.google.com/go/firestore v1.2.0
	firebase.google.com/go v3.8.1+incompatible
	github.com/akutz/memconn v0.1.0
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/rs/cors v1.6.0
	go.opencensus.io v0.22.3
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200610111108-226ff32320da // indirect
	golang.org/x/tools v0.0.0-20200612220849-54c614fe050c // indirect
	google.golang.org/genproto v0.0.0-20200612171551-7676ae05be11
	google.golang.org/grpc v1.29.1
)
