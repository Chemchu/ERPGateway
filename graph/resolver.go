package graph

import "example.com/m/v2/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Venta                     map[string]model.Venta
	Producto                  map[string]model.Producto
	CierreTPVFind             map[string]model.CierreTPVFind
	CierreTPVInput            map[string]model.CierreTPVInput
	CierreTPVMutationResponse map[string]model.CierreTPVMutationResponse
	CierreTpv                 map[string]model.CierreTpv
	Cliente                   map[string]model.Cliente
	ClienteFind               map[string]model.ClienteFind
	ClienteInput              map[string]model.ClienteInput
	ClienteMutationResponse   map[string]model.ClienteMutationResponse
	ClientesFind              map[string]model.ClientesFind
	Credentials               map[string]model.Credentials
	Empleado                  map[string]model.Empleado
	EmpleadoFind              map[string]model.EmpleadoFind
	EmpleadoInput             map[string]model.EmpleadoInput
	EmpleadoMutationResponse  map[string]model.EmpleadoMutationResponse
	EmpleadosFind             map[string]model.EmpleadosFind
	JwtValidationResult       map[string]model.JwtValidationResult
	LoginResult               map[string]model.LoginResult
	ProductoAddInput          map[string]model.ProductoAddInput
	ProductoFind              map[string]model.ProductoFind
	ProductoMutationResponse  map[string]model.ProductoMutationResponse
	ProductoUpdateInput       map[string]model.ProductoUpdateInput
	ProductoVendido           map[string]model.ProductoVendido
	ProductoVendidoInput      map[string]model.ProductoVendidoInput
	ProductosFind             map[string]model.ProductosFind
	ResponseMutation          map[string]model.ResponseMutation
	TPVFind                   map[string]model.TPVFind
	TPVMutationJwtResponse    map[string]model.TPVMutationJwtResponse
	TPVMutationResponse       map[string]model.TPVMutationResponse
	TPVsFind                  map[string]model.TPVsFind
	Tpv                       map[string]model.Tpv
	VentaFields               map[string]model.VentaFields
	VentaMutationResponse     map[string]model.VentaMutationResponse
	VentasFind                map[string]model.VentasFind
}
