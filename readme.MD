# G-Openid

Go package

`
go get github.com/tanopwan/gopenid
`

Example

```go
// Any cache that implement Cache interface
googleService := gopenid.NewGoogleService(cache.NewService())
claims, err := googleService.TokenInfoForProd(idToken)
if err != nil {
    return nil, fmt.Errorf("validate id_token for prod error: %w", err)
}

if err := claims.Valid(); err != nil {
    return nil, fmt.Errorf("claims is not valid")
}

if !claims.VerifyExpiresAt(makeTimestamp(), true) {
    return nil, fmt.Errorf("claims is expired")
}

caud := os.Getenv("GOOGLE_CLIENT_ID")
if aud, ok := (*claims)["aud"].(string); !ok || aud != caud {
    return nil, errors.New("claims' aud is invalid found %s", aud)
}

if iss, ok := (*claims)["iss"].(string); !ok || (iss != ciss1 && iss != ciss2) {
    return nil, errors.New("claims' iss is invalid found %s", iss)
}

email, ok := (*claims)["email"].(string)
if !ok {
    return nil, errors.New("claims' email is invalid")
}
	
```