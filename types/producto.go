package types

type ProductoVendido struct {
	ID              *string  `json:"_id"`
	Nombre          *string  `json:"nombre"`
	Familia         *string  `json:"familia"`
	Proveedor       *string  `json:"proveedor,omitempty"`
	PrecioCompra    *float64 `json:"precioCompra"`
	PrecioVenta     *float64 `json:"precioVenta"`
	PrecioFinal     *float64 `json:"precioFinal"`
	CantidadVendida *int     `json:"cantidadVendida"`
	Dto             *float64 `json:"dto"`
	Iva             *float64 `json:"iva"`
	Margen          *float64 `json:"margen"`
	Ean             *string  `json:"ean"`
}

type ProductoDevuelto struct {
	ID               *string  `json:"_id"`
	Nombre           *string  `json:"nombre"`
	Familia          *string  `json:"familia"`
	Proveedor        *string  `json:"proveedor"`
	PrecioCompra     *float32 `json:"precioCompra"`
	PrecioVenta      *float32 `json:"precioVenta"`
	PrecioFinal      *float32 `json:"precioFinal"`
	CantidadDevuelta *int32   `json:"cantidadDevuelta"`
	Dto              *float32 `json:"dto"`
	Iva              *float32 `json:"iva"`
	Margen           *float32 `json:"margen"`
	Ean              *string  `json:"ean"`
}

type Producto struct {
	ID              *string  `json:"_id"`
	Nombre          *string  `json:"nombre"`
	Familia         *string  `json:"familia"`
	Proveedor       *string  `json:"proveedor"`
	PrecioCompra    *float32 `json:"precioCompra"`
	PrecioVenta     *float32 `json:"precioVenta"`
	Iva             *float32 `json:"iva"`
	Margen          *float32 `json:"margen"`
	Ean             *string  `json:"ean"`
	Promociones     *string  `json:"promociones"`
	Alta            *bool    `json:"alta"`
	Cantidad        *int32   `json:"cantidad"`
	CantidadRestock *int32   `json:"cantidadRestock"`
	CreatedAt       *string  `json:"createdAt"`
	UpdatedAt       *string  `json:"updatedAt"`
}
