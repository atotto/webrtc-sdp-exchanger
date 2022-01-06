# WebRTC Session Description exchange service

Example WebRTC project is here: https://github.com/atotto/mibot/tree/master/webrtc-connector

The following is an example of an exchange using curl.

create offer:

```
curl https://webrtc-sdp-exchanger.appspot.com/sessions/test -d '{"session_description":{"type":"offer","sdp":"v=0\r\no..."}}'
```

get offer:

```
curl https://webrtc-sdp-exchanger.appspot.com/sessions/test/offer
{"session_description":{"type":"offer","sdp":"v=0\r\no..."}}
```

create answer:

```
curl https://webrtc-sdp-exchanger.appspot.com/sessions/test -d '{"session_description":{"type":"answer","sdp":"v=0\r\no..."}}'
```

get answer:

```
curl https://webrtc-sdp-exchanger.appspot.com/sessions/test/answer
{"session_description":{"type":"answer","sdp":"v=0\r\no..."}}
```

