package plant

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindAll() ([]PlantCategory, error) {
	args := m.Called()
	return args.Get(0).([]PlantCategory), args.Error(1)
}

func (m *MockRepository) FindByID(id int) (PlantCategory, error) {
	args := m.Called(id)
	return args.Get(0).(PlantCategory), args.Error(1)
}

func (m *MockRepository) Create(category PlantCategory) (PlantCategory, error) {
	args := m.Called(category)
	return args.Get(0).(PlantCategory), args.Error(1)
}

func (m *MockRepository) Update(category PlantCategory) (PlantCategory, error) {
	args := m.Called(category)
	return args.Get(0).(PlantCategory), args.Error(1)
}

func (m *MockRepository) Delete(category PlantCategory) error {
	args := m.Called(category)
	return args.Error(0)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewPlantCategoryService(mockRepo)

	mockCategories := []PlantCategory{
		{ID: 1, Name: "Category A"},
		{ID: 2, Name: "Category B"},
	}

	mockRepo.On("FindAll").Return(mockCategories, nil)

	result, err := service.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, len(mockCategories), len(result))
}

func TestFindByID(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewPlantCategoryService(mockRepo)

	mockCategory := PlantCategory{ID: 1, Name: "Category A"}

	mockRepo.On("FindByID", mock.AnythingOfType("int")).Return(mockCategory, nil)

	result, err := service.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockCategory.ID, result.ID)
	assert.Equal(t, mockCategory.Name, result.Name)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewPlantCategoryService(mockRepo)

	mockInput := PlantCategoryClimateInput{Name: "Category A"}
	mockImageURL := "http://example.com/image.jpg"

	mockRepo.On("Create", mock.Anything).Return(PlantCategory{ID: 1, Name: mockInput.Name, ImageURL: mockImageURL}, nil)

	result, err := service.Create(mockInput, mockImageURL)

	assert.NoError(t, err)
	assert.Equal(t, mockInput.Name, result.Name)
	assert.Equal(t, mockImageURL, result.ImageURL)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewPlantCategoryService(mockRepo)

	mockID := 1
	mockInput := PlantCategoryClimateInput{Name: "Updated Category A"}
	mockImageURL := "http://example.com/updated-image.jpg"

	mockRepo.On("FindByID", mock.AnythingOfType("int")).Return(PlantCategory{ID: mockID, Name: "Category A"}, nil)
	mockRepo.On("Update", mock.Anything).Return(PlantCategory{ID: mockID, Name: mockInput.Name, ImageURL: mockImageURL}, nil)

	result, err := service.Update(mockID, mockInput, mockImageURL)

	assert.NoError(t, err)
	assert.Equal(t, mockInput.Name, result.Name)
	assert.Equal(t, mockImageURL, result.ImageURL)
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewPlantCategoryService(mockRepo)

	mockID := 1

	mockRepo.On("FindByID", mock.AnythingOfType("int")).Return(PlantCategory{ID: mockID, Name: "Category A"}, nil)
	mockRepo.On("Delete", mock.Anything).Return(nil)

	err := service.Delete(mockID)

	assert.NoError(t, err)
}
