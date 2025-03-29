package link

import (
	"fmt"
	apperror "link-shortener/pkg/app-error"
	"link-shortener/pkg/request"
	"link-shortener/pkg/response"
	"net/http"
)

type Handler struct {
	LinkHandlerDeps
}

type LinkHandlerDeps struct {
	Router  *http.ServeMux
	Service *Service
}

func RegisterLinkHandler(linkHandlerDeps LinkHandlerDeps) {
	handler := &Handler{
		linkHandlerDeps,
	}
	handler.Router.HandleFunc("POST /link", handler.create)
	handler.Router.HandleFunc("GET /{hash}", handler.goTo)
	handler.Router.HandleFunc("DELETE /link", handler.delete)

}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	body, err := request.GetBody[CreateRequestPayload](w, r)
	if err != nil {
		response.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	fmt.Println(body)

	createdLink, err := h.Service.Create(body.Url)
	if err != nil {
		h.handleAppError(err, w)
		return
	}
	response.Json(w, 200, &createdLink)
}

func (h *Handler) goTo(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")
	link, err := h.Service.Find(&FindParams{
		hash: hash,
	})
	if err != nil {
		h.handleAppError(err, w)
		return
	}
	http.Redirect(w, r, link.Url, http.StatusSeeOther)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	body, err := request.GetBody[DeleteRequestPayload](w, r)
	if err != nil {
		response.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = h.Service.Delete(&FindParams{
		hash: body.Hash,
		url:  body.Url,
		id:   body.Id,
	})
	if err != nil {
		h.handleAppError(err, w)
		return
	}

	response.Json(w, 200, &map[string]string{"message": "Link deleted"})
}

func (h *Handler) handleAppError(err error, w http.ResponseWriter) {
	if appErr, ok := err.(*apperror.AppError); ok {
		http.Error(w, appErr.Message, appErr.Code)
	} else {
		http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
	}
}
