package middleware_test

import (
	"errors"
	"myapp/internal/exception"
	"myapp/internal/interface/api/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestHandleError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success/errorなしの場合はStatusOKが返る", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.HandleError())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	t.Run("success/exceptionパッケージのエラーを投げるとそれに対応したステータスコードが返る", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.HandleError())
		router.GET("/test", func(c *gin.Context) {
			c.Error(exception.ServerError)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"code":0,"message":"エラーが発生しました"}`, w.Body.String())
	})

	t.Run("success/複数のエラーがある場合は最初のエラーのステータスコードが返る", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.HandleError())
		router.GET("/test", func(c *gin.Context) {
			c.Error(exception.ServerError)
			c.Error(exception.InvalidRequestError)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"code":0,"message":"エラーが発生しました"}`, w.Body.String())
	})

	t.Run("success/Wrapされたエラーがある場合", func(t *testing.T) {
		t.Run("Unwrapしてexceptionパッケージのエラーが見つかった場合対応するステータスコードが返る", func(t *testing.T) {
			router := gin.New()
			router.Use(middleware.HandleError())
			router.GET("/test", func(c *gin.Context) {
				c.Error(xerrors.Errorf("wrap error: %w", exception.ServerError))
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.JSONEq(t, `{"code":0,"message":"エラーが発生しました"}`, w.Body.String())
		})

		t.Run("Unwrapしてもexceptionパッケージのエラーが見つからない場合は500エラーが返る", func(t *testing.T) {
			router := gin.New()
			router.Use(middleware.HandleError())
			router.GET("/test", func(c *gin.Context) {
				c.Error(xerrors.Errorf("wrap error: %w", errors.New("error")))
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.JSONEq(t, `{"code":0,"message":"エラーが発生しました"}`, w.Body.String())
		})
	})
}
