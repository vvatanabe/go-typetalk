package typetalk

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_MentionsService_ReadMention_should_read_a_mention(t *testing.T) {
	setup()
	defer teardown()
	mentionId := 1
	b, _ := ioutil.ReadFile("../testdata/read-mention.json")
	mux.HandleFunc(fmt.Sprintf("/mentions/%d", mentionId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Mentions.ReadMention(context.Background(), mentionId)
	if err != nil {
		t.Errorf("returned error: %v", err)
	}
	var want *struct {
		Mention *Mention `json:"mention"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Mention) {
		t.Errorf("returned content: got  %v, want %v", result.ID, want.Mention.ID)
	}
}

func Test_MentionsService_GetMentionList_should_get_some_mentions(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/get-mention-list.json")
	mux.HandleFunc("/mentions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValues(t, r, values{
			"from":   10,
			"unread": true,
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Mentions.GetMentionList(context.Background(), &GetMentionListOptions{10, true})
	if err != nil {
		t.Errorf("returned error: %v", err)
	}
	var want *struct {
		Mentions []*Mention `json:"mentions"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Mentions) {
		t.Errorf("returned content: got  %v, want %v", result, want.Mentions)
	}
}
