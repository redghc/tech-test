# Prueba Técnica (A) - Golang Developer

## Descripción

El objetivo de esta prueba técnica es crear un endpoint en un servicio REST que procese un body de entrada y un header de autorización, y continue la solicitud recibida redirigiendo a otro servicio.

## Requerimientos

1. **Endpoint**: Crear un endpoint HTTP que acepte un request con un body y un header de autorización.
2. **Procesamiento del Body**:

   - Verificar si el body contiene los parámetros `shipments.$.origin.zipCode` y `shipments.$.destination.zipCode`
   - Si alguno de los parámetros esta presente realizar el procesamiento siguiente segun corresponda:
     - Asignar el valor de la llave `shipments.$.X.zipCode` a la llave `shipments.$.X.zipcode` (renombrado a lowercase).

3. **Procesamiento del Header**:

   - Obtener el valor del header `Authorization`.

4. **Request a URL Externa**:
   - Realizar una petición HTTP a una URL externa POST: ([Creación de orden](https://sandbox.99minutos.com/api/v3/orders)) utilizando el body modificado y el header `Authorization`.
5. **Respuesta**:
   - El endpoint debe devolver la respuesta recibida desde la URL externa.

## Implementación

- Utilizar [Golang](https://golang.org/) para desarrollar la solución.
- Utilizar servidor HTTP [Fiber](https://gofiber.io/)

## Entrega

- Sube tu solución a un repositorio en GitHub o GitLab.
- Incluye un archivo `README.md` en el repositorio con las instrucciones necesarias para correr tu solución localmente.
- Comparte el enlace al repositorio.

## Ejemplo de Entrada y Salida

### Request de entrada

**Header**:

```http
Authorization: Bearer <token>
```

**Body**:

```json
{
  "shipments": [
    {
      "sender": {
        "firstName": "Name",
        "lastName": "Lastname",
        "email": "sender@email.com",
        "phone": "+521234567890"
      },
      "recipient": {
        "firstName": "Name",
        "lastName": "Lastname",
        "email": "receiver@email.com",
        "phone": "+529999999999"
      },
      "origin": {
        "address": "Av. Pdte. Masaryk 123-Planta Baja, Polanco, Polanco V Secc, Miguel Hidalgo, 11560 Ciudad de México, CDMX",
        "reference": "Edificio de oficinas",
        "country": "MEX",
        "city": "CDMX",
        "zipCode": "11560"
      },
      "destination": {
        "address": "Playa Tabachines 13, Militar Marte, Iztacalco, 08830 Ciudad de México, CDMX",
        "reference": "Casa 2 pisos",
        "country": "MEX",
        "city": "CDMX",
        "zipCode": "08830"
      },
      "items": [
        {
          "size": "xs",
          "weight": 999
        }
      ],
      "payments": {
        "insured": false
      },
      "options": {
        "requiresSignature": false,
        "twoFactorAuth": false,
        "requiresIdentification": false,
        "notes": "Additional information"
      },
      "deliveryType": "SPT"
    }
  ],
  "draft": false
}
```

## Generación de JWT

Para generar un JWT y validar el funcionamiento utilizar el siguiente CURL

```bash
curl --location 'https://sandbox.99minutos.com/api/v3/oauth/token' \
--header 'Content-Type: application/json' \
--data '{
    "client_id": "387f1a91-be22-454c-8eeb-c3018c6026d9",
    "client_secret": "Os0qBLk6hibUu9icUa6iZBAfC_"
}'
```

## Request a URL externa

Realizar un request a la URL de [Creación de orden](https://sandbox.99minutos.com/api/v3/orders) con el body modificado y el header de autorización.

## Respuesta esperada

El servicio devolverá la respuesta obtenida de la URL de [Creación de orden](https://sandbox.99minutos.com/api/v3/orders).

## Datos de la cuenta utilizada

- **Usuario**: tech.test.99@99minutos.com
- **Password**: 99Minutos.com
- **Plataforma**: https://authsandbox.99minutos.com/
