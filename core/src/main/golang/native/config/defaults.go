package config

var (
	defaultNameServers = []string{
		"223.5.5.5",
		"119.29.29.29",
		"8.8.8.8",
		"1.1.1.1",
	}
	defaultFakeIPFilter = []string{
		// stun services
		"+.stun.*.*",
		"+.stun.*.*.*",
		"+.stun.*.*.*.*",

		// Google Voices
		"lens.l.google.com",
		"stun.l.google.com",

		// Nintendo Switch
		"*.n.n.srv.nintendo.net",
	}
	defaultFakeIPRange = "28.0.0.0/8"
)
