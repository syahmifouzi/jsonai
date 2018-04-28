package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//https://stackoverflow.com/questions/43713952/modifying-json-file-using-golang?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa

func exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			return false, nil
		}
		// exist but got other error
		return true, err
	}
	return true, nil
}

func (p *Person) meetNewPerson() {
	fmt.Println("Hai, " + p.ID + ".")
	fmt.Println("Nice to meet you.")

	fmt.Println("What's your full name?")
	c := make(chan string)
	go readStim(false, c)
	fn := toTitledStr(<-c)
	fmt.Println(fn, "got it!")
	p.addPerson(fn)
}

func (p *Person) checkPerson() {
	// fmt.Println(nn)
	db := "./database/person/" + p.ID + ".json"
	// path := filepath.Join("person", p.ID)
	exs, err := exists(db)
	if err != nil {
		log.Fatalln(err)
	}
	if !exs {
		// if err = os.MkdirAll("person", 0644); err != nil {
		// 	log.Fatalln(err)
		// }

		p.meetNewPerson()
		var b []byte
		b, err = json.Marshal(p)
		if err != nil {
			log.Fatalln(err)
		}
		if err = ioutil.WriteFile(db, b, 0644); err != nil {
			log.Fatalln(err)
		}

		return
	}
	var content []byte
	content, err = ioutil.ReadFile(db)
	if err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(content, &p); err != nil {
		log.Fatalln(err)
	}
}

func (p *Person) addPerson(fn string) {

	split := strings.Split(fn, " ")
	l := len(split) - 1
	mn := ""
	if len(split) > 2 {
		mn = split[1]
		for i := 2; i < l; i++ {
			mn += " " + split[i]
		}
	}

	*p = Person{
		FirstName:  split[0],
		MiddleName: mn,
		LastName:   split[l],
		ID:         p.ID,
		NickName:   p.ID,
	}
	// *p.FirstName = split[0]
}

func (p *Person) deletePerson() {

	db := "./database/person/" + p.ID + ".json"

	if err := os.Remove(db); err != nil {
		log.Fatalln(err)
	}
}
