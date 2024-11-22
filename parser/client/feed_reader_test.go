package client

import (
	"path"
	"testing"

	. "github.com/slipros/devicedetector/parser"
	"gotest.tools/assert"
)

func TestFeedReaderParse(t *testing.T) {
	ps := NewFeedReader(path.Join(dir, FixtureFileFeedReader))
	var list []*ClientFixture
	err := ReadYamlFile(`fixtures/feed_reader.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		assert.DeepEqual(t, item.ClientMatchResult, r)
	}
}
