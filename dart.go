package main

import (
	"regexp"
	"strings"
)

// #000 => 0XFF000 // As Flutter Color System Wants
func getColor(s string) string {
	r := regexp.MustCompile(`#`)
	return r.ReplaceAllString(strings.ToUpper(s), "0xFF")
}

//  "black: '#000'," => ("black" , "0XFF000")
func getUngroupedNameAndColor(s string) (name string, color string) {
	//Get The Name
	r := regexp.MustCompile(`\w+:`)
	name = r.FindString(s)
	name = name[:len(name)-1]

	//Get The Color
	if strings.Contains(s, "#") {
		r = regexp.MustCompile(`'#\w+'`)
		color = r.FindString(s)
		color = color[1 : len(color)-1]
		color = getColor(color)
		return
	}

	r = regexp.MustCompile(`'\w+'`)
	color = r.FindString(s)
	color = color[1 : len(color)-1]

	return

}

// "100: '#ebf8ff'," => ("100" , "0XFFEBF8FF" )
func getGroupedNameAndColor(s string) (name string, color string) {
	r := regexp.MustCompile(`\w+:`)
	name = r.FindString(s)
	name = name[:len(name)-1]

	r = regexp.MustCompile(`'#\w+'`)
	color = r.FindString(s)
	color = color[1 : len(color)-1]
	color = getColor(color)
	return
}

// "blue: {" => "blue"
func getWrapperColor(s string) (color string) {
	r := regexp.MustCompile(`\w+:`)
	color = r.FindString(s)
	color = color[:len(color)-1]

	return
}
