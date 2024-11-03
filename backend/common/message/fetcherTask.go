package message

type TaskType int64

const (
	FetchDeveloper = iota
	FetchCreatedRepo
	FetchStarredRepo
	FetchFollow
	FetchFollowing
	FetchFollower
	FetchContributionOfUser
	FetchIssuePROfUser
	FetchCommentOfUser

	FetchRepo
	FetchFork
)

type FetcherTask struct {
	Type TaskType `json:"type"`
	Id   int64    `json:"id"`
}
