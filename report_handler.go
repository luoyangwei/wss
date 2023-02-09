package main

type ReportHandler struct {
	// 使用之前继承 Handle
	Handle
}

func (h *ReportHandler) Route() map[string]H {
	return map[string]H{
		"/report": h.getReport,
	}
}

func (h *ReportHandler) getReport(r *Request) {
	h.Response([]byte("hello word!"))
	return
}
