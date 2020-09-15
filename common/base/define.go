package base

type BaseRequest struct {
	Uid      int64  `form:"uid" json:"uid" binding:"required"`
	DeviceId string `form:"udid" json:"udid" binding:"required"`
}

type Response struct {
	Code        ResponseCode `json:"status"`
	Msg         string       `json:"msg"`
	Data        interface{}  `json:"types"`
	ServiceTime int64        `json:"servicetime"`
}

//ResponseCode 状态返回码
type ResponseCode int

const (
	ResponseCode_Succ ResponseCode = 0

	//1000~1999 通用
	ResponseCode_ServerError   ResponseCode = 1000
	ResponseCode_Param_Less    ResponseCode = 1001
	ResponseCode_Param_Invalid ResponseCode = 1002
	ResponseCode_Data_Invalid  ResponseCode = 1003
	ResponseCode_Data_NotExist ResponseCode = 1004
)

var responseCodeToMsg = map[ResponseCode]string{
	ResponseCode_Succ: "请求成功",

	//1000~1999 通用
	ResponseCode_ServerError:   "服务器错误",
	ResponseCode_Param_Less:    "参数不足",
	ResponseCode_Param_Invalid: "参数无效",
	ResponseCode_Data_Invalid:  "数据无效",
	ResponseCode_Data_NotExist: "数据不存在",
}

type DataStatusCode int

const (
	DataStatusCode_Normal  DataStatusCode = 0
	DataStatusCode_Deleted DataStatusCode = 1

	DataStatusCode_Online  DataStatusCode = 0
	DataStatusCode_Offline DataStatusCode = 1

	DataStatusCode_NotReward DataStatusCode = 0
	DataStatusCode_Rewarded  DataStatusCode = 1
)

const (
	ANDROID string = "android"
	IOS     string = "ios"
	IPAD    string = "ipad"
)
