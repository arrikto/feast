package auth

import (
	"context"

	"github.com/feast-dev/feast/backend/src/apiserver/client"
	util "github.com/feast-dev/feast/backend/src/utils"
	"github.com/pkg/errors"
	authv1 "k8s.io/api/authentication/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TokenReviewAuthenticator struct {
	// tokenHeader in which the authenticator expects to find the ServiceAccountToken
	tokenHeader string
	// tokenPrefix is the prefix encountered before the token
	tokenPrefix string
	// audiences the authenticator identifies as
	audiences []string
	// client to use to do TokenReviews
	client client.TokenReviewInterface
}

func NewTokenReviewAuthenticator(tokenHeader, tokenPrefix string, audiences []string, tokenReviewClient client.TokenReviewInterface) *TokenReviewAuthenticator {
	return &TokenReviewAuthenticator{
		tokenHeader: tokenHeader,
		tokenPrefix: tokenPrefix,
		audiences:   audiences,
		client:      tokenReviewClient,
	}
}

func (tra *TokenReviewAuthenticator) GetUserIdentity(ctx context.Context) (string, []string, error) {
	token, err := singlePrefixedHeaderFromMetadata(ctx, tra.tokenHeader, tra.tokenPrefix)
	if err != nil {
		return "", make([]string, 0), err
	}

	userInfo, err := tra.doTokenReview(ctx, token)
	if err != nil {
		return "", make([]string, 0), util.Wrap(err, "Authentication failure")
	}

	return userInfo.Username, userInfo.Groups, err
}

// ensureAudience makes sure all audience of the authenticator is found in the provided audience list
func (tra *TokenReviewAuthenticator) ensureAudience(audience []string) bool {
	// Create a set (map) to check fast whether something is part of the list
	audienceSet := make(map[string]struct{}, len(audience))
	for _, a := range audience {
		audienceSet[a] = struct{}{}
	}

	// Iterate through the audiences of the authenticator and check if they are part of the provided list
	for _, a := range tra.audiences {
		if _, ok := audienceSet[a]; !ok {
			return false
		}
	}
	return true
}

func (tra *TokenReviewAuthenticator) doTokenReview(ctx context.Context, userIdentity string) (*authv1.UserInfo, error) {
	review, err := tra.client.Create(
		ctx,
		&authv1.TokenReview{
			Spec: authv1.TokenReviewSpec{
				Token:     userIdentity,
				Audiences: tra.audiences,
			},
		},
		v1.CreateOptions{},
	)
	if err != nil {
		return nil, util.NewUnauthenticatedError(err, "Request header error: Failed to review the token provided")
	}

	if !review.Status.Authenticated {
		return nil, util.NewUnauthenticatedError(
			errors.New("Failed to authenticate token review"),
			"Review.Status.Authenticated is false",
		)
	}
	if !tra.ensureAudience(review.Status.Audiences) {
		return nil, util.NewUnauthenticatedError(
			errors.New("Failed to authenticate token review"),
			"Failed to find all of '%v' in audience: %v",
			tra.audiences,
			review.Status.Audiences,
		)
	}

	return &review.Status.User, nil
}
