package helpers_test

import (
	"math"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/alaa-aqeel/todo/src/helpers"
)

func TestNewPagination(t *testing.T) {
	query := url.Values{}
	query.Set("page", "2")
	query.Set("limit", "15")
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)

	p := helpers.NewPagination(request)

	if p.Page() != 2 {
		t.Errorf("Expected page to be 2, got %d", p.Page())
	}
	if p.Limit() != 15 {
		t.Errorf("Expected limit to be 15, got %d", p.Limit())
	}
}

func TestGetPaginationArgs(t *testing.T) {
	query := url.Values{}
	query.Set("page", "-1")   // Invalid page
	query.Set("limit", "200") // Invalid limit
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)

	p := &helpers.Pagination{}
	p.GetPaginationArgs(request)

	if p.Page() != 1 {
		t.Errorf("Expected default page to be 1, got %d", p.Page())
	}
	if p.Limit() != 10 {
		t.Errorf("Expected default limit to be 10, got %d", p.Limit())
	}
}

func TestSetTotal(t *testing.T) {
	query := url.Values{}
	query.Set("page", "1")
	query.Set("limit", "15")
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)
	p := helpers.NewPagination(request)
	p.SetTotal(95)

	if p.Total() != 95 {
		t.Errorf("Expected total to be 95, got %d", p.Total())
	}
	expectedTotalPages := int(math.Ceil(95.0 / 15.0))
	if p.TotalPages() != expectedTotalPages {
		t.Errorf("Expected totalPages to be %d, got %d", expectedTotalPages, p.TotalPages())
	}
}

func TestSetData(t *testing.T) {
	p := &helpers.Pagination{}
	data := []string{"item1", "item2"}
	p.SetData(data)

	if !reflect.DeepEqual(p.ToMap()["data"], data) {
		t.Errorf("Expected data to be %v, got %v", data, p.ToMap()["data"])
	}
}

func TestOffset(t *testing.T) {
	query := url.Values{}
	query.Set("page", "3")
	query.Set("limit", "10")
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)

	p := helpers.NewPagination(request)
	expectedOffset := 20

	if p.Offset() != expectedOffset {
		t.Errorf("Expected offset to be %d, got %d", expectedOffset, p.Offset())
	}
}

func TestNavigation(t *testing.T) {
	query := url.Values{}
	query.Set("page", "2")
	query.Set("limit", "5")
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)
	p := helpers.NewPagination(request)
	p.SetTotal(15)

	if !p.HasNext() {
		t.Error("Expected HasNext to be true")
	}
	if p.NextPage() != 3 {
		t.Errorf("Expected NextPage to be 3, got %d", p.NextPage())
	}
	if !p.HasPrevious() {
		t.Error("Expected HasPrevious to be true")
	}
}

func TestEdgeCases(t *testing.T) {
	query := url.Values{}
	query.Set("page", "1")
	query.Set("limit", "5")
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)
	p := helpers.NewPagination(request)
	p.SetTotal(15)

	if !p.IsFirst() {
		t.Error("Expected IsFirst to be true")
	}
	if p.IsLast() {
		t.Error("Expected IsLast to be false")
	}
}

func TestHasPagination(t *testing.T) {
	query := url.Values{}
	query.Set("page", "1")
	query.Set("limit", "10")
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)
	p := helpers.NewPagination(request)
	p.SetTotal(15)

	if !p.HasPagination() {
		t.Error("Expected HasPagination to be true")
	}

	p.SetTotal(1)
	if p.HasPagination() {
		t.Error("Expected HasPagination to be false")
	}
}

func TestIsEmpty(t *testing.T) {
	query := url.Values{}
	query.Set("page", "1")
	query.Set("limit", "10")
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)
	p := helpers.NewPagination(request)
	p.SetTotal(0)

	if !p.IsEmpty() {
		t.Error("Expected IsEmpty to be true")
	}

	p.SetTotal(10)
	if p.IsEmpty() {
		t.Error("Expected IsEmpty to be false")
	}
}
