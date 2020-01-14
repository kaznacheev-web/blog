package web

// func TestServer_HandleTalksGetAll(t *testing.T) {
// 	articles := []models.Article{}
//
// 	dbMock := new(mocks.StorageManager)
// 	dbMock.On("GetArticles", 1).Return(articles, nil)
//
// 	bodyWant := ``
//
// 	srv := &Server{
// 		sm: dbMock,
// 	}
//
// 	r := httptest.NewRequest("GET", "http://example.com/foo?page=1", nil)
// 	w := httptest.NewRecorder()
// 	srv.HandleArticlesGetAll()(w, r)
// 	resp := w.Result()
// 	bodyGot, _ := ioutil.ReadAll(resp.Body)
//
// 	assert.Equal(t, http.StatusOK, resp.StatusCode, "wrong status")
// 	assert.Equal(t, bodyGot, bodyWant, "wrong body")
// }
