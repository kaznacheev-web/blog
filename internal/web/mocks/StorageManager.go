// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/kaznacheev-web/blog/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// StorageManager is an autogenerated mock type for the StorageManager type
type StorageManager struct {
	mock.Mock
}

// GetArticle provides a mock function with given fields: ctx, slug
func (_m *StorageManager) GetArticle(ctx context.Context, slug string) (*models.Article, error) {
	ret := _m.Called(ctx, slug)

	var r0 *models.Article
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Article); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticleCount provides a mock function with given fields: ctx
func (_m *StorageManager) GetArticleCount(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticles provides a mock function with given fields: ctx, page
func (_m *StorageManager) GetArticles(ctx context.Context, page int) ([]models.Article, error) {
	ret := _m.Called(ctx, page)

	var r0 []models.Article
	if rf, ok := ret.Get(0).(func(context.Context, int) []models.Article); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSimplePage provides a mock function with given fields: ctx, slug
func (_m *StorageManager) GetSimplePage(ctx context.Context, slug string) (*models.SimplePage, error) {
	ret := _m.Called(ctx, slug)

	var r0 *models.SimplePage
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.SimplePage); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SimplePage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTalk provides a mock function with given fields: ctx, slug
func (_m *StorageManager) GetTalk(ctx context.Context, slug string) (*models.Talk, error) {
	ret := _m.Called(ctx, slug)

	var r0 *models.Talk
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Talk); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Talk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTalkCount provides a mock function with given fields: ctx
func (_m *StorageManager) GetTalkCount(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTalks provides a mock function with given fields: ctx, page
func (_m *StorageManager) GetTalks(ctx context.Context, page int) ([]models.Talk, error) {
	ret := _m.Called(ctx, page)

	var r0 []models.Talk
	if rf, ok := ret.Get(0).(func(context.Context, int) []models.Talk); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Talk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}