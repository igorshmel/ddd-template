# DDD Template with Unit of Work pattern 

Это полностью функциональны шаблон сервиса, написанный на языке Golang, 
построенный на принципах чистой архитектуры, использующий паттерн Unit of Work 
для корректной работы с базой данных.

## Как это работает

1. Скомпилировать исполняемый файл можно стандартным для GO способом  
   go build /app/cmd/main.go
2. Предварительно необходимо заполнить конфигурационный файл с доступом к базе данных Postgres
   Приложение само создаст необходимые таблицы при запуске
3. Шаблон реализует функционал создания заказа для корзины покупок. 
   Данный функционал требует согласованности при записи данных в БД.
   Поэтому на данном примере легко протестировать паттерн Unit of Work.
4. Шаблон построен на архитектуре DDD (Domain-Driven Design), поэтому некоторые места в коде выглядят как дикий overhead.
   Однако следует понимать, что реальное приложение с постоянно меняющимися бизнес-требованиями - довольно страшный зверь, 
   и данный подход - это одна из попыток накормить и укротить этого зверя.
5. Протестировать работу приложения и реализацию паттерна Unit of Work, при записи данных в БД можно следующим образом:
   Используя openapi.yaml нужно отправить несколько запросов в систему.
   1. Создать пользователя
   2. Создать несколько продуктов
   3. Создать одну или несколько позиций заказа (применяется Unit of Work)
   4. Создать заказ (применяется Unit of Work)

go version go1.18.2