package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInplaceMsdSort(t *testing.T) {
	assert := assert.New(t)

	original := []string{
		"bed", "bug", "dad", "yes", "zoo",
		"now", "for", "tip", "ilk", "dim",
		"tag", "jot", "sob", "nob", "sky",
		"hut", "men", "egg", "few", "jay",
		"owl", "joy", "rap", "gig", "wee",
		"was", "wad", "fee", "tap", "tar",
		"dug", "jam", "all", "bad", "yet",
	}
	expected := []string{
		"all", "bad", "bed", "bug", "dad",
		"dim", "dug", "egg", "fee", "few",
		"for", "gig", "hut", "ilk", "jam",
		"jay", "jot", "joy", "men", "nob",
		"now", "owl", "rap", "sky", "sob",
		"tag", "tap", "tar", "tip", "wad",
		"was", "wee", "yes", "yet", "zoo",
	}

	InplaceMsdSort(original)
	assert.Equal(expected, original)
}
