package blesser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"veyron.io/veyron/veyron/services/identity"
	"veyron.io/veyron/veyron/services/identity/revocation"

	"veyron.io/veyron/veyron2"
	"veyron.io/veyron/veyron2/ipc"
	"veyron.io/veyron/veyron2/security"
	"veyron.io/veyron/veyron2/vdl/vdlutil"
	"veyron.io/veyron/veyron2/vlog"
)

type googleOAuth struct {
	rt                 veyron2.Runtime
	authcodeClient     struct{ ID, Secret string }
	accessTokenClients []struct{ ID string }
	duration           time.Duration
	domain             string
	dischargerLocation string
	revocationManager  *revocation.RevocationManager
}

// GoogleParams represents all the parameters provided to NewGoogleOAuthBlesserServer
type GoogleParams struct {
	// The Veyron runtime to use
	R veyron2.Runtime
	// The OAuth client IDs for the clients of the BlessUsingAccessToken RPCs.
	AccessTokenClients []struct {
		ID string
	}
	// The duration for which blessings will be valid.
	BlessingDuration time.Duration
	// If non-empty, only email addresses from this domain will be blessed.
	DomainRestriction string
	// The object name of the discharger service. If this is empty then revocation caveats will not be granted.
	DischargerLocation string
	// The revocation manager that generates caveats and manages revocation.
	RevocationManager *revocation.RevocationManager
}

// NewGoogleOAuthBlesserServer provides an identity.OAuthBlesserService that uses authorization
// codes to obtain the username of a client and provide blessings with that name.
//
// For more details, see documentation on Google OAuth 2.0 flows:
// https://developers.google.com/accounts/docs/OAuth2
//
// Blessings generated by this server expire after duration. If domain is non-empty, then blessings
// are generated only for email addresses from that domain.
func NewGoogleOAuthBlesserServer(p GoogleParams) interface{} {
	return identity.NewServerOAuthBlesser(&googleOAuth{
		rt:                 p.R,
		duration:           p.BlessingDuration,
		domain:             p.DomainRestriction,
		dischargerLocation: p.DischargerLocation,
		revocationManager:  p.RevocationManager,
		accessTokenClients: p.AccessTokenClients,
	})
}

func (b *googleOAuth) BlessUsingAccessToken(ctx ipc.ServerContext, accesstoken string) (vdlutil.Any, error) {
	if len(b.accessTokenClients) == 0 {
		return nil, fmt.Errorf("server not configured for blessing based on access tokens")
	}
	// URL from: https://developers.google.com/accounts/docs/OAuth2UserAgent#validatetoken
	tokeninfo, err := http.Get("https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=" + accesstoken)
	if err != nil {
		return nil, fmt.Errorf("unable to use token: %v", err)
	}
	if tokeninfo.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to verify access token: %v", tokeninfo.StatusCode)
	}
	// tokeninfo contains a JSON-encoded struct
	var token struct {
		IssuedTo      string `json:"issued_to"`
		Audience      string `json:"audience"`
		UserID        string `json:"user_id"`
		Scope         string `json:"scope"`
		ExpiresIn     int64  `json:"expires_in"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		AccessType    string `json:"access_type"`
	}
	if err := json.NewDecoder(tokeninfo.Body).Decode(&token); err != nil {
		return "", fmt.Errorf("invalid JSON response from Google's tokeninfo API: %v", err)
	}
	audienceMatch := false
	for _, c := range b.accessTokenClients {
		if token.Audience == c.ID {
			audienceMatch = true
			break
		}
	}
	if !audienceMatch {
		vlog.Infof("Got access token [%+v], wanted one of client ids %v", token, b.accessTokenClients)
		return "", fmt.Errorf("token not meant for this purpose, confused deputy? https://developers.google.com/accounts/docs/OAuth2UserAgent#validatetoken")
	}
	if !token.VerifiedEmail {
		return nil, fmt.Errorf("email not verified")
	}
	return b.bless(ctx, token.Email)
}

func (b *googleOAuth) bless(ctx ipc.ServerContext, name string) (vdlutil.Any, error) {
	if len(b.domain) > 0 && !strings.HasSuffix(name, "@"+b.domain) {
		return nil, fmt.Errorf("blessings for name %q are not allowed due to domain restriction", name)
	}
	self := b.rt.Identity()
	var err error
	// Use the blessing that was used to authenticate with the client to bless it.
	if self, err = self.Derive(ctx.LocalID()); err != nil {
		return nil, err
	}
	var revocationCaveat security.ThirdPartyCaveat
	if b.revocationManager != nil {
		revocationCaveat, err = b.revocationManager.NewCaveat(b.rt.Identity().PublicID(), b.dischargerLocation)
		if err != nil {
			return nil, err
		}
	}

	return revocation.Bless(self, ctx.RemoteID(), name, b.duration, nil, revocationCaveat)
}
