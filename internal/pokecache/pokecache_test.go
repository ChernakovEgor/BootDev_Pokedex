package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key   string
		value string
	}{
		{"http://example.com", "data"},
		{"http://google.com", "another_data"},
		{"http://pathofexile.com", "epic quarterstaf"},
	}

	cache := NewCache(time.Second * 5)

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			cache.Add(c.key, []byte(c.value))

			got, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("no entry for key %s", c.key)
			}

			if string(got) != c.value {
				t.Errorf("got %s want %s", string(got), c.value)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	cacheTime := time.Millisecond * 5
	waitTime := time.Millisecond * 10

	cache := NewCache(cacheTime)
	cache.Add("http://example.com", []byte("data"))

	if _, ok := cache.Get("http://example.com"); !ok {
		t.Error("expected to get value")
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("http://example.com"); ok {
		t.Error("expected to get no value")
	}
}
