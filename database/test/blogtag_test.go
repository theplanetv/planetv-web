package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Connection_BlogTag(t *testing.T) {
	service := BlogTagService{}

	t.Run("Database connection success", func(t *testing.T) {
		// Execute
		err := service.Open()
		defer service.Close()

		// Assert check error
		assert.NoError(t, err)
	})
}

func Test_Count_BlogTag(t *testing.T) {
	service := BlogTagService{}

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

func Test_GetFirst_BlogTag(t *testing.T) {
	service := BlogTagService{}

	t.Run("Get first value success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		item, err := service.GetFirst()

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, item, BlogTag{})
		assert.NotEmpty(t, item)
	})
}

func Test_GetLast_BlogTag(t *testing.T) {
	service := BlogTagService{}

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
		assert.IsType(t, item, BlogTag{})
		assert.NotEmpty(t, item)
		assert.GreaterOrEqual(t, item.Id, count)
	})
}

func Test_GetAll_BlogTag(t *testing.T) {
	service := BlogTagService{}

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
		assert.IsType(t, []BlogTag{}, data)
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
		assert.IsType(t, []BlogTag{}, data)
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
		assert.IsType(t, []BlogTag{}, data)
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
		assert.IsType(t, []BlogTag{}, data)
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
		assert.IsType(t, []BlogTag{}, data)
		assert.NotEmpty(t, data)
		assert.Contains(t, data, last)
	})
}

func Test_Create_BlogTag(t *testing.T) {
	service := BlogTagService{}

	t.Run("Create success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get count before creating
		beforeCount, _ := service.Count()

		// Execute testing method
		item := BlogTag{
			BlogcategoryId: 1,
			Name:           "example tag",
		}
		err := service.Create(&item)

		// Get count after creating
		afterCount, _ := service.Count()

		// Assert check error
		assert.NoError(t, err)
		assert.Less(t, beforeCount, afterCount)
	})
}

func Test_Update_BlogTag(t *testing.T) {
	service := BlogTagService{}

	t.Run("Update unsuccess without id", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		item := BlogTag{Id: 0}
		err := service.Update(&item)

		// Assert check error
		assert.Error(t, err)
	})

	t.Run("Update unsuccess without BlogCategory id and name", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTag{
			Id: lastId,
		}
		err := service.Update(&item)

		// Assert check error
		assert.Error(t, err)
	})

	t.Run("Update success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTag{
			Id:             lastId,
			BlogcategoryId: 2,
			Name:           "this is example tag",
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})

	t.Run("Update success without BlogCategory id", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTag{
			Id:   lastId,
			Name: "this is an example tag",
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})

	t.Run("Update success without name", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogTag{
			Id:             lastId,
			BlogcategoryId: 3,
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})
}

func Test_Remove_BlogTag(t *testing.T) {
	service := BlogTagService{}

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
