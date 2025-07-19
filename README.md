Api hecha en GO para obtener los datos de clientes y productos.

La api se corre con el siguiente comando: go run .\main.go

Se levanta por defecto en el puerto 8082.

La api tiene 2 endpoints los cuales son los siguientes

GET http://localhost:8082/api/customer/{id}

Ejemplo de curl

curl --request GET \
  --url http://localhost:8082/api/customer/1 \
  --header 'Content-Type: text/plain' \
  --header 'User-Agent: insomnia/11.3.0'

POST http://localhost:8082/api/product/get-list-of-products

Ejemplo de curl

curl --request POST \
  --url http://localhost:8082/api/product/get-list-of-products \
  --header 'Content-Type: text/plain' \
  --header 'User-Agent: insomnia/11.3.0' \
  --data '[1,2]'