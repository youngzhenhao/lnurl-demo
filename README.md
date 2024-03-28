# lnurl-demo

**This demo hasn't been fully tested yet.**

## Design

![LNURL](./LNURL.jpg)

## Steps

1. Alice runs service on Phone

```go
api.RouterRunOnPhone()
```

2. Server runs service 

```go
api.RouterRunOnServer()
```

3. Alice uploads info to Server

```go
api.UploadUserInfo(name, socket) 
```

4. Alice gets LNURL, (Alice generates QR code, Bob scans QR code)

*Front-end implementation*

5. Bob uses LNURL to pay amount of stas

```go
api.PayToLnurl(lnu, amount)
```

6. Bob gets response of transaction hash