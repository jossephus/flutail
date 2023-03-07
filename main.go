package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const dataFile = "stubs/tailwind.js"
const outputFile = "stubs/tailwindColors.dart"
const whatToLookFor = "colors:"

func main() {

	f, err := os.Open(dataFile)

	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	tailwind, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalln(err)
	}

	tailwindstring := string(tailwind)

	i := strings.Index(tailwindstring, whatToLookFor)

	partToBeFiltered := tailwindstring[i:]

	var openingBraces []rune
	var closingBraces []rune

	// len(whatToLookFor) + 1 => One For Javascript Formatter
	lastIndex := strings.IndexFunc(partToBeFiltered[len(whatToLookFor)+1:], func(r rune) bool {
		if r == '{' {
			openingBraces = append(openingBraces, r)
			//fmt.Println(len(openingBraces))
		}
		if r == '}' {
			closingBraces = append(closingBraces, r)
			//fmt.Println(len(closingBraces))
		}
		if len(openingBraces) == len(closingBraces) {
			return true
		}
		return false
	})

	// len(whatToLookFor) + 2 => 1 for  let's get what's taken from us => 1 for I don't know
	whatWeReallyWant := partToBeFiltered[:lastIndex+len(whatToLookFor)+2]

	Analayze(whatWeReallyWant)
}

func Analayze(data string) {

	splitted := strings.FieldsFunc(data, func(r rune) bool {
		if r == '\n' {
			return true
		}
		return false
	})

	var splittedTrimmed []string

	// splitted[1:len(splitted - 1)] => let's remove the whatToLookFor Part
	foreach(splitted[1:len(splitted)-1], func(s string) {
		splittedTrimmed = append(splittedTrimmed, strings.TrimSpace(s))
	})

	// This is where the main things happen It can be refactored to work with other filters
	var notGrouped = make(map[string]string)
	var upperColor string
	var groupedSlice = make(map[string][]map[string]string)

	flag := false

	foreach(splittedTrimmed, func(s string) {
		if strings.Index(s, "{") != -1 {
			upperColor = getWrapperColor(s)
			//upperColor = s
			flag = true
			return
			//fmt.Println(s)
		}
		if strings.Index(s, "}") != -1 {

			flag = false
			return
		}
		if !flag {

			name, color := getUngroupedNameAndColor(strings.TrimSpace(s))
			notGrouped[name] = color
			return
		}
		name, color := getGroupedNameAndColor(strings.TrimSpace(s))

		groupedSlice[upperColor] = append(groupedSlice[upperColor], map[string]string{
			name: color,
		})
	})
	fmt.Println("Whoo")

	// for k, v := range notGrouped {
	// 	fmt.Printf("%s\t\t=>\t\t%s\n", k, v)
	// 	//fmt.Println("key", k, "value", v)
	// }
	// fmt.Println()
	// fmt.Println()
	// for k, v := range groupedSlice {
	// 	fmt.Println(k, " { ")
	// 	for _, sv := range v {
	// 		for sk, v := range sv {
	// 			fmt.Printf("\t\t%s\t\t=>\t\t%s\n", sk, v)
	// 		}
	// 		//fmt.Println(k, sv)
	// 	}
	// 	fmt.Println("}")
	// }

	createDartFileAndWriteMe(notGrouped, groupedSlice)

}

func createDartFileAndWriteMe(notGrouped map[string]string, groupedSlice map[string][]map[string]string) {
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalln("Can not create dart file ")
	}
	defer f.Close()
	var b bytes.Buffer

	b.WriteString("import 'package:flutter/material.dart';\n")

	for k, v := range notGrouped {
		b.WriteString("const ")
		b.WriteString("tail" + k)
		b.WriteString(" = ")
		b.WriteString(" Color(" + v + ");")
		b.WriteRune('\n')
	}

	// const tailRed100 = Color(0xFF)

	for key, values := range groupedSlice {

		for _, sv := range values {

			for sk, v := range sv {
				b.WriteString("const ")
				b.WriteString("tail" + strings.Title(key) + sk)
				b.WriteString(" = ")
				b.WriteString(" Color(" + v + ");")
				b.WriteRune('\n')
			}
			//fmt.Println(k, sv)
		}
	}

	f.Write([]byte(b.String()))

}

func foreach(data []string, callback func(string)) {
	for _, v := range data {
		callback(v)
	}
}
