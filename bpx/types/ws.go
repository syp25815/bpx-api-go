package types

type WsResp struct {
	Stream string `json:"stream"`
}

type WsDepth struct {
	Data struct {
		E  int64      `json:"E"`
		T  int64      `json:"T"`
		U  int        `json:"U"`
		A  [][]string `json:"a"`
		B  [][]string `json:"b"`
		E1 string     `json:"e"`
		S  string     `json:"s"`
		U1 int        `json:"u"`
	} `json:"data"`
	Stream string `json:"stream"`
}

type WsTrade struct {
	Data struct {
		E  int64  `json:"E"`
		T  int64  `json:"T"`
		A  string `json:"a"`
		B  string `json:"b"`
		E1 string `json:"e"`
		M  bool   `json:"m"`
		P  string `json:"p"`
		Q  string `json:"q"`
		S  string `json:"s"`
		T1 int    `json:"t"`
	} `json:"data"`
	Stream string `json:"stream"`
}
type WsTicker struct {
	Data struct {
		E  int64  `json:"E"`
		V  string `json:"V"`
		C  string `json:"c"`
		E1 string `json:"e"`
		H  string `json:"h"`
		L  string `json:"l"`
		N  int    `json:"n"`
		O  string `json:"o"`
		S  string `json:"s"`
		V1 string `json:"v"`
	} `json:"data"`
	Stream string `json:"stream"`
}
type WsBookTicker struct {
	Data struct {
		A  string `json:"A"`
		B  string `json:"B"`
		E  int64  `json:"E"`
		T  int64  `json:"T"`
		A1 string `json:"a"`
		B1 string `json:"b"`
		E1 string `json:"e"`
		S  string `json:"s"`
		U  int    `json:"u"`
	} `json:"data"`
	Stream string `json:"stream"`
}

type WsKline struct {
	Data struct {
		E  int64  `json:"E"`
		T  string `json:"T"`
		X  bool   `json:"X"`
		C  string `json:"c"`
		E1 string `json:"e"`
		H  string `json:"h"`
		L  string `json:"l"`
		N  int    `json:"n"`
		O  string `json:"o"`
		S  string `json:"s"`
		T1 string `json:"t"`
		V  string `json:"v"`
	} `json:"data"`
	Stream string `json:"stream"`
}

type WsOrder struct {
	Data struct {
		E  int64  `json:"E"`
		S  string `json:"S"`
		T  int64  `json:"T"`
		V  string `json:"V"`
		X  string `json:"X"`
		Z  string `json:"Z"`
		E1 string `json:"e"`
		F  string `json:"f"`
		I  string `json:"i"`
		O  string `json:"o"`
		P  string `json:"p"`
		Q  string `json:"q"`
		S1 string `json:"s"`
		Z1 string `json:"z"`
	} `json:"data"`
	Stream string `json:"stream"`
}
