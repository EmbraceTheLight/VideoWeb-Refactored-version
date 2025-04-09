package biz

// 视频热度相关
const (
	// AddHotEachView 每访问一次增加的热度
	AddHotEachView = 1

	// AddHotEachComment 每评论一次增加的热度
	AddHotEachComment = 3 * AddHotEachView

	// AddHotEachBarrage 每发表一次弹幕增加的热度
	AddHotEachBarrage = 3 * AddHotEachComment

	// AddHotEachLike 每点赞一次增加的热度
	AddHotEachLike = 10 * AddHotEachView

	// AddHotEachShell 每投一个贝壳增加的热度
	AddHotEachShell = 3 * AddHotEachView

	// AddHotEachFavorite 每收藏一次增加的热度
	AddHotEachFavorite = 75 * AddHotEachView

	// AddHotEachShare 每分享一次增加的热度
	AddHotEachShare = 50 * AddHotEachView
)

// 用户-视频状态相关
const (
	LikeStatus = 1 << iota
	FavoriteStatus
	ShareStatus
)
