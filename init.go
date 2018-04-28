package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// Person To make JSON file, the format should be CAPITAL LETTER
type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
	NickName   string
	ID         string
}

//PersonRes to store the responses
type PersonRes struct {
	String string
	Split  []string
}

func readStim(q bool, c chan string) {
	if q {
		close(c)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		// handle error.
		log.Fatalln(scanner.Err())
		// c <- strings.Join([]string{"readStim() err:", scanner.Err().Error()}, " ")
	}
	c <- strings.ToLower(removeWS(scanner.Text()))
}

func toTitledStr(s string) string {

	// Function replacing words (assuming lower case input)
	replace := func(w string) string {
		switch w {
		case "with", "in", "a":
			return w
		}
		return strings.Title(w)
	}

	r := regexp.MustCompile(`\w+`)
	s = r.ReplaceAllStringFunc(s, replace)
	return s
}

func removeWS(s string) string {

	// r, _ := regexp.Compile(`^\s+|\s{2,}|\s+$`)
	//First, remove start & end whitespace
	r1, _ := regexp.Compile(`^\s+|\s+$`)
	//then, remove middle double whitespace
	r2, _ := regexp.Compile(`\s{2,}`)

	s = r1.ReplaceAllString(s, "")
	s = r2.ReplaceAllString(s, " ")

	return s
}

func (pr *PersonRes) personRes(r string) {
	(*pr).String = r
	(*pr).Split = strings.Split(r, " ")
	// return strings.Join([]string(c), ", ")
	// var a [n]byte
	// copy(a[:], split[0])
	// s := string(a[:3]) adjkadlfa
}

func (pr *PersonRes) talkAbout(s string) bool {
	s0 := pr.Split[0]
	// fmt.Println(s)
	switch s {
	case "endConversation":
		var a [3]byte
		copy(a[:], s0)
		for _, v := range vocab.Exit {
			// fmt.Println(s0, "==", v)
			if s0 == v || string(a[:3]) == v {
				return true
			}
		}
	case "endMemory":
		for _, v := range vocab.Delete {
			// fmt.Println(s0, v)
			if s0 == v {
				return true
			}
		}
	}

	return false
}
