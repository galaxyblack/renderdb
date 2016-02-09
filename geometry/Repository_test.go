package geometry

import (
	"errors"
	"testing"

	"github.com/dhconnelly/rtreego"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ungerik/go3d/float64/vec3"
)

type mockDatabase struct {
	mock.Mock
}

func (m *mockDatabase) add(o Object) (int64, error) {
	args := m.Called(o)
	return int64(args.Int(0)), args.Error(1)
}

func (m *mockDatabase) getMany(ids []int64) (<-chan *data, <-chan error) {
	args := m.Called(ids)
	return args.Get(0).(chan *data), args.Get(1).(chan error)
}

func (m *mockDatabase) getAll() (<-chan *data, <-chan error) {
	args := m.Called()
	return args.Get(0).(chan *data), args.Get(1).(chan error)
}

type mockFilterOptions struct {
	mock.Mock
}

func (m *mockFilterOptions) Apply(bounds []*vec3.Box) []int {
	args := m.Called(bounds)
	return args.Get(0).([]int)
}

// createGetManyResult is a helper function to emulate how getMany() returns
// values/error. Usage:
//   .Returns(createGetManyResult(data1, data2, err)) // Returns two values, then fails
func createGetManyResult(values ...interface{}) (chan *data, chan error) {
	dataCh := make(chan *data)
	errCh := make(chan error)
	go func() {
		defer close(dataCh)
		for _, v := range values {
			if data, ok := v.(*data); ok {
				dataCh <- data
			} else if err, ok := v.(error); ok {
				errCh <- err
				return
			} else {
				panic("Values must be error or *data")
			}
		}
	}()
	return dataCh, errCh
}

func flattenGetInsideVolumeResults(objCh <-chan Object, errCh <-chan error) ([]Object, error) {
	objects := []Object{}
	doBreak := false
	for !doBreak {
		select {
		case obj, more := <-objCh:
			if more {
				objects = append(objects, obj)
			} else {
				doBreak = true
			}
		case err := <-errCh:
			return objects, err
		}
	}
	return objects, nil
}

func TestRepository_Add_ValidGeometry_AddsToTreeAndDatabase(t *testing.T) {
	// Arrange
	obj := new(SimpleObject)
	obj.bounds = &vec3.Box{}

	mockDb := new(mockDatabase)
	mockDb.On("add", obj).Return(1, nil)

	rtree := rtreego.NewTree(3, 5, 10)
	repo := defaultRepository{mockDb, rtree}

	// Act
	id, err := repo.Add(obj)

	// Assert
	assert.Equal(t, int64(1), id)
	assert.Nil(t, err)
	assert.Equal(t, 1, rtree.Size())
	mockDb.AssertExpectations(t)
}

func TestRepository_Add_DatabaseReturnsError_DoesNotAddToTree(t *testing.T) {
	// Arrange
	obj := new(SimpleObject)
	obj.bounds = &vec3.Box{}

	mockDb := new(mockDatabase)
	mockDb.On("add", obj).Return(0, errors.New("error"))

	rtree := rtreego.NewTree(3, 5, 10)
	repo := defaultRepository{mockDb, rtree}

	// Act
	_, err := repo.Add(obj)

	// Assert
	assert.Equal(t, 0, rtree.Size())
	assert.NotNil(t, err)
}

func TestRepository_GetInsideVolume_NothingInsideVolume_ReturnsEmpty(t *testing.T) {
	// Arrange
	objBounds := vec3.Box{vec3.T{1, 1, 1}, vec3.T{2, 2, 2}}
	obj := &SimpleObject{
		bounds:       &objBounds,
		geometryData: []byte{},
		metadata:     nil,
	}

	mockDb := new(mockDatabase)
	mockDb.On("add", obj).Return(1, nil)
	mockDb.On("getMany", []int64{}).Return(createGetManyResult())

	rtree := rtreego.NewTree(3, 5, 10)
	repo := defaultRepository{mockDb, rtree}
	repo.Add(obj)

	// Act
	bounds := vec3.Box{vec3.T{5, 5, 5}, vec3.T{6, 6, 6}}
	objects, err := flattenGetInsideVolumeResults(repo.GetInsideVolume(bounds))

	// Assert
	mockDb.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Empty(t, objects)
}

func TestRepository_GetInsideVolume_OneInsideVolume_ReturnsObject(t *testing.T) {
	// Arrange
	objBounds := vec3.Box{vec3.T{0.5, 0.5, 0.5}, vec3.T{1.5, 1.5, 1.5}}
	obj := &SimpleObject{
		bounds:       &objBounds,
		geometryData: []byte{},
		metadata:     nil,
	}

	data := new(data)
	data.id = 1
	mockDb := new(mockDatabase)
	mockDb.On("add", obj).Return(1, nil)
	mockDb.On("getMany", []int64{1}).Return(createGetManyResult(data))
	rtree := rtreego.NewTree(3, 5, 10)
	repo := defaultRepository{mockDb, rtree}
	repo.Add(obj)

	// Act
	searchBounds := vec3.Box{vec3.T{0.5, 0.5, 0.5}, vec3.T{1.5, 1.5, 1.5}}
	objects, err := flattenGetInsideVolumeResults(repo.GetInsideVolume(searchBounds))

	// Assert
	mockDb.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(objects))
}

