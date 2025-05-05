package biz

const (
	byUpvote  = "cnt_upvote"
	byCreated = "created_at"

	separator = "::"

	asc  = "asc"
	desc = "desc"
)

const (
	HotEachView = 1

	// HotEachComment 每评论一次增加的热度
	HotEachComment = 10 * HotEachView
)
