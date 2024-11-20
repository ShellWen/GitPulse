package locks

import "strconv"

const (
	lockKeyPrefix = "lock"
	separator     = "|"
)

const (
	UpdateDeveloper = iota
	UpdateCreatedRepo
	UpdateStarredRepo

	UpdateFollowing
	UpdateFollower

	UpdateContributionOfUser
	UpdateIssuePROfUser
	UpdateCommentOfUser
	UpdateReviewOfUser

	UpdateRepo
	UpdateFork

	UpdateRegion
	UpdateLanguages
	UpdatePulsePoint
	UpdateSummary
)

func GetNewLockKey(updateType int, id int64) string {
	return lockKeyPrefix + separator + strconv.Itoa(updateType) + separator + strconv.Itoa(int(id))
}
