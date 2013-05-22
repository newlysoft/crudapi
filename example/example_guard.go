package main

import (
	"github.com/sauerbraten/crudapi"
	"net/url"
)

type MapGuard struct {
	ValidActions map[string][]crudapi.Action
}

// no authentication; only valid actions are authorized
func (a MapGuard) AuthenticateAndAuthorize(action crudapi.Action, kind string, params url.Values) crudapi.GuardResponse {
	for _, validAction := range a.ValidActions[kind] {
		if validAction == action {
			return crudapi.GuardResponse{true, true, ""}
		}
	}

	return crudapi.GuardResponse{true, false, "action not allowed for this kind of resource"}
}
