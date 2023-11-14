//go:build !boringcrypto && !goexperiment.systemcrypto && !goexperiment.cngcrypto && !goexperiment.opensslcrypto

package fips

func IsAvailable() bool {
	return false
}
