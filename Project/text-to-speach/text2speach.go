package main

import htgotts "github.com/hegedustibor/htgo-tts"

func main() {
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("heyyyy hie abc i hope you are doing great")
}
