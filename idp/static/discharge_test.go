// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package static_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-macaroon-bakery/macaroon-bakery/v3/httpbakery"

	"github.com/kian99/candid/idp"
	"github.com/kian99/candid/idp/static"
	"github.com/kian99/candid/internal/candidtest"
	"github.com/kian99/candid/internal/discharger"
	"github.com/kian99/candid/internal/identity"
)

func TestInteractiveDischarge(t *testing.T) {
	c := qt.New(t)
	defer c.Done()

	store := candidtest.NewStore()
	sp := store.ServerParams()
	sp.IdentityProviders = []idp.IdentityProvider{
		static.NewIdentityProvider(getSampleParams()),
	}
	candid := candidtest.NewServer(c, sp, map[string]identity.NewAPIHandlerFunc{
		"discharger": discharger.NewAPIHandler,
	})
	dischargeCreator := candidtest.NewDischargeCreator(candid)
	dischargeCreator.AssertDischarge(c, httpbakery.WebBrowserInteractor{
		OpenWebBrowser: candidtest.PasswordLogin(c, "user1", "pass1"),
	})
}
