# Inspector

Inspector is a simple HTTP server that simple JSON encodes the incoming request and writes it out as the response, making it easy to inspect headers.

## Example Response

```json
{
  "Method": "GET",
  "URL": {
    "Scheme": "",
    "Opaque": "",
    "User": null,
    "Host": "",
    "Path": "/v2",
    "RawQuery": "",
    "Fragment": ""
  },
  "Header": {
    "Accept": [
      "*/*"
    ],
    "Client-Ip": [
      "192.168.59.3"
    ],
    "Connection": [
      "close"
    ],
    "User-Agent": [
      "curl/7.37.1"
    ],
    "X-Forwarded-For": [
      "192.168.59.3"
    ],
    "X-Real-Ip": [
      "192.168.59.3"
    ],
    "X-Request-Id": [
      "ae9uok67udhrcp2l0crg"
    ],
    "X-Request-Start": [
      "t=1435856093.982"
    ]
  },
  "Host": "192.168.59.103:8080",
  "RemoteAddr": "172.17.0.18:42107",
  "RequestURI": "/v2"
}
```

In go, the response can be encoded back into an `http.Request` with:

```go
var req http.Request
json.Unmarshal(r.Body, &req)
```
