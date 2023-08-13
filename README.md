# 🍲boiler

Esta app CLI inicia un nuevo módulo copiando un módulo de plantilla publico (ej. [github.com/gophers-latam/minimalchi](https://github.com/gophers-latam/minimalchi)).

*```boiler``` esta basado en la implementación experimental de ```gonew``` para extender y localizar a Español.

**Uso base:**

```shell
boiler src.modName@version [dst.name/mod [dirName]]

# src.modName = fuente.nombre-modulo
# @version = tag/rama a descargar
# [...] -opcional, donde:
# dst.name/mod = destino.nombre/modulo
# [dirName] = nombrado personalizado directorio

# Ejemplo: 
boiler golang.org/x/example/hello mi.dominio/hola mi-hola
```

```src.modName@version``` y ```dst.name/mod``` deben usar un versionamiento válido.

**Como instalar:**

- Para estar disponible via ```$GOPATH/bin```
```shell
go install github.com/gophers-latam/boiler@latest
```

**Plantillas propuestas:**
- [gophers-latam/minimal](https://github.com/gophers-latam/minimal)
- [gophers-latam/minimalchi](https://github.com/gophers-latam/minimalchi)
- [gophers-latam/small-template](https://github.com/gophers-latam/small-template)
- **Hexagonal-arch** [```pendiente```]  por [@AndresXLP](https://github.com/AndresXLP)

## TODO:

- [ ] Mejorar salida de ayuda.
- [ ] Considerar integrar libreria de terceros tipo: termdash, cobra.
- [ ] Agregar instruccion para agregar templates.

**Agregar flujo interactivo, no salida de error por defecto:**
- [ ] Mejorar y agregar instrucción interactiva para configuración de inicialización: *version go module*, *tidy dependencias*, *renombrado mejorado*, etc.
- [ ] Agregar instrucción para eliminar o agregar archivos genericos.
- [ ] Refactorizar y agregar patrón concurrente.


<hr style="height:1px;border-width:0;color:gray;background-color:gray">

...*Más detalles sobre requerimientos en [Boiler Google Doc](https://docs.google.com/document/d/1xI7gqO1E9h2sKU574zyUrgKXZWyLpmVhiTgWcTZ7DvI/edit?usp=sharing)*