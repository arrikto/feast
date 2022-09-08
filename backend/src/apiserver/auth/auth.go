// Copyright 2022 Arrikto Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
