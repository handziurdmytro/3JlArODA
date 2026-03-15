Щоб підняти міграції, требе використати таку команду, ти маєш бути в папці з міграціями:
migrate -path migrations/ -database "postgres://USER:PASSWORD@127.0.0.1:5432/DB_NAME?sslmode=disable" up
USER - має бути ім'я юзера, на яке ти створював бд
DB_NAME - назва бдшки
PASSWORD - пароль від твоєї карточки, монобанк
