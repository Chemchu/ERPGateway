package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"example.com/m/v2/graph/generated"
	"example.com/m/v2/graph/model"
)

func (r *mutationResolver) AddCliente(ctx context.Context, nif string, nombre string, calle *string, cp *string) (*model.ClienteMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCliente(ctx context.Context, id string) (*model.ClienteMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCliente(ctx context.Context, id string, nif *string, nombre *string, calle *string, cp *string) (*model.ClienteMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddEmpleado(ctx context.Context, nombre string, apellidos string, dni string, rol string, genero *string, email string) (*model.EmpleadoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEmpleado(ctx context.Context, id string) (*model.EmpleadoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEmpleado(ctx context.Context, id string, nombre *string, apellidos *string, dni *string, rol *string, genero *string, email *string, horasPorSemana *float64, fechaAlta *string) (*model.EmpleadoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddProductosFile(ctx context.Context, csv string) (*model.ResponseMutation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddClientesFile(ctx context.Context, csv string) (*model.ResponseMutation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddVentasFile(ctx context.Context, ventasJSON string) (*model.ResponseMutation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddProducto(ctx context.Context, producto model.ProductoAddInput) (*model.ProductoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProducto(ctx context.Context, id string) (*model.ProductoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProducto(ctx context.Context, producto model.ProductoUpdateInput) (*model.ProductoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddCierreTpv(ctx context.Context, cierre model.CierreTPVInput) (*model.CierreTPVMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCierreTpv(ctx context.Context, id string) (*model.TPVMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCierreTpv(ctx context.Context, cierre *model.CierreTPVInput) (*model.CierreTPVMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddTpv(ctx context.Context, nombre string, enUsoPor *string, libre *bool, cajaInicial *int) (*model.TPVMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTpv(ctx context.Context, id string) (*model.TPVMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTpv(ctx context.Context, id string, nombre *string, enUsoPor *string, libre *bool, cajaInicial *int) (*model.TPVMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OcupyTpv(ctx context.Context, idEmpleado string, idTpv string, cajaInicial float64) (*model.TPVMutationJwtResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FreeTpv(ctx context.Context, idEmpleado string, idTpv string) (*model.TPVMutationJwtResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddVenta(ctx context.Context, fields model.VentaFields) (*model.VentaMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVenta(ctx context.Context, id string) (*model.VentaMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVenta(ctx context.Context, id string, productos []*model.ProductoVendidoInput, dineroEntregadoEfectivo *float64, descuentoPorcentaje *float64, precioVentaTotal float64, cambio *float64, clienteID *model.ClienteInput, vendidoPor *model.EmpleadoInput, modificadoPor *model.EmpleadoInput, tipo *string, descuentoEfectivo *float64) (*model.VentaMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, loginValues model.Credentials) (*model.LoginResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ValidateJwt(ctx context.Context, jwt string) (*model.JwtValidationResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cliente(ctx context.Context, find model.ClienteFind) (*model.Cliente, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Clientes(ctx context.Context, find *model.ClientesFind, limit *int) ([]*model.Cliente, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Empleado(ctx context.Context, find model.EmpleadoFind) (*model.Empleado, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Empleados(ctx context.Context, find *model.EmpleadosFind, limit *int) ([]*model.Empleado, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Producto(ctx context.Context, find model.ProductoFind) (*model.Producto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Productos(ctx context.Context, find *model.ProductosFind, limit *int) ([]*model.Producto, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CierreTpv(ctx context.Context, find model.CierreTPVFind) (*model.CierreTpv, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CierresTPVs(ctx context.Context, find *model.CierresTPVFind, limit *int) ([]*model.CierreTpv, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tpv(ctx context.Context, find model.TPVFind) (*model.Tpv, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tpvs(ctx context.Context, find *model.TPVsFind, limit *int) ([]*model.Tpv, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Venta(ctx context.Context, id string) (*model.Venta, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ventas(ctx context.Context, find *model.VentasFind, limit *int, order *string, offset *int) ([]*model.Venta, error) {
	// graphqlRedirect.RedirectQuery("aaa")
	// ctx
	fmt.Println("#############################")
	fmt.Println(ctx.Value("find"))
	fmt.Println("#############################")
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
