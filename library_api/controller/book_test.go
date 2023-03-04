package controller_test

import (
	"bytes"
	"errors"
	"library_api/controller"
	"library_api/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCreateBooks(t *testing.T) {
	url := "/create_book"
	jsonBody := []byte(`{
		"name":"Test name",
		"author":"Test Author",
		"year":"2000"
		}`)

	t.Run("success create book", func(t *testing.T) {
		bodyReader := bytes.NewBuffer(jsonBody)
		req, err := http.NewRequest(http.MethodPost, "/create_book", bodyReader)
		if err != nil {
			t.Fatal(err)
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBookRepository := mock.NewMockBookRepository(mockCtrl)
		mockBookRepository.EXPECT().Create(mock.BookForCreate).Return(1, nil)

		c := &controller.Book{Book_repository: mockBookRepository}
		r := *mux.NewRouter()
		rr := httptest.NewRecorder()
		r.HandleFunc(url, c.CreateBook).Methods(http.MethodPost)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Empty(t, rr.Body.String())
	})

	t.Run("failed book upload", func(t *testing.T) {
		bodyReader := bytes.NewBuffer(jsonBody)
		req, err := http.NewRequest(http.MethodPost, "/create_book", bodyReader)
		if err != nil {
			t.Fatal(err)
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBookRepository := mock.NewMockBookRepository(mockCtrl)
		mockBookRepository.EXPECT().Create(mock.BookForCreate).Return(1, errors.New("book upload error"))

		c := &controller.Book{Book_repository: mockBookRepository}
		r := *mux.NewRouter()
		rr := httptest.NewRecorder()
		r.HandleFunc(url, c.CreateBook).Methods(http.MethodPost)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, ``, rr.Body.String())
	})
}

func TestListOfBooks(t *testing.T) {
	url := "/books"
	jsonBody := []byte(`[{"name":"Test name","author":"Test Author","year":"2000"}]`)

	t.Run("success show list of books", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/books", nil)
		if err != nil {
			t.Fatal(err)
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBookRepository := mock.NewMockBookRepository(mockCtrl)
		mockBookRepository.EXPECT().List().Return(mock.BooksForShowing, nil)

		c := &controller.Book{Book_repository: mockBookRepository}
		r := *mux.NewRouter()
		rr := httptest.NewRecorder()
		r.HandleFunc(url, c.List).Methods(http.MethodGet)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, jsonBody, rr.Body.Bytes())
	})

	t.Run("failed show list of books", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/books", nil)
		if err != nil {
			t.Fatal(err)
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBookRepository := mock.NewMockBookRepository(mockCtrl)
		mockBookRepository.EXPECT().List().Return(nil, errors.New("book upload error"))

		c := &controller.Book{Book_repository: mockBookRepository}
		r := *mux.NewRouter()
		rr := httptest.NewRecorder()
		r.HandleFunc(url, c.List).Methods(http.MethodGet)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Empty(t, rr.Body.String())
	})
}
