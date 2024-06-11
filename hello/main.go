package main

import (
	"flag"
	"fmt"
)

func main() {
	var greetFlag = flag.String("lang", "en", "greeting the user with different languages")
	flag.Parse()
	
	greeting := greet(language(*greetFlag)) 
	fmt.Println(greeting)
}


type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε", // Greek
	"en": "Hello world", // English
	"fr": "Bonjour le monde", // French
	"he": " שלום עולם ", // Hebrew
	"ur": " ہیلو ", // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in various languages
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
