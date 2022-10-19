Тестовое задание №2 (backend)<br>
Создание кеширующего веб-сервера<br>
Задача<br>
Следуя описанию и рекомендациям написать веб-сервер для сохранения и раздачи электронных документов. Вы можете выбрать СУБД для реализации приложения по своему усмотрению.

Язык реализации задания
go

Описание приложения
Реализовать http-сервер, который реализует описанное ниже REST API.

1.	Регистрация (логин, пароль – создание нового пользователя)
2.	Аутентификация (получить токен авторизации по логину и паролю)
3.	Загрузка нового документа
4.	Получение списка документов
5.	Получение одного документа
6.	Удаление документа
7.	Завершение авторизованной сессии работы

Пользователь может загружать файлы, управлять коллекцией и выборочно делиться ими с другими. При этом система должна рассчитываться на значительную нагрузку при получении списков и файлов. Запросы HEAD не должны возвращать данных. Запросы GET/HEAD должны отдаваться из внутреннего кеша. Остальные запросы должны инвалидировать кеш, желательно выборочно.

Уровни сложности:

•	Первый: запросы 4 и 5.<br>
•	Второй: добавить 3 и 6.<br>
•	Третий: все запросы.<br>
•	Четвертый: все запросы + выдача результата из кэша (если есть)<br>

Общая модель ответа для всех методов:

{<br>
"error": { <br>
      "code": 123,<br>
      "text": "so sad"<br>
      },<br>
"response": {<br>
      ...<br>
      },<br>
"data": {<br>
      ...<br>
      }<br>
}<br>


•	Поля error, response, data присутствуют, только если заполнены<br>
•	Поле response для подтверждения действий<br>
•	Поле data для выдачи содержимого<br>


•	HTTP-cтатусы:<br>
•	Все ок — 200<br>
•	Некорректные параметры — 400<br>
•	Не авторизован — 401<br>
•	Нет прав доступа — 403<br>
•	Неверный метод запроса — 405<br>
•	Нежданчик — 500<br>
•	Метод не реализован - 501<br>



Описание REST API:<br>
1.	Регистрация [POST] /api/register<br>
      •	Вход<br>
      •	token — токен администратора<br>
      •	Фиксированный, задается в конфиге приложения<br>
      •	login — логин нового пользователя<br>
      •	Минимальная длина 8, латиница и цифры<br>
      •	pswd — пароль нового пользователя<br>
      •	минимальная длина 8,<br>
      •	минимум 2 буквы в разных регистрах<br>
      •	минимум 1 цифра<br>
      •	минимум 1 символ (не буква и не цифра)<br>

•	Выход<br>


`{
"response": {
"login": "test"
}
}`



2. Аутентификация [POST] /api/auth<br>
      •	Вход - форма<br>
      •	login<br>
      •	pswd<br>
      •	Выход<br>
      {<br>
      "response": {<br>
      "token": "sfuqwejqjoiu93e29"<br>
      }<br>
      }<br>

3. Загрузка нового документа [POST] /api/docs<br>
      •	Вход — multipart form<br>
      •	meta — параметры запроса. Модель:<br>

      {<br>
      "name": "photo.jpg", "file": true, "public": false,<br>
      "token": "sfuqwejqjoiu93e29", "mime": "image/jpg", "grant": [<br>
      "login1",<br>
      "login2",<br>
      ]<br>
      }<br>
      •	json — данные документа<br>
      •	может отсутствовать<br>
      •	модель не определена<br>
      •	file — файл документа<br>
      •	Выход<br>


      {
      "data": {
            "json": { ... }, 
            "file": "photo.jpg"
      }
      }


4. Получение списка документов [GET, HEAD] /api/docs <br>
      •	Вход <br>

•	token <br>
•	login - опционально — если не указан — то список своих <br>
•	key - имя колонки для фильтрации <br>
•	value - значение фильтра <br>
•	limit - кол-во документов в списке <br>
•	Выход - сортировать по имени и дате создания <br>


`{
"data": {
"docs": [
{
"id": "qwdj1q4o34u34ih759ou1", "name": "photo.jpg",
"mime": "image/jpg", "file": true, "public": false,
"created": "2018-12-24 10:30:56",
"grant": [
"login1",
"login2",
]
}
]
}
}`

5. Получение одного документа [GET, HEAD] /api/docs/<id> <br>
      •	Вход <br>
      •	token <br>
      •	Выход <br>
      •	Если файл — выдать файл с нужным mime <br>
      •	Если JSON: <br>


      {
      "data": {
      ...
      }
      }

6. Удаление документа [DELETE] /api/docs/<id> <br>
      •	Вход <br>
      •	token <br>
      •	Выход <br>


      {
      "response": {
      "qwdj1q4o34u34ih759ou1": true
      }
      }

7. Завершение авторизованной сессии работы [DELETE] <br>
      /api/auth/<token> <br>
      •	Выход <br>


      {
      "response": {
      "qwdj1q4o34u34ih759ou1": true
      }
      }

Библиотеки для реализации (*возможно использовать и другие)<br>
•   HTTP<br>
•   https://golang.org/pkg/net/ <br>
•   https://golang.org/pkg/net/http/ <br>
•	https://golang.org/pkg/net/url/ <br>
•	https://golang.org/pkg/mime/ <br>
•	https://golang.org/pkg/net/http/httptest/ <br>
•	https://golang.org/pkg/mime/multipart/ <br>
•	https://godoc.org/github.com/gocraft/web <br>
•	https://godoc.org/github.com/jarcoal/httpmock <br>
•	Утилиты <br>
•	https://golang.org/pkg/regexp/ <br>
•	https://golang.org/pkg/sync/ <br>
•	https://godoc.org/github.com/satori/go.uuid <br>
•	https://godoc.org/github.com/pkg/errors <br>
•	Базы данных <br>
•	https://golang.org/pkg/database/sql/ <br>
•	https://godoc.org/github.com/lib/pq <br>
•	https://github.com/jackc/pgx <br>
•	https://godoc.org/github.com/globalsign/mgo <br>
•	https://godoc.org/github.com/jmoiron/sqlx <br>
•	https://godoc.org/github.com/gwenn/gosqlite <br>
