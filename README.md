# WebRTC Session Description exchange service

(WIP)

create offer:

```
curl -X POST https://webrtc-sdp-exchanger.appspot.com/sessions/test '{"session_description":{"type":"offer","sdp":"v=0\r\no..."}}'
```

get offer:

```
curl https://webrtc-sdp-exchanger.appspot.com/sessions/test/offer
{"session_description":{"type":"offer","sdp":"v=0\r\no..."}}
```

create answer:

```
curl -X POST https://webrtc-sdp-exchanger.appspot.com/sessions/test '{"session_description":{"type":"answer","sdp":"v=0\r\no..."}}'
```

get answer:

```
curl https://webrtc-sdp-exchanger.appspot.com/sessions/test/answer
{"session_description":{"type":"answer","sdp":"v=0\r\no..."}}
```

