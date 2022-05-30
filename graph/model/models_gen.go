// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CierreTpv struct {
	ID                   string    `json:"_id"`
	Tpv                  string    `json:"tpv"`
	CajaInicial          *float64  `json:"cajaInicial"`
	AbiertoPor           *Empleado `json:"abiertoPor"`
	CerradoPor           *Empleado `json:"cerradoPor"`
	Apertura             *string   `json:"apertura"`
	Cierre               *string   `json:"cierre"`
	NumVentas            *int      `json:"numVentas"`
	VentasEfectivo       *float64  `json:"ventasEfectivo"`
	VentasTarjeta        *float64  `json:"ventasTarjeta"`
	VentasTotales        *float64  `json:"ventasTotales"`
	DineroEsperadoEnCaja *float64  `json:"dineroEsperadoEnCaja"`
	DineroRealEnCaja     *float64  `json:"dineroRealEnCaja"`
	DineroRetirado       *float64  `json:"dineroRetirado"`
	FondoDeCaja          *float64  `json:"fondoDeCaja"`
	Beneficio            *float64  `json:"beneficio"`
	Nota                 *string   `json:"nota"`
}

type CierreTPVFind struct {
	ID    *string `json:"_id"`
	Fecha *string `json:"fecha"`
}

type CierreTPVInput struct {
	ID                   *string        `json:"_id"`
	Tpv                  string         `json:"tpv"`
	CajaInicial          float64        `json:"cajaInicial"`
	AbiertoPor           *EmpleadoInput `json:"abiertoPor"`
	CerradoPor           *EmpleadoInput `json:"cerradoPor"`
	Apertura             string         `json:"apertura"`
	Cierre               *string        `json:"cierre"`
	NumVentas            int            `json:"numVentas"`
	VentasEfectivo       float64        `json:"ventasEfectivo"`
	VentasTarjeta        float64        `json:"ventasTarjeta"`
	VentasTotales        float64        `json:"ventasTotales"`
	DineroEsperadoEnCaja float64        `json:"dineroEsperadoEnCaja"`
	DineroRealEnCaja     float64        `json:"dineroRealEnCaja"`
	DineroRetirado       float64        `json:"dineroRetirado"`
	FondoDeCaja          float64        `json:"fondoDeCaja"`
	Nota                 *string        `json:"nota"`
}

type CierreTPVMutationResponse struct {
	Message    string     `json:"message"`
	Successful bool       `json:"successful"`
	Token      *string    `json:"token"`
	Cierre     *CierreTpv `json:"cierre"`
}

type CierresTPVFind struct {
	Fecha *string `json:"fecha"`
}

type Cliente struct {
	ID     string `json:"_id"`
	Nif    string `json:"nif"`
	Nombre string `json:"nombre"`
	Calle  string `json:"calle"`
	Cp     string `json:"cp"`
}

type ClienteFind struct {
	ID  *string `json:"_id"`
	Nif *string `json:"nif"`
}

type ClienteInput struct {
	ID     string `json:"_id"`
	Nif    string `json:"nif"`
	Nombre string `json:"nombre"`
	Calle  string `json:"calle"`
	Cp     string `json:"cp"`
}

type ClienteMutationResponse struct {
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}

