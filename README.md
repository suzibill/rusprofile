# rusprofile

**Тестовое задание от штрафов.нет**   
Необходимо сделать gRPC обёртку над сайтом https://www.rusprofile.ru/

## API

Сервис реализует один метод, принимающий на вход ИНН компании, ищущий компанию на rusprofile,
и возвращающий её ИНН, КПП, название, ФИО руководителя.

## Технологии

* Go
* API через [gRPC](https://grpc.io/docs/languages/go/quickstart/).
* API через HTTP с помощью [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).
* Swagger UI с документацией, сгенерированной из .proto файла с помощью protoc-gen-swagger.
  Документация доступна по пути `/swaggerui`.
* Упакован в Docker контейнер.


## Настройка и запуск проекта

1. Склонировать репозиторий с помощью команды:
   ```
   git clone https://github.com/suzibill/rusprofile
   ```
2. Перейти в папку с проектом:
   ```
   cd rusprofile
   ```
3. Установить Docker с [официального сайта](https://www.docker.com/products/docker-desktop)
4. Сбилдить Docker-образ :
   ```
   make docker-build
   ```
5. Запустить сервис в Docker
  ```
   make docker-run
   ```
6. Получить swaggerui документацию
   ```
   127.0.0.1:8080/swaggerui
   ```
7. Получить информацию о компании по ИНН 
    ```
    http://127.0.0.1:8080/v1/companies/{INN}
    ```
