package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Connection_BlogTagFile(t *testing.T) {
	service := BlogTagFileService{}

	t.Run("Database connection success", func(t *testing.T) {
		// Execute
		err := service.Open()
		defer service.Close()

		// Assert check error
		assert.NoError(t, err)
	})
}

func Test_Count_BlogTagFile(t *testing.T) {
	service := BlogTagFileService{}

	t.Run("Count success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		count, err := service.Count()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEmpty(t, count)
	})
}

func Test_GetFirst_BlogTagFile(t *testing.T) {
	service := BlogTagFileService{}

	t.Run("Get first value success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		item, err := service.GetFirst()

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, item, BlogTagFile{})
		assert.NotEmpty(t, item)
	})
}

func Test_GetLast_BlogTagFile(t *testing.T) {
	service := BlogTagFileService{}

	t.Run("Get last value success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get count
		count, _ := service.Count()

		// Execute testing method
		item, err := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, item, BlogTagFile{})
		assert.NotEmpty(t, item)
		assert.GreaterOrEqual(t, item.Id, count)
	})
}

func Test_GetAll_BlogTagFile(t *testing.T) {
	service := BlogTagFileService{}

	t.Run("Get all with limit and page success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		limit := 10
		page := 1
		data, err := service.GetAll(&limit, &page)

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, []BlogTagFile{}, data)
		assert.NotEmpty(t, data)
	})

	t.Run("Get all with limit < 5", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		limit := 4
		page := 1
		data, err := service.GetAll(&limit, &page)

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, []BlogTagFile{}, data)
		assert.NotEmpty(t, data)
		assert.Equal(t, 5, len(data))
	})

	t.Run("Get all with limit > 50", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		limit := 51
		page := 1
		data, err := service.GetAll(&limit, &page)

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, []BlogTagFile{}, data)
		assert.NotEmpty(t, data)
		assert.Equal(t, 50, len(data))
	})

	t.Run("Get all with page < 1", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get first
		first, _ := service.GetFirst()

		// Execute testing method
		limit := 10
		page := 0
		data, err := service.GetAll(&limit, &page)

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, []BlogTagFile{}, data)
		assert.NotEmpty(t, data)
		assert.Contains(t, data, first)
	})

	t.Run("Get all with page > ceil (count / page)", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get count and last
		count, _ := service.Count()
		last, _ := service.GetLast()

		// Execute testing method
		limit := 10
		page := count
		data, err := service.GetAll(&limit, &page)

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, []BlogTagFile{}, data)
		assert.NotEmpty(t, data)
		assert.Contains(t, data, last)
	})
}

func Test_Create_BlogTagFile(t *testing.T) {
	service := BlogTagFileService{}

	t.Run("Create unsuccess if same value exists", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get count before creating
		beforeCount, _ := service.Count()

		// Execute testing method
		item := BlogTagFile{
			BlogtagId:  10,
			BlogfileId: 1,
		}
		err := service.Create(&item)

		// Get count after creating
		afterCount, _ := service.Count()

		// Assert check error
		assert.Error(t, err)
		assert.Equal(t, beforeCount, afterCount)
	})

	t.Run("Create success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get count before creating
		beforeCount, _ := service.Count()

		// Execute testing method
		item := BlogTagFile{
			BlogtagId:  11,
			BlogfileId: 1,
		}
		err := service.Create(&item)

		// Get count after creating
		afterCount, _ := service.Count()

		// Assert check error
		assert.NoError(t, err)
		assert.Less(t, beforeCount, afterCount)
	})
}

func Test_Update_BlogTagFile(t *testing.T) {
	service := BlogTagFileService{}

	t.Run("Update unsuccess without id", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		item := BlogTagFile{Id: 0}
		err := service.Update(&item)

		// Assert check error
		assert.Error(t, err)
	})

	t.Run("Update unsuccess without BlogTag id and BlogFile id", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTagFile{
			Id: lastId,
		}
		err := service.Update(&item)

		// Assert check error
		assert.Error(t, err)
	})

	t.Run("Update unsuccess if the same value exists", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTagFile{
			Id:         lastId,
			BlogtagId:  10,
			BlogfileId: 1,
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.Error(t, err)
		assert.Equal(t, beforeValue, afterValue)
	})

	t.Run("Update success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTagFile{
			Id:         lastId,
			BlogtagId:  11,
			BlogfileId: 2,
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})

	t.Run("Update success without BlogTag id", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTagFile{
			Id:         lastId,
			BlogfileId: 3,
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})

	t.Run("Update success without BlogFile id", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTagFile{
			Id:        lastId,
			BlogtagId: 12,
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})
}

func Test_Remove_BlogTagFile(t *testing.T) {
	service := BlogTagFileService{}

	t.Run("Remove success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get count before removing
		beforeCount, _ := service.Count()

		// Execute testing method
		lastValue, _ := service.GetLast()
		err := service.Remove(&lastValue.Id)

		// Get count after removing
		afterCount, _ := service.Count()

		// Assert check error
		assert.NoError(t, err)
		assert.Greater(t, beforeCount, afterCount)
	})
}
