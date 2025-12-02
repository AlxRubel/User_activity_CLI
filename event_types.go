

type CommitCommentEvent struct {
}

type CreateEvent struct {
}
type DeleteEvent struct {
}
type DiscussionEvent struct {
}
type ForkEvent struct {
}
type GollumEvent struct {
}
type IssueCommentEvent struct {
}
type IssuesEvent struct {
}
type MemberEvent struct {
}
type PublicEvent struct {
}
type PullRequestEvent struct {
}
type PullRequestReviewEvent struct {
}
type PullRequestReviewCommentEvent struct {
}
type PushEvent struct {
	Repository_ID int    `json:"repository_id"`
	Push_ID       int    `json:"push_id"`
	Ref           string `json:"ref"`
	Head          string `json:"head"`
	Before        string `json:"before"`
}
type ReleaseEvent struct {
}
type WatchEvent struct {
}