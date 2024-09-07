package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Connection_BlogFile(t *testing.T) {
	service := BlogFileService{}

	t.Run("Database connection success", func(t *testing.T) {
		// Execute
		err := service.Open()
		defer service.Close()

		// Assert check error
		assert.NoError(t, err)
	})
}

func Test_Count_BlogFile(t *testing.T) {
	service := BlogFileService{}

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

func Test_GetFirst_BlogFile(t *testing.T) {
	service := BlogFileService{}

	t.Run("Get first value success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		item, err := service.GetFirst()

		// Assert check error
		assert.NoError(t, err)
		assert.IsType(t, item, BlogFile{})
		assert.NotEmpty(t, item)
	})
}

func Test_GetLast_BlogFile(t *testing.T) {
	service := BlogFileService{}

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
		assert.IsType(t, item, BlogFile{})
		assert.NotEmpty(t, item)
		assert.GreaterOrEqual(t, item.Id, count)
	})
}

func Test_GetAll_BlogFile(t *testing.T) {
	service := BlogFileService{}

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
		assert.IsType(t, []BlogFile{}, data)
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
		assert.IsType(t, []BlogFile{}, data)
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
		assert.IsType(t, []BlogFile{}, data)
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
		assert.IsType(t, []BlogFile{}, data)
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
		assert.IsType(t, []BlogFile{}, data)
		assert.NotEmpty(t, data)
		assert.Contains(t, data, last)
	})
}

func Test_Create_BlogFile(t *testing.T) {
	service := BlogFileService{}

	t.Run("Create success", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get count before creating
		beforeCount, _ := service.Count()

		// Execute testing method
		item := BlogFile{
			Filename: "example.md",
		}
		err := service.Create(&item)

		// Get count after creating
		afterCount, _ := service.Count()

		// Assert check error
		assert.NoError(t, err)
		assert.Less(t, beforeCount, afterCount)
	})
}

func Test_Update_BlogFile(t *testing.T) {
	service := BlogFileService{}

	t.Run("Update unsuccess without id", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Execute testing method
		item := BlogFile{Id: 0}
		err := service.Update(&item)

		// Assert check error
		assert.Error(t, err)
	})

	t.Run("Update unsuccess without Filename, Created date and Updated date", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogFile{
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
		currentTime := time.Time{}
		item := BlogFile{
			Id:          lastId,
			Filename:    "blogfile-example.md",
			CreatedDate: currentTime.AddDate(0, 0, -1),
			UpdatedDate: currentTime.AddDate(0, 0, 1),
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})

	t.Run("Update success with Filename", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		item := BlogFile{
			Id:       lastId,
			Filename: "example-blogfile.md",
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})

	t.Run("Update success with Created date", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		currentTime := time.Time{}
		item := BlogFile{
			Id:          lastId,
			CreatedDate: currentTime.AddDate(0, 0, -2),
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})

	t.Run("Update success with Updated date", func(t *testing.T) {
		// Open database connection and close after end
		_ = service.Open()
		defer service.Close()

		// Get value before updating
		beforeValue, _ := service.GetLast()

		// Execute testing method
		lastId := beforeValue.Id
		currentTime := time.Time{}
		item := BlogFile{
			Id:          lastId,
			UpdatedDate: currentTime.AddDate(0, 0, 2),
		}
		err := service.Update(&item)

		// Get value after updating
		afterValue, _ := service.GetLast()

		// Assert check error
		assert.NoError(t, err)
		assert.NotEqual(t, beforeValue, afterValue)
	})
}

func Test_Remove_BlogFile(t *testing.T) {
	service := BlogFileService{}

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
