package mock

import (
	"library_api/entity"
	"reflect"

	"github.com/golang/mock/gomock"
)

var BookForCreate = entity.Book{
	Name:   "Test name",
	Author: "Test Author",
	Year:   2000,
}


var BooksForShowing = []entity.Book{
	{
	Name:   "Test name",
	Author: "Test Author",
	Year:   2000,
	},
}

// MockReader is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryRecorder
}

// MockBookRepository is the mock recorder for MockBookRepository.
type MockBookRepositoryRecorder struct {
	mock *MockBookRepository
}

// MockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockBookRepository) Create(book entity.Book) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", book)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBookRepositoryRecorder) Create(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookRepository)(nil).Create), book)
}

// GetBooks mocks base method.
func (m *MockBookRepository) List() ([]entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooks indicates an expected call of GetBooks.
func (mr *MockBookRepositoryRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBookRepository)(nil).List))
}
