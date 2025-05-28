package kebabcase

import "testing"

var ops int = 1e6

type sample struct {
	str, out string
}

func TestKebabcase(t *testing.T) {
	samples := []sample{
		{"@49L0S145_¬fwHƒ0TSLNVp", "49l0s145-fw-h-0tslnvp"},
		{"lk0B@bFmjrLQ_Z6YL", "lk0-b-b-fmjr-lq-z6yl"},
		{"samPLE text", "sam-ple-text"},
		{"sample text", "sample-text"},
		{"sample-text", "sample-text"},
		{"sample_text", "sample-text"},
		{"sample___text", "sample-text"},
		{"sampleText", "sample-text"},
		{"inviteYourCustomersAddInvites", "invite-your-customers-add-invites"},
		{"sample 2 Text", "sample-2-text"},
		{"   sample   2    Text   ", "sample-2-text"},
		{"   $#$sample   2    Text   ", "sample-2-text"},
		{"SAMPLE 2 TEXT", "sample-2-text"},
		{"___$$Base64Encode", "base64-encode"},
		{"FOO:BAR$BAZ", "foo-bar-baz"},
		{"FOO#BAR#BAZ", "foo-bar-baz"},
		{"something.com", "something-com"},
		{"$something%", "something"},
		{"something.com", "something-com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
		{"CStringRef", "cstring-ref"},
		{"5test", "5test"},
		{"test5", "test5"},
		{"THE5r", "the5r"},
		{"5TEst", "5test"},
		{"_5TEst", "5test"},
		{"@%#&5TEst", "5test"},
		{"edf_6N", "edf-6n"},
		{"f_pX9", "f-p-x9"},
		{"p_z9Rg", "p-z9-rg"},
		{"2FA Enabled", "2fa-enabled"},
		{"Enabled 2FA", "enabled-2fa"},
	}

	for _, sample := range samples {
		if out := Kebabcase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}
