package strutils

import "github.com/fatih/camelcase"

// Split splits the camelcase word and returns a list of words. It also
// supports digits. Both lower camel case and upper camel case are supported.
// For more info please check: http://en.wikipedia.org/wiki/CamelCase
//
// Examples
//
//	"" =>                     [""]
//	"lowercase" =>            ["lowercase"]
//	"Class" =>                ["Class"]
//	"MyClass" =>              ["My", "Class"]
//	"MyC" =>                  ["My", "C"]
//	"HTML" =>                 ["HTML"]
//	"PDFLoader" =>            ["PDF", "Loader"]
//	"AString" =>              ["A", "String"]
//	"SimpleXMLParser" =>      ["Simple", "XML", "Parser"]
//	"vimRPCPlugin" =>         ["vim", "RPC", "Plugin"]
//	"GL11Version" =>          ["GL", "11", "Version"]
//	"99Bottles" =>            ["99", "Bottles"]
//	"May5" =>                 ["May", "5"]
//	"BFG9000" =>              ["BFG", "9000"]
//	"BöseÜberraschung" =>     ["Böse", "Überraschung"]
//	"Two  spaces" =>          ["Two", "  ", "spaces"]
//	"BadUTF8\xe2\xe2\xa1" =>  ["BadUTF8\xe2\xe2\xa1"]
//
// Splitting rules
//
//  1. If string is not valid UTF-8, return it without splitting as
//     single item array.
//  2. Assign all unicode characters into one of 4 sets: lower case
//     letters, upper case letters, numbers, and all other characters.
//  3. Iterate through characters of string, introducing splits
//     between adjacent characters that belong to different sets.
//  4. Iterate through array of split strings, and if a given string
//     is upper case:
//     if subsequent string is lower case:
//     move last character of upper case string to beginning of
//     lower case string
func Split(s string) []string {
	return camelcase.Split(s)
}

// SplitAll splits a slice of strings into a slice of words by splitting on camelcase.
func SplitAll(names ...string) []string {
	ss := make([]string, 0, len(names))
	for _, name := range names {
		ss = append(ss, Split(name)...)
	}
	return ss
}
