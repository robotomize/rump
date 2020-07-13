# rump

Микросервис обновления игрового состояния и получения актуальных данных по игроку

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
