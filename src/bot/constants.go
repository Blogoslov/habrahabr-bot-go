package bot


const articleRegexPattern = "(?:https://|)habrahabr.ru/(?:post|company/\\w+/blog)/\\d{1,6}(?:/|)"

const userRegexPattern = "^https://habrahabr.ru/users/[\\w\\s_]+/$"
 
const messageText = "{title} <a href='{IV}'>(IV)</a>\n\n<a href='{link}'>Перейти к статье</a>\n\n<a href='{link}#comments'>Перейти к комментариям</a>"

const maxArticlesLimit = 25 // Служит для ограничения отправки статей, чтобы Telegram не заблокировал бота

// ссылка на InstantView с {url} вместо ссылки на статью
const instantViewURL = "https://t.me/iv?url={url}&rhash=2cb77307aed3b2"

const allArticlesURL = "https://habrahabr.ru/rss/all/?with_hubs=true:?with_tags=true:"

const bestArticlesURL = "https://habrahabr.ru/rss/best/?with_hubs=true:?with_tags=true:"


const helpText = `
📝 <b>КОМАНДЫ</b>:
/help – показать помощь
/my_tags – показать 📃 список тегов, на которые пользователь подписан
/add_tags – добавить теги (пример: /add_tags IT Алгоритмы)
/del_tags – удалить теги (пример: /del_tags IT Алгоритмы)
/del_all_tags – ❌ удалить ВСЕ теги
/copy_tags – ✂️ скопировать теги из профиля на habrahabr'e (пример: /copy_tags https://habrahabr.ru/users/kirtis/)
/stop – 🔕 приостановить рассылку (для продолжения рассылки - /start)
/get_best – получить лучшие статьи за день (по-умолчанию присылается 5, но можно через пробел указать другое количество)
Также вы можете отправить боту ссылку на статью, после чего он отправит ссылку на InstantView этой статьи

❗️ <b>ВАЖНО</b>:
1) Если пользователь не указал никаких тегов, то ему будут присылаться все статьи
2) Если тег содержит пробелы, то они должны быть заменены на нижнее подчёркивание (Разработка под Windows -> Разработка_под_Windows)
3) Теги, которые должны быть удалены или добавлены, записываются через пробел

📌 <b>УТОЧНЕНИЯ</b>:
1) Хаб == Тег 
2) Регистр тегов не важен
3) При обнаружении 🐞 багов – писать @Tirsias
`

const botFatherCommands =
`
help - показать помощь
my_tags - показать список тегов
add_tags - добавить теги
del_tags - удалить теги
del_all_tags - удалить ВСЕ теги
copy_tags - скопировать теги из профиля на habrahabr'e
stop - приостановить рассылку
get_best - получить лучшие статьи за день
`