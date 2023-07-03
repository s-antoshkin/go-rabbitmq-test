
Запуск контейнера с RabbitMQ (`http://localhost:15672/`):
```bash
docker run -d --hostname rabbitmq --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

Интерактивный доступ к admin CLI tool:
```bash
docker exec -it 7ad(номер контейнера с RabbitMQ) bash
```
Отправка сообщения:
```bash
rabbitmqadmin publish exchange=amq.default routing_key="TestQueue" payload="Hello World"
```