package gopenid

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"time"
)

// GoogleService contains business logic to the domain
type GoogleService struct {
	cache Cache
}

// NewGoogleService return new object
func NewGoogleService(cache Cache) *GoogleService {
	return &GoogleService{
		cache: cache,
	}
}

const timeout = time.Duration(10 * time.Second)
const (
	googleDomain = "https://accounts.google.com"
	openIDURI    = "/.well-known/openid-configuration"
)

func getKeyFunction(cc Cache) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		keyID, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("expecting JWT header to have string kid")
		}

		if cc != nil {
			value := cc.Get(keyID)
			if value != nil {
				fmt.Printf("cache found google_public for key: %s and value: %+v\n", keyID, value)
				return value, nil
			}
		}

		fmt.Printf("cache not found google_public for key: %s\n", keyID)

		client := &http.Client{Timeout: timeout}

		var jwksURI string
		{
			u, err := url.ParseRequestURI(googleDomain)
			if err != nil {
				return nil, errors.Wrap(ErrInternalServerError, fmt.Sprintf("failed to parse request uri: %s", err.Error()))
			}
			u.Path = openIDURI
			urlStr := u.String()

			r, err := http.NewRequest("GET", urlStr, nil)
			if err != nil {
				return nil, errors.Wrap(ErrInternalServerError, fmt.Sprintf("failed to create request body: %s", err.Error()))
			}

			resp, err := client.Do(r)
			if err != nil {
				return nil, errors.Wrap(ErrExternalDAO, fmt.Sprintf("failed to read response: %s", err.Error()))
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Printf("failed to close response body with reason: %s\n", err.Error())
				}
			}()

			respBody := struct {
				Issuer                            string   `json:"issuer"`
				AuthorizationEndpoint             string   `json:"authorization_endpoint"`
				TokenEndpoint                     string   `json:"token_endpoint"`
				UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
				RevocationEndpoint                string   `json:"revocation_endpoint"`
				JwksURI                           string   `json:"jwks_uri"`
				ResponseTypesSupported            []string `json:"response_types_supported"`
				SubjectTypesSupported             []string `json:"subject_types_supported"`
				IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
				ScopesSupported                   []string `json:"scopes_supported"`
				TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
				ClaimsSupported                   []string `json:"claims_supported"`
				CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"`
			}{}

			err = json.NewDecoder(resp.Body).Decode(&respBody)
			if err != nil {
				return nil, errors.Wrap(ErrInternalServerError, fmt.Sprintf("failed to decode response body: %s", err.Error()))
			}

			jwksURI = respBody.JwksURI
		}

		// we want to verify a JWT
		set, err := jwk.FetchHTTP(jwksURI)
		if err != nil {
			return nil, err
		}

		if key := set.LookupKeyID(keyID); len(key) == 1 {
			m, err := key[0].Materialize()
			if err != nil {
				return nil, err
			}
			cc.SetExpire(keyID, m, time.Hour*20)
			fmt.Printf("cached keyID[%s] for %d hours\n", keyID, 20)
			return m, nil
		}

		return nil, errors.New("unable to find key (key is too old)")
	}
}

// TokenInfoForProd dao
func (s *GoogleService) TokenInfoForProd(IDToken string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(IDToken, getKeyFunction(s.cache))
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	// for key, value := range claims {
	// 	fmt.Printf("%s\t%v\n", key, value)
	// }

	return &claims, nil
}
