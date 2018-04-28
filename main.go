package main

import (
	"fmt"
)

//https://stackoverflow.com/questions/14230145/what-is-the-best-way-to-convert-byte-array-to-string?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
//https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa

func main() {

	fmt.Println("Huarrgh...")

	initVocab()
	c := make(chan string)

	fmt.Println("Hai...")
	fmt.Println("I am Comel")
	fmt.Println("To whom am I speaking with?")
	go readStim(false, c)
	n := <-c
	for _, mn := range master {
		if n == mn {
			fmt.Println("Hello there, Master", mn, "!")
			// checkPerson(ctx, client, n)
			loveMaster()
			fmt.Println("comel say: Bye!!!")
			return
		}
	}
	fmt.Println("Hmm...")
	p := Person{
		ID: n,
	}
	p.checkPerson()
	go p.interactions(c)
	fmt.Println(<-c)

}

func (p *Person) interactions(c chan string) {

	c2 := make(chan string)

	fmt.Println("comel say: Hello", p.FirstName)
	fmt.Println("comel say: How can I help?")

	go readStim(false, c2)

	for r := range c2 {
		go func(r2 string) {
			pr := PersonRes{}
			pr.personRes(r2)
			if pr.talkAbout("endConversation") {
				fmt.Println("comel say:", r2)
				readStim(true, c2)
			} else {
				fmt.Println("comel say:", r2)
				if pr.talkAbout("endMemory") {
					fmt.Println("comel say: I will now forget about you")
					p.deletePerson()
					readStim(true, c2)
				}
				readStim(false, c2)
			}

		}(r)
	}

	c <- "Done interacting"
}

func loveMaster() {
	c := make(chan string)
	go readStim(false, c)

	for r := range c {
		go func(r2 string) {
			pr := PersonRes{}
			pr.personRes(r2)
			if pr.talkAbout("endConversation") {
				readStim(true, c)
			} else {
				pr.read()
				readStim(false, c)
			}

		}(r)
	}
}
