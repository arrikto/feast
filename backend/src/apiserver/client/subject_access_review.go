package client

import (
	"context"
	"time"

	"github.com/cenkalti/backoff"
	util "github.com/feast-dev/feast/backend/src/utils"
	"github.com/golang/glog"
	authzv1 "k8s.io/api/authorization/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SubjectAccessReviewInterface interface {
	Create(ctx context.Context, sar *authzv1.SubjectAccessReview, opts v1.CreateOptions) (result *authzv1.SubjectAccessReview, err error)
}

func createSubjectAccessReviewClient(clientParams util.ClientParameters) (SubjectAccessReviewInterface, error) {
	clientSet, err := getKubernetesClientset(clientParams)
	if err != nil {
		return nil, err
	}
	return clientSet.AuthorizationV1().SubjectAccessReviews(), nil
}

// CreateSubjectAccessReviewClientOrFatal creates a new SubjectAccessReview client.
func CreateSubjectAccessReviewClientOrFatal(initConnectionTimeout time.Duration, clientParams util.ClientParameters) SubjectAccessReviewInterface {
	var client SubjectAccessReviewInterface
	var err error
	var operation = func() error {
		client, err = createSubjectAccessReviewClient(clientParams)
		return err
	}
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = initConnectionTimeout
	err = backoff.Retry(operation, b)

	if err != nil {
		glog.Fatalf("Failed to create SubjectAccessReview client. Error: %v", err)
	}
	return client
}
