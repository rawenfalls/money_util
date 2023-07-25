# Консольная утилита, которая выводит курсы валют ЦБ РФ за определенную дату.
### Для получения курсов необходимо использовать официальный API ЦБ РФ https://www.cbr.ru/development/sxml/.
* Установка:
```
git clone https://github.com/rawenfalls/money_util.git
```
* Сборка:
```
go build money_util.go
```
* Интерфейс:
	* currency_rates --code=USD --date=2022-10-08

* Описание параметров:
	* --code - код валюты в формате ISO 4217
	* date - дата в формате YYYY-MM-DD
* Пример работы:
```
money_util.exe --code=USD --date=2023-07-25
//USD (Доллар США): 90,4890
```