type ClientesFind struct {
	Ids    []string `json:"_ids"`
	Nombre *string  `json:"nombre"`
	Query  *string  `json:"query"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Empleado struct {
	ID             string   `json:"_id"`
	Nombre         string   `json:"nombre"`
	Apellidos      string   `json:"apellidos"`
	Dni            *string  `json:"dni"`
	Rol            string   `json:"rol"`
	Genero         *string  `json:"genero"`
	Email          string   `json:"email"`
	HorasPorSemana *float64 `json:"horasPorSemana"`
	FechaAlta      *string  `json:"fechaAlta"`
}

type EmpleadoFind struct {
	ID     *string `json:"_id"`
	Nombre *string `json:"nombre"`
	Dni    *string `json:"dni"`
}

type EmpleadoInput struct {
	ID             string   `json:"_id"`
	Nombre         string   `json:"nombre"`
	Apellidos      string   `json:"apellidos"`
	Dni            *string  `json:"dni"`
	Rol            string   `json:"rol"`
	Genero         *string  `json:"genero"`
	Email          string   `json:"email"`
	HorasPorSemana *float64 `json:"horasPorSemana"`
	FechaAlta      *string  `json:"fechaAlta"`
}

type EmpleadoMutationResponse struct {
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}

type EmpleadosFind struct {
	Ids    []string `json:"_ids"`
	Nombre *string  `json:"nombre"`
	Rol    *string  `json:"rol"`
	Query  *string  `json:"query"`
}

type JwtValidationResult struct {
	Validado bool `json:"validado"`
}

type LoginResult struct {
	Message string  `json:"message"`
	Success bool    `json:"success"`
	Token   *string `json:"token"`
}

type Producto struct {
	ID              string    `json:"_id"`
	Nombre          *string   `json:"nombre"`
	Proveedor       *string   `json:"proveedor"`
	Familia         *string   `json:"familia"`
	PrecioVenta     *float64  `json:"precioVenta"`
	PrecioCompra    *float64  `json:"precioCompra"`
	Iva             *float64  `json:"iva"`
	Margen          *float64  `json:"margen"`
	Promociones     []*string `json:"promociones"`
	Ean             *string   `json:"ean"`
	Cantidad        *int      `json:"cantidad"`
	CantidadRestock *int      `json:"cantidadRestock"`
	Alta            *bool     `json:"alta"`
	Img             *string   `json:"img"`
	CreatedAt       *string   `json:"createdAt"`
	UpdatedAt       *string   `json:"updatedAt"`
}

type ProductoAddInput struct {
	Nombre          string    `json:"nombre"`
	Proveedor       *string   `json:"proveedor"`
	Familia         *string   `json:"familia"`
	PrecioVenta     float64   `json:"precioVenta"`
	PrecioCompra    float64   `json:"precioCompra"`
	Iva             float64   `json:"iva"`
	Margen          float64   `json:"margen"`
	Promociones     []*string `json:"promociones"`
	Ean             string    `json:"ean"`
	Cantidad        *int      `json:"cantidad"`
	CantidadRestock *int      `json:"cantidadRestock"`
	Alta            bool      `json:"alta"`
	Img             *string   `json:"img"`
}

type ProductoFind struct {
	ID     *string `json:"_id"`
	Nombre *string `json:"nombre"`
	Ean    *string `json:"ean"`
}

type ProductoMutationResponse struct {
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}

type ProductoUpdateInput struct {
	ID              string    `json:"_id"`
	Nombre          *string   `json:"nombre"`
	Proveedor       *string   `json:"proveedor"`
	Familia         *string   `json:"familia"`
	PrecioVenta     *float64  `json:"precioVenta"`
	PrecioCompra    *float64  `json:"precioCompra"`
	Iva             *float64  `json:"iva"`
	Margen          *float64  `json:"margen"`
	Promociones     []*string `json:"promociones"`
	Ean             *string   `json:"ean"`
	Cantidad        *int      `json:"cantidad"`
	CantidadRestock *int      `json:"cantidadRestock"`
	Alta            *bool     `json:"alta"`
	Img             *string   `json:"img"`
}

type ProductoVendido struct {
	ID              string   `json:"_id"`
	Nombre          *string  `json:"nombre"`
	Proveedor       *string  `json:"proveedor"`
	Familia         *string  `json:"familia"`
	PrecioVenta     *float64 `json:"precioVenta"`
	PrecioCompra    *float64 `json:"precioCompra"`
	Iva             *float64 `json:"iva"`
	Margen          *float64 `json:"margen"`
	Ean             *string  `json:"ean"`
	CantidadVendida *int     `json:"cantidadVendida"`
	CreatedAt       *string  `json:"createdAt"`
	UpdatedAt       *string  `json:"updatedAt"`
}

type ProductoVendidoInput struct {
	ID              string   `json:"_id"`
	Nombre          *string  `json:"nombre"`
	Proveedor       *string  `json:"proveedor"`
	Familia         *string  `json:"familia"`
	PrecioVenta     *float64 `json:"precioVenta"`
	PrecioCompra    *float64 `json:"precioCompra"`
	Iva             *float64 `json:"iva"`
	Margen          *float64 `json:"margen"`
	Ean             *string  `json:"ean"`
	CantidadVendida *int     `json:"cantidadVendida"`
	CreatedAt       *string  `json:"createdAt"`
	UpdatedAt       *string  `json:"updatedAt"`
	Dto             *float64 `json:"dto"`
}

type ProductosFind struct {
	Ids       []string `json:"_ids"`
	Nombre    *string  `json:"nombre"`
	Familia   *string  `json:"familia"`
	Proveedor *string  `json:"proveedor"`
	Query     *string  `json:"query"`
}

type ResponseMutation struct {
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}

type Tpv struct {
	ID          string    `json:"_id"`
	Nombre      *string   `json:"nombre"`
	EnUsoPor    *Empleado `json:"enUsoPor"`
	Libre       *bool     `json:"libre"`
	CajaInicial *int      `json:"cajaInicial"`
	CreatedAt   *string   `json:"createdAt"`
	UpdatedAt   *string   `json:"updatedAt"`
}

type TPVFind struct {
	ID         *string `json:"_id"`
	Nombre     *string `json:"nombre"`
	EmpleadoID *string `json:"empleadoId"`
}

type TPVMutationJwtResponse struct {
	Token      string `json:"token"`
	Successful *bool  `json:"successful"`
}

type TPVMutationResponse struct {
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}

type TPVsFind struct {
	Libre bool `json:"libre"`
}

type Venta struct {
	ID                      string             `json:"_id"`
	Productos               []*ProductoVendido `json:"productos"`
	DineroEntregadoEfectivo *float64           `json:"dineroEntregadoEfectivo"`
	DineroEntregadoTarjeta  *float64           `json:"dineroEntregadoTarjeta"`
	PrecioVentaTotalSinDto  *float64           `json:"precioVentaTotalSinDto"`
	PrecioVentaTotal        *float64           `json:"precioVentaTotal"`
	Cambio                  *float64           `json:"cambio"`
	Cliente                 *Cliente           `json:"cliente"`
	VendidoPor              *Empleado          `json:"vendidoPor"`
	ModificadoPor           *Empleado          `json:"modificadoPor"`
	Tipo                    *string            `json:"tipo"`
	DescuentoEfectivo       *float64           `json:"descuentoEfectivo"`
	DescuentoPorcentaje     *float64           `json:"descuentoPorcentaje"`
	Tpv                     *string            `json:"tpv"`
	CreatedAt               *string            `json:"createdAt"`
	UpdatedAt               *string            `json:"updatedAt"`
}

type VentaFields struct {
	Productos               []*ProductoVendidoInput `json:"productos"`
	DineroEntregadoEfectivo float64                 `json:"dineroEntregadoEfectivo"`
	DineroEntregadoTarjeta  float64                 `json:"dineroEntregadoTarjeta"`
	PrecioVentaTotalSinDto  float64                 `json:"precioVentaTotalSinDto"`
	PrecioVentaTotal        float64                 `json:"precioVentaTotal"`
	Cambio                  float64                 `json:"cambio"`
	Cliente                 *ClienteInput           `json:"cliente"`
	VendidoPor              *EmpleadoInput          `json:"vendidoPor"`
	ModificadoPor           *EmpleadoInput          `json:"modificadoPor"`
	Tipo                    string                  `json:"tipo"`
	DescuentoEfectivo       float64                 `json:"descuentoEfectivo"`
	DescuentoPorcentaje     float64                 `json:"descuentoPorcentaje"`
	Tpv                     string                  `json:"tpv"`
}

type VentaMutationResponse struct {
	ID         *string `json:"_id"`
	Message    string  `json:"message"`
	Successful bool    `json:"successful"`
	CreatedAt  *string `json:"createdAt"`
}

type VentasFind struct {
	Ids          []string `json:"_ids"`
	ClienteID    *string  `json:"clienteId"`
	Tipo         *string  `json:"tipo"`
	VendedorID   *string  `json:"vendedorId"`
	CreatedAt    *string  `json:"createdAt"`
	FechaInicial *string  `json:"fechaInicial"`
	FechaFinal   *string  `json:"fechaFinal"`
	Tpv          *string  `json:"tpv"`
	Query        *string  `json:"query"`
}