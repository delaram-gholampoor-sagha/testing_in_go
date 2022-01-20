package httpservice_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Delaram-Gholampoor-Sagha/testing_in_go/httpservice"
	"github.com/stretchr/testify/assert"
)




func TestHomePage(t *testing.T) {
     req := httptest.NewRequest("GET" , "/" , nil)
	 rr := httptest.NewRecorder()
	 handler := http.HandlerFunc(httpservice.HomePage)
	 handler.ServeHTTP(rr , req)
	 assert.Equal(t , http.StatusOK , rr.Code)
}