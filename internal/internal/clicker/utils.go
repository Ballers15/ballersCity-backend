package clicker

import (
	"context"
	"fmt"
	"github.com/nostressdev/nerrors"
	"google.golang.org/grpc/metadata"
)

const AccessToken = "AccessToken"

type ContextInfo struct {
	AccountID string
}

var ErrNoMetadataString = fmt.Sprintf("no metadata")
var ErrNoAccessTokenString = fmt.Sprintf("no access token in metadata")
var ErrInvalidAppIDString = fmt.Sprintf("invalid app ID")
var ErrTokenNotAvailableString = fmt.Sprintf("token is not available")

func (s *TvxServer) getToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md, ok = metadata.FromOutgoingContext(ctx)
	}
	if !ok {
		return "", nerrors.Validation.New(ErrNoMetadataString)
	}
	accessTokens := md.Get(AccessToken)
	if accessTokens == nil || len(accessTokens) == 0 {
		return "", nerrors.Validation.New(ErrNoAccessTokenString)
	}
	accessToken := accessTokens[0]
	return accessToken, nil
}
