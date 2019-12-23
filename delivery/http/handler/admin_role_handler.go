package handler

import (
	"encoding/json"
	"net/http"

	"github.com/betsegawlemma/restaurant-rest/user"
	"github.com/julienschmidt/httprouter"
)

// AdminRoleHandler is used to implement role related http requests
type AdminRoleHandler struct {
	roleService user.RoleService
}

// NewAdminRoleHandler initializes and returns new AdminRoleHandler object
func NewAdminRoleHandler(roleSrv user.RoleService) *AdminRoleHandler {
	return &AdminRoleHandler{roleService: roleSrv}
}

// GetRoles handles GET /v1/admin/roles requests
func (arh *AdminRoleHandler) GetRoles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	roles, errs := arh.roleService.Roles()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleRole handles GET /v1/admin/roles/:id requests
func (arh *AdminRoleHandler) GetSingleRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

// PutRole handles PUT /v1/admin/roles/:id requests
func (arh *AdminRoleHandler) PutRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}

// PostRole handles POST /v1/admin/roles requests
func (arh *AdminRoleHandler) PostRole(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

}

// DeleteRole handles DELETE /v1/admin/roles/:id requests
func (arh *AdminRoleHandler) DeleteRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

}
