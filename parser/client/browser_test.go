package client

import (
	"path"
	"testing"

	. "github.com/slipros/devicedetector/parser"
	"gotest.tools/assert"
)

func TestBrowserParse(t *testing.T) {
	ps := NewBrowser(path.Join(dir, FixtureFileBrowser))
	var list []*ClientFixture
	err := ReadYamlFile(`fixtures/browser.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		assert.DeepEqual(t, item.ClientMatchResult, r)
	}
}
