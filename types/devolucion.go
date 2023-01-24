package types

type Devolucion struct {
	ID                 *string             `json:"_id"`
	ProductosDevueltos *[]ProductoDevuelto `json:"productosDevueltos"`
	DineroDevuelto     *float32            `json:"dineroDevuelto"`
	VentaOriginal      *Venta              `json:"ventaOriginal"`
	Tpv                *string             `json:"tpv"`
	Cliente            *Cliente            `json:"cliente"`
	Trabajador         *Empleado           `json:"trabajador"`
	ModificadoPor      *Empleado           `json:"modificadoPor"`
	CreatedAt          *string             `json:"createdAt"`
	UpdatedAt          *string             `json:"updatedAt"`
}
