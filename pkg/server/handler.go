package server

import (
	"net/http"

	"github.com/tonyhhyip/seau/api"
	"github.com/tonyhhyip/seau/pkg/server/modules"
	"github.com/tonyhhyip/seau/pkg/server/repository"
)

type Handler struct {
	Store    repository.Store
	Registry *modules.Registry
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	repo, err := h.Store.GetByDomain(req.Host)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
	} else if repo == nil {
		resp.WriteHeader(http.StatusNotFound)
	} else {
		h.doHandle(repo, resp, req)
	}
}

func (h *Handler) doHandle(repo *repository.Repository, resp http.ResponseWriter, req *http.Request) {
	if !repo.AllowPublicRead && !h.checkAuth(req) {
		resp.WriteHeader(http.StatusForbidden)
		return
	}

	p, exists := h.Registry.RegisterTable.Load(repo.Handler)
	if !exists {
		resp.WriteHeader(http.StatusBadGateway)
		return
	}

	plugin := p.(api.Plugin)

	handler := plugin.Handler()
	handler.ServeHTTP(resp, req)
}

func (h *Handler) checkAuth(req *http.Request) bool {
	return req.Header.Get("Authentication") != ""
}
