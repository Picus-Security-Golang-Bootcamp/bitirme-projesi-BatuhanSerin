package product

import (
	"reflect"
	"testing"

	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/models"
)

func Test_responseToProduct(t *testing.T) {
	type args struct {
		p *api.Product
	}
	tests := []struct {
		name string
		args args
		want *models.Product
	}{
		{name: "Test_responseToProduct", args: args{p: &api.Product{Name: ""}}, want: &models.Product{Name: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := responseToProduct(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responseToProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productsToResponse(t *testing.T) {
	type args struct {
		ps *[]models.Product
	}
	tests := []struct {
		name string
		args args
		want []*api.Product
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productsToResponse(tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productsToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productsToResponseWithoutCategory(t *testing.T) {
	type args struct {
		ps *[]models.Product
	}
	tests := []struct {
		name string
		args args
		want []*api.Product
	}{
		{name: "Test_productsToResponseWithoutCategory", args: args{ps: &[]models.Product{models.Product{Name: ""}}}, want: []*api.Product{&api.Product{Name: ""}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productsToResponseWithoutCategory(tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productsToResponseWithoutCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductToResponseWithoutCategory(t *testing.T) {
	type args struct {
		p *models.Product
	}
	tests := []struct {
		name string
		args args
		want *api.Product
	}{
		{name: "TestProductToResponseWithoutCategory", args: args{p: &models.Product{Name: ""}}, want: &api.Product{Name: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductToResponseWithoutCategory(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductToResponseWithoutCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductToResponse(t testing.T) {
	type args struct {
		p models.Product
	}
	tests := []struct {
		name string
		args args
		want api.Product
	}{
		{name: "TestProductToResponse", args: args{p: models.Product{}}, want: api.Product{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductToResponse(&tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
