// Code generated by goctl. DO NOT EDIT.
package types

type FeedReq struct {
	LatestTime string `form:"latest_time,optional"` // 可选参数，限制返回视频的最新投稿时间戳
	Token      string `form:"token,optional"`       // 可选参数，登录用户设置
}

type FeedResp struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
	NextTime   int64   `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type Video struct {
	Id            int64  `json:"id" form:"id"`                         // 视频唯一标识
	Author        User   `json:"author" form:"author"`                 // 视频作者信息
	PlayUrl       string `json:"play_url" form:"play_url"`             // 视频播放地址
	CoverUrl      string `json:"cover_url" form:"cover_url"`           // 视频封面地址
	FavoriteCount int64  `json:"favorite_count" form:"favorite_count"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count" form:"comment_count"`   // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite" form:"is_favorite"`       // true-已点赞，false-未点赞
	Title         string `json:"title" form:"title"`                   // 视频标题
}

type User struct {
	Id            int64  `json:"id" form:"id"`                         // 用户id
	Username      string `json:"name" form:"username"`                 // 用户名称
	FollowCount   int64  `json:"follow_count" form:"follow_count"`     // 关注总数
	FollowerCount int64  `json:"follower_count" form:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow" form:"is_follow"`           // true-已关注，false-未关注
}

type Comment struct {
	Id         int64  `json:"id" form:"id"`                   //评论id
	User       User   `json:"user" form:"user"`               //评论用户信息
	Content    string `json:"content" form:"content"`         //评论内容
	CreateDate string `json:"create_date" form:"create_date"` //评论发布日期，格式 mm-dd
}

type Req struct {
	Username string `form:"username"` // 用户名，最长32个字符
	Password string `form:"password"` // 密码，最长32个字符
}

type Resp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserID     int64  `json:"user_id"`     // 用户id
	Token      string `json:"token"`       // 用户鉴权token
}

type UserInfoReq struct {
	UserId int64  `form:"user_id"` // 用户ID
	Token  string `form:"token"`   // 鉴权token
}

type UserInfoResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserInfo   User   `json:"user"`        // 用户信息
}

type PublishActionReq struct {
	Token string `form:"token"` // 鉴权token
	Title string `form:"title"` // 视频标题
}

type PublishActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type PublishListReq struct {
	UserID int64  `form:"user_id"` // 用户ID
	Token  string `form:"token"`   // 鉴权token
}

type PublishListResp struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

type FavoriteActionReq struct {
	Token      string `form:"token"`       // 鉴权token
	VideoID    int64  `form:"video_id"`    // 视频ID
	ActionType int32  `form:"action_type"` // 1-点赞，2-取消点赞
}

type FavoriteActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FavoriteListReq struct {
	UserID int64  `form:"user_id"` // 用户ID
	Token  string `form:"token"`   // 鉴权token
}

type FavoriteListResp struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

type CommentActionReq struct {
	Token       string `form:"token"`                 // 用户鉴权token
	VideoId     int64  `form:"video_id"`              // 视频id
	ActionType  int32  `form:"action_type"`           // 1-发布评论，2-删除评论
	CommentText string `form:"comment_text,optional"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `form:"comment_id,optional"`   // 要删除的评论id，在action_type=2的时候使用
}

type CommentActionResq struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	Comment    Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
}

type CommentListReq struct {
	Token   string `form:"token"`    // 鉴权token
	VideoId int64  `form:"video_id"` // 视频id
}

type CommentListResp struct {
	StatusCode  int32     `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`   // 返回状态描述
	CommentList []Comment `json:"comment_list"` // 评论列表
}

type ActionReq struct {
	Token      string `form:"token"`       // 鉴权token
	ToUserId   int64  `form:"to_user_id"`  // 对方ID
	ActionType int32  `form:"action_type"` // 1-关注，2-取消关注
}

type ActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type ListReq struct {
	UserID int64  `form:"user_id"` // 用户ID
	Token  string `form:"token"`   // 鉴权token
}

type ListResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserList   []User `json:"user_list"`   // 用户列表
}
