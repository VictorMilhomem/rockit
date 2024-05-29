package main

import (
	"log"
	"os"
  "time"
	"github.com/VictorMilhomem/rockit/cmd/playlist"
	"github.com/gopxl/beep"
  "github.com/gopxl/beep/mp3"
  "github.com/gopxl/beep/speaker"
)


func playMusic(path string) error {
	
	f, err := os.Open(path)
	if err != nil {
     return err
	}
 	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done 


	return nil
}



func main() {
	 //err := playMusic("/home/lil/Music/Charlie Brown/Charlie Brown - Ceu Azul.mp3")
	list := playlist.NewPlaylist("/home/lil/Music/Charlie Brown", "Charlie Brown")
	err := list.FetchPlaylistMusics()
	if err != nil {
	  	log.Fatal(err)
	 }
	
}
