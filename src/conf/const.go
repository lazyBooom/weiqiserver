package conf

//基础信息
const (
	BLACK_PLAYER    = 0 // 黑子玩家
	WHITE_PLAYER    = 1 // 白棋玩家
	WINNER          = 1 // 胜利者
	FAILER          = 2 // 失败者
	UN_FINISHED     = 0 // 未完成对局
	FINISHED_WIN    = 1 // 完成并且胜利
	FINISHED_FAILED = 2 // 完成并且失败
)

//错误码
const (
	SUCCEED            = 0
	ERR_BAD_PARAM      = 100 //错误滴参数
	ERR_SERVER_ERR     = 113 //服务器错误
	ERR_INVITE_OFFLINE = 101 //邀请好友不在线
)
