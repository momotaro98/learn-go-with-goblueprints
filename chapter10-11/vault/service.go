package vault

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"encoding/json"
)

// Service provides password hashing capabilities.
type Service interface {
	Hash(ctx context.Context, password string) (string, error)
	Validate(ctx context.Context, password, hash string) (bool, error)
}

// NewService makes a new Service.
func NewService() Service {
	return vaultService{}
}

type vaultService struct {}

func (vaultService) Hash(ctx context.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (vaultService) Validate(ctx context.Context, password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

type hashRequest struct {
	Password string `json:"password"`
}
type hashResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}
type validateRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}
type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

func decodeHashRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req hashRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeHashEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hashRequest)
		v, err := srv.Hash(ctx, req.Password)
		if err != nil {
			// An Error returned from an endpoint is intended to indicate a transport error,
			// and perhaps Go kit will be configured to retry the call a few times by some
			// middleware. If our service methods return an error, it is considered a business
			// logic error and will probably always return the same error for the same input,
			// so it's not worth retrying. This is why we wrap the error into the response
			// and return it to the client so that they can deal with it.
			return hashResponse{v, err.Error()}, nil
		}
		return hashResponse{v, ""}, nil
	}
}

func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		v, err := srv.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return validateResponse{false, err.Error()}, nil
		}
		return validateResponse{v, ""}, nil
	}
}