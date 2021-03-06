package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"github.com/132yse/acgzone-server/api/db"
	"github.com/132yse/acgzone-server/api/def"
)

func GetCount(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid, _ := strconv.Atoi(p.ByName("pid"))
	resp, err := db.GetCount(pid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Count{Pid: resp.Pid, Pv: resp.Pv, Cv: resp.Cv}
		sendCountResponse(w, res, 201)
	}
}
//
//func AddPageView(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	pid, _ := strconv.Atoi(p.ByName("pid"))
//
//	err := db.AddPageView(pid)
//	if err != nil {
//		sendErrorResponse(w, def.ErrorDB)
//		return
//	} else {
//		sendErrorResponse(w, def.Success)
//
//	}
//}
