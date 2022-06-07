## How to Test 

1. Run ngrok Server 
```shell
ngrok http 8080
```
2. Copy & Paste Forwarding URL into endpoint of Webhook at Qase.

```
{URL}/webhook/qase/test

//example
https://f0d1-220-72-239-218.jp.ngrok.io/webhook/qase/test
```

3. Run Go Server with below command in terminal.
```go
go run src/main.go
```

4. Test with Postman or Qase whatever you prefer.