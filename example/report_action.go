package main

import "github.com/luoyangwei/wss"

type ReportAction struct {
	// 使用之前继承 Handle
	wss.ResponseEmpowerment
}

func (action *ReportAction) Route() wss.Routers {
	return []wss.Router{
		{
			Method:  wss.MethodPost,
			Pattern: "/getReport",
			Action:  action.getReport,
		},
	}
}

func (action *ReportAction) getReport(request *wss.Request) {
	action.Response([]byte("hello word!"))
}
