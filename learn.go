package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/rs/xid"
)

//https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/

func (pr PersonRes) read() {
	// vocab2 := make(map[string]VoMaster)

	for _, v := range pr.Split {
		searchVocab(v)
	}
}

func remember(q, a string) {
	db := "./vocabularies/" + q + ".json"
	s := "{\"answer\":[\"" + a + "\"]}"

	var result map[string]interface{}
	var err error
	if err = json.Unmarshal([]byte(s), &result); err != nil {
		log.Fatalln(err)
	}
	var b []byte
	b, err = json.Marshal(result)
	if err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile(db, b, 0644); err != nil {
		log.Fatalln(err)
	}
}

func understand(v string) bool {

	return true
}

func genXid() string {
	return xid.New().String()
	// fmt.Printf("github.com/rs/xid: %s\n", id.String())
}
