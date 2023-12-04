package gcf

import (
	"fmt"
	"github.com/HRMonitorr/packagebuku/service"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GetOneEmployee", GetOneEmployee)
}

func GetOneEmployee(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://rofinafiin.github.io, https://hrmonitorr.github.io, https://hrmonitorr.advocata.me")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization, Token, Login")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://rofinafiin.github.io, https://hrmonitorr.github.io, https://hrmonitorr.advocata.me")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, service.ListRepo("GIHUBTOKEN", r))

}
