# rump

Микросервис обновления игрового состояния и получения актуальных данных по игроку

### Docker

```sh
sudo docker build -t rump .
```

### Ручной вариант сборки
```sh
  git clone https://github.com/robotomize/rump.git
  основной сервис
  go build cmd/syncrcvpos
  консольная утилита для получения позиции по игроку
  go build tools/rcvposcli
  консольная утилита для отправки игрока(позиции)
  go build tools/syncposcli
  утилита для нагрузочного тестирования(зачаток)
  go build tools/stubsrvcli
```
### Ручной запуск тестов 
```sh
go test -v -cover -covermode=atomic ./...
```
### Для сборки всех бинарников
```sh
make build
```
### Для запуска тестов 
```sh
make test
```
### Использование
Запускаем сервер
Утилита syncposcli отправляет позицию игрока в сервис
```sh 
syncposcli --id 1 --x 10 --y 11 --z 20
syncposcli --id 1 --x 10 --y 11 --z 15
```
Утилита rcvposcli забирает позицию
```sh 
rcvposcli --id 1
```
