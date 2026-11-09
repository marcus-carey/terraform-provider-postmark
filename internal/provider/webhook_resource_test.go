package provider

import "testing"

func TestParseWebhookImportID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		serverID  int64
		webhookID int
		wantErr   bool
	}{
		{name: "valid", id: "12345:18121569", serverID: 12345, webhookID: 18121569},
		{name: "webhook id only", id: "18121569", wantErr: true},
		{name: "too many parts", id: "1:2:3", wantErr: true},
		{name: "non-numeric server", id: "walm-qa:18121569", wantErr: true},
		{name: "non-numeric webhook", id: "12345:abc", wantErr: true},
		{name: "zero server", id: "0:18121569", wantErr: true},
		{name: "negative webhook", id: "12345:-1", wantErr: true},
		{name: "empty", id: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serverID, webhookID, err := parseWebhookImportID(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q", tt.id)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if serverID != tt.serverID || webhookID != tt.webhookID {
				t.Fatalf("got (%d, %d), want (%d, %d)", serverID, webhookID, tt.serverID, tt.webhookID)
			}
		})
	}
}

func TestServerScopedClientDoesNotMutateSharedClient(t *testing.T) {
	shared := &postmarkClientForTest{}
	base := shared.client("[SERVER-TOKEN]", "account")

	a := serverScopedClient(base, "server-a")
	b := serverScopedClient(base, "server-b")

	if base.ServerToken != "[SERVER-TOKEN]" {
		t.Fatalf("shared client token changed to %q", base.ServerToken)
	}
	if a.ServerToken != "server-a" || b.ServerToken != "server-b" {
		t.Fatalf("scoped tokens wrong: a=%q b=%q", a.ServerToken, b.ServerToken)
	}
	if a.AccountToken != "account" || b.AccountToken != "account" {
		t.Fatalf("scoped clients lost the account token")
	}
}
