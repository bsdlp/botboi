package kitsu

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/jsonapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLibraryEntry(t *testing.T) {
	testObj := `{
    "data": [
        {
            "id": "35128879",
            "type": "libraryEntries",
            "links": {
                "self": "https://kitsu.io/api/edge/library-entries/35128879"
            },
            "attributes": {
                "createdAt": "2018-09-17T07:23:40.396Z",
                "updatedAt": "2018-09-28T08:16:08.529Z",
                "status": "current",
                "progress": 3,
                "volumesOwned": 0,
                "reconsuming": false,
                "reconsumeCount": 0,
                "notes": null,
                "private": false,
                "reactionSkipped": "unskipped",
                "progressedAt": "2018-09-28T08:16:08.528Z",
                "startedAt": "2018-09-28T08:15:32.510Z",
                "finishedAt": null,
                "rating": "0.0",
                "ratingTwenty": null
            },
            "relationships": {
                "user": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/user",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/user"
                    }
                },
                "anime": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/anime",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/anime"
                    }
                },
                "manga": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/manga",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/manga"
                    }
                },
                "drama": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/drama",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/drama"
                    }
                },
                "review": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/review",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/review"
                    }
                },
                "mediaReaction": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/media-reaction",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/media-reaction"
                    }
                },
                "media": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/media",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/media"
                    }
                },
                "unit": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/unit",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/unit"
                    }
                },
                "nextUnit": {
                    "links": {
                        "self": "https://kitsu.io/api/edge/library-entries/35128879/relationships/next-unit",
                        "related": "https://kitsu.io/api/edge/library-entries/35128879/next-unit"
                    }
                }
            }
        }
    ],
    "meta": {
        "statusCounts": {
            "planned": 29,
            "current": 5,
            "completed": 9
        },
        "count": 43
    },
    "links": {
        "first": "https://kitsu.io/api/edge/library-entries?filter%5Bkind%5D=anime&filter%5Buser_id%5D=373871&page%5Blimit%5D=40&page%5Boffset%5D=0&sort=status%2C-progressed_at",
        "next": "https://kitsu.io/api/edge/library-entries?filter%5Bkind%5D=anime&filter%5Buser_id%5D=373871&page%5Blimit%5D=40&page%5Boffset%5D=40&sort=status%2C-progressed_at",
        "last": "https://kitsu.io/api/edge/library-entries?filter%5Bkind%5D=anime&filter%5Buser_id%5D=373871&page%5Blimit%5D=40&page%5Boffset%5D=3&sort=status%2C-progressed_at"
    }
}`

	out, err := jsonapi.UnmarshalManyPayload(strings.NewReader(testObj), reflect.TypeOf(&LibraryEntry{}))
	require.NoError(t, err)

	assert.Empty(t, out)
}
