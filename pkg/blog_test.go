package pkg

import "testing"

func TestSaveArticle(t *testing.T) {

	blog := New()

	blog.SaveArticle(Article{"My title", "My Post Body"})

	if blog.Articles[0].Title != "My title" {
		t.Errorf("Item was not added")
	}
}

func TestFetchAllArticles(t *testing.T) {

	blog := New()

	blog.SaveArticle(Article{"My title", "My Post Body"})

	articles := blog.FetchAll()

	if len(articles) == 0 {
		t.Errorf("Fetch All fails")
	}
}