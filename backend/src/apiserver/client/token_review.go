package client

import (
	"context"
	"time"

	"github.com/cenkalti/backoff"
	util "github.com/feast-dev/feast/backend/src/utils"
	"github.com/golang/glog"
	authv1 "k8s.io/api/authentication/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TokenReviewInterface interface {
	Create(ctx context.Context, tokenReview *authv1.TokenReview, opts v1.CreateOptions) (result *authv1.TokenReview, err error)
}

func createTokenReviewClient(clientParams util.ClientParameters) (TokenReviewInterface, error) {
	clientSet, err := getKubernetesClientset(clientParams)
	if err != nil {
		return nil, err
	}
	return clientSet.AuthenticationV1().TokenReviews(), nil
}

// CreateTokenReviewClientOrFatal creates a new TokenReview client.
func CreateTokenReviewClientOrFatal(initConnectionTimeout time.Duration, clientParams util.ClientParameters) TokenReviewInterface {
	var client TokenReviewInterface
	var err error
	var operation = func() error {
		client, err = createTokenReviewClient(clientParams)
		return err
	}
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = initConnectionTimeout
	err = backoff.Retry(operation, b)

	if err != nil {
		glog.Fatalf("Failed to create TokenReview client. Error: %v", err)
	}
	return client
}
