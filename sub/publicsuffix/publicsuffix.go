// A public suffix is one under which Internet users can directly register names.
// TLD (top level domain) com au
// eTLD (effective TLD) com.au
// Another name for "an eTLD" is "a public suffix".
// Often, what's more of interest is the eTLD+1
// https://publicsuffix.org/
package publicsuffix

import "golang.org/x/net/publicsuffix"

var (
	// 返回域名的 eTLD+1
	EffectiveTLDPlusOne = publicsuffix.EffectiveTLDPlusOne
)
