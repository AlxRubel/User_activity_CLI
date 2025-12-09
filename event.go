package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type EventRaw struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Actor      Actor           `json:"actor"`
	Repo       Repo            `json:"repo"`
	Payload    json.RawMessage `json:"payload"`
	Public     bool            `json:"public"`
	Created_at string          `json:"created_at"`
	Org        Org             `json:"org"`
}
type Payload interface {
	humanReadable() string
}
type Event struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Actor      Actor     `json:"actor"`
	Repo       Repo      `json:"repo"`
	Payload    Payload   `json:"payload"`
	Public     bool      `json:"public"`
	Created_at time.Time `json:"created_at"`
	Org        Org       `json:"org"`
}

func (ev *Event) UnmarshalJSON(data []byte) error {
	eventRaw := EventRaw{}
	if err := json.Unmarshal(data, &eventRaw); err != nil {
		return err
	}

	ev.ID = eventRaw.ID
	ev.Type = eventRaw.Type
	ev.Actor = eventRaw.Actor
	ev.Repo = eventRaw.Repo
	ev.Public = eventRaw.Public
	ev.Created_at, _ = time.Parse(time.RFC3339, eventRaw.Created_at)
	ev.Org = eventRaw.Org

	switch eventRaw.Type {
	case "CreateEvent":
		ev.Payload = &CreateEvent{}
	case "ForkEvent":
		ev.Payload = &ForkEvent{}
	case "IssueCommentEvent":
		ev.Payload = &IssueCommentEvent{}
	case "IssuesEvent":
		ev.Payload = &IssuesEvent{}
	case "PullRequestEvent":
		ev.Payload = &PullRequestEvent{}
	case "PushEvent":
		ev.Payload = &PushEvent{}
	case "WatchEvent":
		ev.Payload = &WatchEvent{}
	default:
		ev.Payload = nil
		return nil
	}

	if err := json.Unmarshal(eventRaw.Payload, ev.Payload); err != nil {
		return err
	}

	return nil
}

type Actor struct {
	// ID            int    `json:"id"`
	Login         string `json:"login"`
	Display_login string `json:"display_login"`
	// Gravatar_ID   string `json:"gravatar_id"`
	// URL           string `json:"url"`
	// Avatar_URL    string `json:"avatar_url"`
}

type Repo struct {
	// ID   int    `json:"id"`
	Name string `json:"name"`
	// URL  string `json:"url"`
}

type Org struct {
	// ID          int    `json:"id"`
	Login string `json:"login"`
	// Gravatar_ID string `json:"gravatar_id"`
	// URL         string `json:"url"`
	// Avatar_URL  string `json:"avatar_url"`
}

func (event Event) humanReadable() string {
	var result string
	switch event.Type {
	case "CreateEvent":
		result = "[" + event.Created_at.String() + "] " + event.Actor.Login +
			" created " + event.Payload.(*CreateEvent).Master_branch +
			" in " + event.Repo.Name
	case "ForkEvent":
		result = "[" + event.Created_at.String() + "] " + event.Actor.Login +
			" forked "
	case "IssueCommentEvent":
	case "IssuesEvent":
	case "PullRequestEvent":
	case "PushEvent":
		resp, err := http.Get("https://api.github.com/repos/" + event.Repo.Name + "/compare/" + event.Payload.(*PushEvent).Before + "..." + event.Payload.(*PushEvent).Head)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		commits := PushEventCompare{}
		err = json.Unmarshal(body, &commits)
		if err != nil {
			panic(err)
		}

		result = "[" + event.Created_at.String() + "] " + event.Actor.Login +
			" pushed " + strconv.Itoa(commits.Total_commits) + " commits to " +
			event.Repo.Name

	case "WatchEvent":
		result = "[" + event.Created_at.String() + "] " + event.Actor.Login +
			" " + event.Payload.(*WatchEvent).Action + " " + event.Repo.Name
	default:
		return ""
	}
	return result
}
