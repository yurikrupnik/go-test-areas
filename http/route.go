package http

import (
	"github.com/go-chi/chi"
	"go-test-areas/domain"
)

func Routes(r chi.Router, h *Handler) {
	r.Route("/pets", func(r chi.Router) {
		//r.Get("/", h.ListPets)
		r.With(ValidateQueryParam(ListPetsQuery{})).Get("/", h.ListPets)
		r.With(ValidateBody(domain.Pet{})).Post("/", h.AddPet)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(ValidateURLParam("id"))
			r.Get("/", h.GetPet)
			r.Delete("/", h.DeletePet)
		})
	})
}
