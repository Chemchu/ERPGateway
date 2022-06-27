package graphQL

func QUERY_VENTAS() string {
	return `query Ventas($find: VentasFind) {
  ventas(find: $find) {
    _id
    productos {
      _id
      nombre
      proveedor
      familia
      precioVenta
      precioCompra
      precioFinal
      iva
      margen
      ean
      cantidadVendida
      createdAt
      updatedAt
      dto
    }
    dineroEntregadoEfectivo
    dineroEntregadoTarjeta
    precioVentaTotalSinDto
    precioVentaTotal
    cambio
    cliente {
      _id
      nif
      nombre
      calle
      cp
    }
    vendidoPor {
      _id
      nombre
      apellidos
      rol
      email
    }
    modificadoPor {
      _id
      nombre
      apellidos
      rol
      email
    }
    tipo
    descuentoEfectivo
    descuentoPorcentaje
    tpv
    createdAt
    updatedAt
  }
}`
}
