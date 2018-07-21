package view_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/novakit/nova"
	"github.com/novakit/router"
	"github.com/novakit/testkit"
	"github.com/novakit/view"
)

func sanitizeHTML(s string) (r string) {
	s = strings.TrimSpace(s)
	rn := strings.Split(s, "\n")
	for _, n := range rn {
		n = strings.TrimSpace(n)
		r += n
	}
	return
}

func TestView_Binary(t *testing.T) {
	n := nova.New()
	n.Use(view.Handler(view.Options{
		Directory: "testdata",
	}))
	router.Route(n).Get("/hello").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Binary([]byte("hello"))
		return nil
	})
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	res := testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "application/octet-stream" || res.String() != "hello" {
		t.Error("failed")
	}
}

func TestView_Text(t *testing.T) {
	n := nova.New()
	n.Use(view.Handler(view.Options{
		Directory: "testdata",
	}))
	router.Route(n).Get("/hello").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Text("hello")
		return nil
	})
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	res := testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "text/plain" || res.String() != "hello" {
		t.Error("failed")
	}
}

func TestView_JSON(t *testing.T) {
	n := nova.New()
	n.Use(view.Handler(view.Options{
		Directory: "testdata",
	}))
	router.Route(n).Get("/hello").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.JSON(map[string]string{"A": "B"})
		return nil
	})
	router.Route(n).Get("/hello2").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Data["A"] = "B"
		v.DataAsJSON()
		return nil
	})

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	res := testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "application/json" || res.String() != `{"A":"B"}` {
		t.Error("failed")
	}

	req, _ = http.NewRequest(http.MethodGet, "/hello2", nil)
	res = testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "application/json" || res.String() != `{"A":"B"}` {
		t.Error("failed")
	}
}

func TestView_HTML(t *testing.T) {
	n := nova.New()
	n.Use(view.Handler(view.Options{
		Directory: "testdata",
	}))
	router.Route(n).Get("/hello1").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Data["Key1"] = "AAA"
		v.HTML("dir1/dir11")
		return nil
	})
	router.Route(n).Get("/hello2").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Data["Key1"] = "AAA"
		v.Data["Key2"] = "BBB"
		v.HTML("dir2/dir22")
		return nil
	})
	router.Route(n).Get("/hello3").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Data["Key1"] = "AAA"
		v.Data["Key2"] = "BBB"
		v.Data["Key3"] = "CCC"
		v.HTML("dir2/dir21/dir211")
		return nil
	})
	req, _ := http.NewRequest(http.MethodGet, "/hello1", nil)
	res := testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "text/html" || sanitizeHTML(res.String()) != "AAA" {
		t.Error("failed 1", sanitizeHTML(res.String()))
	}
	req, _ = http.NewRequest(http.MethodGet, "/hello2", nil)
	res = testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "text/html" || sanitizeHTML(res.String()) != "AAABBB" {
		t.Error("failed 2", sanitizeHTML(res.String()))
	}
	req, _ = http.NewRequest(http.MethodGet, "/hello3", nil)
	res = testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "text/html" || sanitizeHTML(res.String()) != "AAABBBCCC" {
		t.Error("failed 3", sanitizeHTML(res.String()))
	}
}

func TestView_HTMLBinFS(t *testing.T) {
	n := nova.New()
	n.Use(view.Handler(view.Options{
		Directory: "testdata",
		BinFS:     true,
	}))
	router.Route(n).Get("/hello1").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Data["Key1"] = "AAA"
		v.HTML("dir1/dir11")
		return nil
	})
	router.Route(n).Get("/hello2").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Data["Key1"] = "AAA"
		v.Data["Key2"] = "BBB"
		v.HTML("dir2/dir22")
		return nil
	})
	router.Route(n).Get("/hello3").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Data["Key1"] = "AAA"
		v.Data["Key2"] = "BBB"
		v.Data["Key3"] = "CCC"
		v.HTML("dir2/dir21/dir211")
		return nil
	})
	req, _ := http.NewRequest(http.MethodGet, "/hello1", nil)
	res := testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "text/html" || sanitizeHTML(res.String()) != "AAA" {
		t.Error("failed 1", sanitizeHTML(res.String()))
	}
	req, _ = http.NewRequest(http.MethodGet, "/hello2", nil)
	res = testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "text/html" || sanitizeHTML(res.String()) != "AAABBB" {
		t.Error("failed 2", sanitizeHTML(res.String()))
	}
	req, _ = http.NewRequest(http.MethodGet, "/hello3", nil)
	res = testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "text/html" || sanitizeHTML(res.String()) != "AAABBBCCC" {
		t.Error("failed 3", sanitizeHTML(res.String()))
	}
}

type dummyI18n struct {
}

func (_ *dummyI18n) T(key string, args ...string) string {
	r := []string{key}
	r = append(r, args...)
	return strings.Join(r, "")
}

func TestView_HTMLI18n(t *testing.T) {
	n := nova.New()
	n.Use(func(c *nova.Context) error {
		c.Values[view.I18nContextKey] = &dummyI18n{}
		c.Next()
		return nil
	})
	n.Use(view.Handler(view.Options{
		Directory: "testdata",
	}))
	router.Route(n).Get("/hello").Use(func(c *nova.Context) error {
		v := view.Extract(c)
		v.Data["Key4"] = "DDD"
		v.HTML("dir3/dir31")
		return nil
	})
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	res := testkit.NewDummyResponse()
	n.ServeHTTP(res, req)
	if res.Header().Get(view.ContentType) != "text/html" || sanitizeHTML(res.String()) != "AAABBBCCCAAABBBCCCDDD" {
		t.Error("failed 1", sanitizeHTML(res.String()))
	}
}
