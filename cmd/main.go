package main

import (
	"log"
	"os"
  "time"
  "path/filepath"
  "fmt"

	"github.com/gopxl/beep"
  "github.com/gopxl/beep/mp3"
  "github.com/gopxl/beep/speaker"
)


type Playlist struct {
	name string
	path string
	musics []string
}

func NewPlaylist(path, name string) *Playlist {
	return &Playlist {
		name: name,
		path: path,
		musics: nil,
	}
}


func (p *Playlist) fetchPlaylistMusics() error {
	dir := p.path
	var musics []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		musics = append(musics, path)
		
		return nil
	})

	if err != nil {
		return err
	}
	
	p.musics = musics
	for _, file := range p.musics {
		fmt.Println(file)
	}

	return nil
}



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
	list := NewPlaylist("/home/lil/Music/Charlie Brown", "Charlie Brown")
	err := list.fetchPlaylistMusics()
	if err != nil {
	  	log.Fatal(err)
	 }
	
}
