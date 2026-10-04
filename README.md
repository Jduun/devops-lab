# devops-lab

Поднимем `hello-world` сервис на Go в Minikube кластере с 2 репликами.

## Docker-образ
### Сборка
```sh
docker build -t jduun/hello-world:0.1.0 .
```

### Публикация
```sh
docker login
docker push jduun/hello-world:0.1.0
```
Образ на Docker Hub: https://hub.docker.com/r/jduun/hello-world.

### Проверка на уязвимости
```sh
trivy image jduun/hello-world:0.1.0@sha256:1353ad931a4160a5d8047eff165545b75145628b26f0e4d7ecfa58bfa660bd42
```

## API приложения

### `GET /`
#### Описание
Основной эндпоинт сервиса.

#### Ответ
Статус: `200 OK`

```json
{
  "message": "Hello World"
}
```

### `GET /ready`
#### Описание
Проверка готовности сервиса.

#### Ответ
Статус: `200 OK`

```json
{
  "status": "ok"
}
```


### `GET /health`
#### Описание
Проверка работоспособности приложения.

#### Ответ
Статус: `200 OK`

```json
{
  "status": "ok"
}
```

## Запуск в Minikube

### Запуск Minikube кластера
```sh
minikube start
```

### Просмотр сгенерированных манифестов
```sh
helm template prod ./charts/hello-world
```

### Установка Helm chart
```sh
helm upgrade --install prod ./charts/hello-world
```

### Проверяем установку
```sh
kubectl get deploy,pods,svc
```

### Проброс портов
```sh
minikube tunnel
```

### Получаем `EXTERNAL-IP`
```sh
kubectl get svc prod-hello-world
```

### Запрос к сервису
```sh
curl "http://<EXTERNAL-IP>:80"
```
Или переходим в браузере по этой ссылке.

### Удаление релиза
```sh
helm uninstall prod
```

## Результаты

### Скриншот
![screen.png](images/screen.png)

### Схема Minikube-кластера

![cluster.png](images/cluster.png)
