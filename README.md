# rusprofile

Необходимо сделать gRPC обёртку над сайтом https://www.rusprofile.ru/

## API

Сервис должен реализовывать один метод, принимающий на вход ИНН компании, ищущий компанию на rusprofile,
и возвращающий её ИНН, КПП, название, ФИО руководителя.

## Технологии

* Сервис должен быть написан на Go с использованием Go Modules.
* Предоставлять API через [gRPC](https://grpc.io/docs/languages/go/quickstart/).
* Предоставлять API через HTTP с помощью [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).
* Предоставлять Swagger UI с документацией, сгенерированной из .proto файла с помощью protoc-gen-swagger.
  Документация должна быть доступна по пути `/swaggerui`.
* Быть упакован в Docker контейнер.
