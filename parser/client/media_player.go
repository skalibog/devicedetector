package client

import (
	"path"
)

const ParserNameMediaPlayer = `mediaplayer`
const FixtureFileMediaPlayer = `mediaplayers.yml`

func init() {
	RegClientParser(ParserNameMediaPlayer,
		func(dir string) ClientParser {
			return NewMediaPlayer(path.Join(dir, FixtureFileMediaPlayer))
		})
}

func NewMediaPlayer(fileName string) *MediaPlayer {
	c := &MediaPlayer{}
	c.ParserName = ParserNameMediaPlayer
	if err := c.Load(fileName); err != nil {
		return nil
	}
	return c
}

// Client parser for mediaplayer detection
type MediaPlayer struct {
	ClientParserAbstract
}
