# G-Openid

Go package

`
go get github.com/tanopwan/gopenid
`

Example

```go
import (
	"github.com/tanopwan/gopenid"
	"github.com/pkg/errors"
	)

// Any cache that implement Cache interface
googleService := gopenid.NewGoogleService(cache.NewService())
claims, err := googleService.TokenInfoForProd(idToken)
if err != nil {
    return nil, errors.Wrap(err, "validate id_token for prod error")
}

if err := claims.Valid(); err != nil {
    return nil, errors.Wrap(err, "claims is not valid")
}

if !claims.VerifyExpiresAt(makeTimestamp(), true) {
    return nil, errors.New("claims is expired")
}

caud := os.Getenv("GOOGLE_CLIENT_ID")
if aud, ok := (*claims)["aud"].(string); !ok || aud != caud {
    return nil, errors.New("claims' aud is invalid found " + aud)
}

if iss, ok := (*claims)["iss"].(string); !ok || (iss != ciss1 && iss != ciss2) {
    return nil, errors.New("claims' iss is invalid found " + iss)
}

email, ok := (*claims)["email"].(string)
if !ok {
    return nil, errors.New("claims' email is invalid")
}
	
```