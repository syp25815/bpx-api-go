package bpx

type SystemStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// 获取系统状态
func Status() SystemStatus {
	var resp SystemStatus
	newAgent().
		Get(API_BASE + "api/v1/status").
		EndStruct(&resp)
	return resp
}

// Responds：pong
func Ping() string {
	_, data, _ := newAgent().
		Get(API_BASE + "api/v1/ping").
		End()
	return data
}

// 获取当前系统时间：返回毫秒
func SystemTime() string {
	_, data, _ := newAgent().
		Get(API_BASE + "api/v1/time").
		End()
	return data
}
