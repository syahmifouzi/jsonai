package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//https://stackoverflow.com/questions/14668850/list-directory-in-go?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa

var master = [5]string{
	"MeMeijin",
	"MiMo",
	"mimo",
	"MiMi",
	"mimi",
}

var vocab VoMaster

// var mapVocab map[string]interface{}

// VoMaster ...
type VoMaster struct {
	Exit         []string `json:"exit"`
	Delete       []string `json:"delete"`
	FirstPerson  []string `json:"firstPerson"`
	SecondPerson []string `json:"secondPerson"`
	ThirdPerson  []string `json:"thirdPerson"`
	Question     []string `json:"question"`
}

func initVocab() {
	db := "./database/vocabulary.json"
	content, err := ioutil.ReadFile(db)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println("content:", content)
	if err = json.Unmarshal(content, &vocab); err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(map[string]vocab)
	// if err = json.Unmarshal(content, &mapVocab); err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(vocab.Master)
}

func searchVocab(s string) {
	db := "./vocabularies/" + s + ".json"
	a := ""
	exs, err := exists(db)
	if err != nil {
		log.Fatalln(err)
	}
	if !exs {
		// if err = os.MkdirAll("person", 0644); err != nil {
		// 	log.Fatalln(err)
		// }
		fmt.Println("comel say: ermm...")
		fmt.Println("comel say: I dont understand...")
		fmt.Println("comel say: Can you teach me the meaning please?")
		c := make(chan string)
		go readStim(false, c)
		r := <-c
		if r == "y" {
			fmt.Println("comel say: Okay!")
			go readStim(false, c)
			a = <-c
			fmt.Println("comel say:", s, "-->", a, "?")
			go readStim(false, c)
			r = <-c
			if r == "y" {
				fmt.Println("Got it!")
				remember(s, a)
			}
		}

		return
	}
	var replys map[string]interface{}
	content, err := ioutil.ReadFile(db)
	if err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(content, &replys); err != nil {
		log.Fatalln(err)
	}

	rep := replys["answer"].([]interface{})

	for _, r := range rep {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(r.(string))
	}

}

// files, err := ioutil.ReadDir("./vocabularies/")
// if err != nil {
// 	log.Fatalln(err)
// }

// for _, f := range files {
// 	fmt.Println(f.Name())
// }
