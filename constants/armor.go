package constants

// Definitions for armored data
const (
	ArmorHeaderVersion = "GopenPGP 0.0.1 (" + Version + ")"
	ArmorHeaderComment = "https://gopenpgp.org"
	PGPMessageHeader   = "PGP MESSAGE"
	PublicKeyHeader    = "PGP PUBLIC KEY BLOCK"
	PrivateKeyHeader   = "PGP PRIVATE KEY BLOCK"
)