package service

import (
	"context"
	"encoding/json"
	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/HRMonitorr/packagebuku"
	"net/http"
	"os"
)

func ListRepo(personalToken string, r *http.Request) string {
	resp := new(packagebuku.ResponseBack)
	data := new(packagebuku.Req)
	resp.Status = http.StatusBadRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = http.StatusOK
		response, er := packagebuku.ListRepositoriesOnlydDetail(context.Background(), os.Getenv(personalToken), data.OwnerName)
		if er != nil {
			resp.Message = "error " + er.Error()
		}
		resp.Message = "Berhasil ambil list"
		resp.Data = response
	}
	response := pasproj.ReturnStringStruct(resp)
	return response

}
