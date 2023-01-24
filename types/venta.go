package types

type Venta struct {
	ID                      *string            `json:"_id"`
	Productos               *[]ProductoVendido `json:"productos"`
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
	CreatedAt               *int64             `json:"createdAt,string"`
	UpdatedAt               *int64             `json:"updatedAt,string"`
}
