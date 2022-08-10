package auth

import (
	"context"
)

type HTTPHeaderAuthenticator struct {
	userIDHeader string
	userIDPrefix string
}

func NewHTTPHeaderAuthenticator(header, prefix string) *HTTPHeaderAuthenticator {
	return &HTTPHeaderAuthenticator{userIDHeader: header, userIDPrefix: prefix}
}

func (ha *HTTPHeaderAuthenticator) GetUserIdentity(ctx context.Context) (string, []string, error) {
	userID, err := singlePrefixedHeaderFromMetadata(ctx, ha.userIDHeader, ha.userIDPrefix)
	if err != nil {
		return "", make([]string, 0), err
	}
	return userID, make([]string, 0), nil
}
