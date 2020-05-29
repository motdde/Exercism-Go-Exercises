/*
Package twofer returns a string

One for X, one for me.

Where X is the given name.

However, if the name is missing, return the string:

One for you, one for me.
*/
package twofer

// ShareWith take a string arguement and returns a string
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
