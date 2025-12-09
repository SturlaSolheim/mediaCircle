package handlers

import( 
	"net/http"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/anmeldelse"
	"github.com/SturlaSolheim/mediaCircleBackend/service"
)


type AnmeldelseHandlerImpl struct {
	anmeldelseService service.AnmeldelseService
}

func NewOpenAPIAnmeldelseHandler(
	anmeldelseService service.AnmeldelseService,
) anmeldelse.ServerInterface {
	return &AnmeldelseHandlerImpl{
		anmeldelseService: anmeldelseService,
	}
}
func (h *AnmeldelseHandlerImpl) GetAlleAnmeldelserForBruker(w http.ResponseWriter, r *http.Request, bruker string) {
	
	anmeldelser, err := h.anmeldelseService.GetAlleAnmeldelserForEnBruker()
	if err != nil {
		WriteInternalServerError(w, "Internal server error")
		return
	}
	WriteJSON(w, http.StatusOK, anmeldelser)
}


