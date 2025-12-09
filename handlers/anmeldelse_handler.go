package handlers

import( 
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json")
	
	anmeldelser, err := h.anmeldelseService.GetAlleAnmeldelserForEnBruker()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(anmeldelser)
}


