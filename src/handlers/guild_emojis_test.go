package handlers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuildEmojis(t *testing.T) {
	t.Run("checkRoleWhitelist", func(t *testing.T) {
		fixtures := []struct {
			whitelist []string
			roles     []string
			pass      bool
		}{
			{
				[]string{"a", "b", "c"},
				[]string{"d", "e", "a"},
				true,
			},
			{
				[]string{"a", "b", "c"},
				[]string{"d", "e", "f"},
				false,
			},
			{
				[]string{},
				[]string{},
				true,
			},
			{
				[]string{},
				[]string{"a"},
				true,
			},
		}

		for i, fixture := range fixtures {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				assert.Equal(t, fixture.pass, checkEmojiRoleWhitelist(fixture.whitelist, fixture.roles))
			})
		}
	})
}
