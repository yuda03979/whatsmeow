package whatsmeow

import (
	"testing"

	"go.mau.fi/whatsmeow/types"
)

func TestSkipRetryPreKey(t *testing.T) {
	group := types.JID{User: "120363000000000000", Server: types.GroupServer}
	dm := types.JID{User: "1555550100", Server: types.DefaultUserServer}
	broadcast := types.JID{User: "status", Server: types.BroadcastServer}

	cases := []struct {
		name                 string
		enabled              bool
		forceIncludeIdentity bool
		chat                 types.JID
		want                 bool
	}{
		{"group without a sender key", true, true, group, true},
		{"group retry that does not force the identity", true, false, group, false},
		{"direct message", true, true, dm, false},
		{"broadcast list", true, true, broadcast, false},
		{"disabled, group without a sender key", false, true, group, false},
		{"disabled, direct message", false, true, dm, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := skipRetryPreKey(tc.enabled, tc.forceIncludeIdentity, tc.chat); got != tc.want {
				t.Errorf("skipRetryPreKey(%t, %t, %s) = %t, want %t", tc.enabled, tc.forceIncludeIdentity, tc.chat, got, tc.want)
			}
		})
	}
}
