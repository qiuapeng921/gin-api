package consts

const (
	SUCCESS = 200
	ERROR   = 500
)

var MsgFlags = map[int]string{
	SUCCESS: "success",
	ERROR:   "error",
}

func GetMsg(code int) string {
	message, ok := MsgFlags[code]
	if ok {
		return message
	}

	return MsgFlags[ERROR]
}
