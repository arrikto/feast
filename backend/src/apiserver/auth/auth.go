package auth

import (
	"context"

	"github.com/feast-dev/feast/backend/src/apiserver/client"
	"github.com/feast-dev/feast/backend/src/apiserver/common"
	util "github.com/feast-dev/feast/backend/src/utils"
	"github.com/pkg/errors"
)

type Authenticator interface {
	GetUserIdentity(ctx context.Context) (string, []string, error)
}

var IdentityHeaderMissingError = util.NewUnauthenticatedError(
	errors.New("Request header error: there is no user identity header."),
	"Request header error: there is no user identity header.",
)

func GetAuthenticators(tokenReviewClient client.TokenReviewInterface) []Authenticator {
	return []Authenticator{
		NewHTTPHeaderAuthenticator(common.GetKubeflowUserIDHeader(), common.GetKubeflowUserIDPrefix()),
		NewTokenReviewAuthenticator(
			common.AuthorizationBearerTokenHeader,
			common.AuthorizationBearerTokenPrefix,
			[]string{common.GetTokenReviewAudience()},
			tokenReviewClient,
		),
	}
}
