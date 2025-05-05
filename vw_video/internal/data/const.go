package data

// 视频热度相关
const (
	// HotEachView 每访问一次增加的热度
	HotEachView = 1

	// HotEachComment 每评论一次增加的热度
	HotEachComment = 10 * HotEachView

	// HotEachReply 每回复一次增加的热度
	HotEachReply = HotEachComment / 2

	// HotEachBarrage 每发表一次弹幕增加的热度
	HotEachBarrage = 5 * HotEachView

	// HotEachUpvote 每点赞一次增加的热度
	HotEachUpvote = 2 * HotEachView

	// HotEachShell 每投一个贝壳增加的热度
	HotEachShell = 3 * HotEachView

	// HotEachFavorite 每收藏一次增加的热度
	HotEachFavorite = 20 * HotEachView

	// HotEachShare 每分享一次增加的热度
	HotEachShare = 50 * HotEachView
)
