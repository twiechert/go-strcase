// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package strcase

// DotCase converts a string into dot case.
func DotCase(s string) string {
	return lowerDelimiterCase(s, '.')
}
