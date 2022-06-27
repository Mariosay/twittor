package middlew

import (
	"net/http"

	"github.com/Mariosay/twittor/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "conexion perdidad con la base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
