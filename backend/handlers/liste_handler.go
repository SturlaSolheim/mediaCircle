package handlers

import( 
	"net/http"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/liste"
	"github.com/SturlaSolheim/mediaCircleBackend/service"
)


type ListeHandlerImpl struct {
	listeService service.ListeService
}

func NewOpenAPIListeHandler(
	listeService service.ListeService,
) liste.ServerInterface {
	return &ListeHandlerImpl{
		listeService: listeService,
	}
}
func (h *ListeHandlerImpl) GetListe(w http.ResponseWriter, r *http.Request, id int) {
	
	liste, err := h.listeService.GetListe()
	if err != nil {
		WriteInternalServerError(w, "Internal server error")
		return
	}
	WriteJSON(w, http.StatusOK, liste)
}

func (h *ListeHandlerImpl) GetListeTest(w http.ResponseWriter, r *http.Request, id int) {
	
	liste, err := h.listeService.GetListe()
	if err != nil {
		WriteInternalServerError(w, "Internal server error")
		return
	}
	WriteJSON(w, http.StatusOK, liste)
}

