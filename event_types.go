package main

type CreateEvent struct {
	Ref           string `json:"ref"`
	Ref_type      string `json:"ref_type"`
	Full_ref      string `json:"full_ref"`
	Master_branch string `json:"master_branch"`
	Description   string `json:"description"`
	Pusher_type   string `json:"pusher_type"`
}

func (event CreateEvent) humanReadable() string {
	return ""
}

// FIXME not complete
type ForkEvent struct {
	Action string `json:"action"`
}

func (event ForkEvent) humanReadable() string {
	return ""
}

// FIXME not complete
type IssueCommentEvent struct {
	Action string `json:"action"`
}

func (event IssueCommentEvent) humanReadable() string {
	return ""
}

// FIXME not complete
type IssuesEvent struct {
	Action string `json:"action"`
}

func (event IssuesEvent) humanReadable() string {
	return ""
}

// FIXME not complete
type PullRequestEvent struct {
	Action string `json:"action"`
}

func (event PullRequestEvent) humanReadable() string {
	return ""
}

type PushEvent struct {
	// Repository_ID int    `json:"repository_id"`
	// Push_ID       int    `json:"push_id"`
	Ref    string `json:"ref"`
	Head   string `json:"head"`
	Before string `json:"before"`
}
type PushEventCompare struct {
	Total_commits int `json:"total_commits"`
}

func (event PushEvent) humanReadable() string {
	return "Pushed to "
}

type WatchEvent struct {
	Action string `json:"action"`
}

func (event WatchEvent) humanReadable() string {
	return ""
}
