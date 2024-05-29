package playlist


import (
	"fmt"
	
	"os"
	"path/filepath"
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


func (p *Playlist) FetchPlaylistMusics() error {
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