func TestRepository_GetInsideVolume_DatabaseReturnsError_ReturnsError(t *testing.T) {
	// Arrange
	objBounds := vec3.Box{vec3.T{0.5, 0.5, 0.5}, vec3.T{1.5, 1.5, 1.5}}
	obj := &SimpleObject{
		bounds:       &objBounds,
		geometryData: []byte{},
		metadata:     nil,
	}

	mockDb := new(mockDatabase)
	mockDb.On("add", obj).Return(1, nil)
	mockDb.On("getMany", []int64{1}).Return(createGetManyResult(errors.New("error")))
	rtree := rtreego.NewTree(3, 5, 10)
	repo := defaultRepository{mockDb, rtree}
	repo.Add(obj)

	// Act
	searchBounds := vec3.Box{vec3.T{0.5, 0.5, 0.5}, vec3.T{1.5, 1.5, 1.5}}
	objects, err := flattenGetInsideVolumeResults(repo.GetInsideVolume(searchBounds))

	// Assert
	mockDb.AssertExpectations(t)
	assert.Error(t, err)
	assert.Equal(t, 0, len(objects))
}

func TestRepository_GetInsideVolume_DatabaseReturnsOneThenError_ReturnsError(t *testing.T) {
	// Arrange
	objBounds := vec3.Box{vec3.T{0.5, 0.5, 0.5}, vec3.T{1.5, 1.5, 1.5}}
	obj1 := &SimpleObject{
		bounds:       &objBounds,
		geometryData: []byte("1"),
		metadata:     nil,
	}
	obj2 := &SimpleObject{
		bounds:       &objBounds,
		geometryData: []byte("2"),
		metadata:     nil,
	}

	data := new(data)
	data.id = 1
	mockDb := new(mockDatabase)
	mockDb.On("add", obj1).Return(1, nil)
	mockDb.On("add", obj2).Return(2, nil)
	mockDb.On("getMany", []int64{1, 2}).Return(createGetManyResult(data, errors.New("error")))
	rtree := rtreego.NewTree(3, 5, 10)
	repo := defaultRepository{mockDb, rtree}
	repo.Add(obj1)
	repo.Add(obj2)

	// Act
	searchBounds := vec3.Box{vec3.T{0.5, 0.5, 0.5}, vec3.T{1.5, 1.5, 1.5}}
	_, err := flattenGetInsideVolumeResults(repo.GetInsideVolume(searchBounds))

	// Assert
	mockDb.AssertExpectations(t)
	assert.Error(t, err)
}

func TestRepository_GetInsideVolume_WithFilterGeometryOptions_ReturnsFiltered(t *testing.T) {
	// Arrange
	objBounds := vec3.Box{vec3.T{0.5, 0.5, 0.5}, vec3.T{1.5, 1.5, 1.5}}
	obj1 := &SimpleObject{
		bounds:       &objBounds,
		geometryData: []byte("1"),
		metadata:     nil,
	}
	obj2 := &SimpleObject{
		bounds:       &objBounds,
		geometryData: []byte("2"),
		metadata:     nil,
	}

	data := new(data)
	data.id = 1
	mockDb := new(mockDatabase)
	mockDb.On("add", obj1).Return(1, nil)
	mockDb.On("add", obj2).Return(2, nil)
	mockDb.On("getMany", []int64{1}).Return(createGetManyResult(data))
	rtree := rtreego.NewTree(3, 5, 10)
	repo := defaultRepository{mockDb, rtree}
	repo.Add(obj1)
	repo.Add(obj2)

	mockOptions := new(mockFilterOptions)
	mockOptions.On("Apply", []*vec3.Box{&objBounds, &objBounds}).Return([]int{0})

	// Act
	searchBounds := vec3.Box{vec3.T{0.5, 0.5, 0.5}, vec3.T{1.5, 1.5, 1.5}}
	result, err := flattenGetInsideVolumeResults(repo.GetInsideVolume(searchBounds, mockOptions))

	// Assert
	mockDb.AssertExpectations(t)
	mockOptions.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result))
}

func TestRepository_LoadFromDatabase_LoadsObjectsFromDatabase(t *testing.T) {
	// Arrange
	mockDb := new(mockDatabase)
	mockDb.On("getAll").Return(createGetManyResult(&data{id: 1}, &data{id: 2}))
	rtree := rtreego.NewTree(3, 5, 10)
	repo := defaultRepository{mockDb, rtree}

	// Act
	err := repo.loadFromDatabase()

	// Assert
	mockDb.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, 2, rtree.Size())
}
