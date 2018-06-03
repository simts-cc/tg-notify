# Telegram Notification

## Installation
```bash
git clone https://github.com/stanlyliao/tg-notify.git
```

## Package Management for Golang: 
```bash
glide install
```

## Clone the environmental file
```bash
cp .env.example .env
```

## Environmental variables
```bash
SERVER_PORT=8080

TG_BOT_TOKEN=
# private channel
TG_CHAN_ID1=-100000000
# public channel
TG_CHAN_ID2=@tg_notify_id

DB_WRITER=127.0.0.1
DB_READER=127.0.0.1
DB_PORT=3306
DB_NAME=al_notify
DB_USERNAME=root
DB_PASSWORD=root
DB_CHARSET=utf8
DB_MAX_IDLE=32
DB_MAX_OPEN=32
```

## Database Schema
```sql
CREATE TABLE `api_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uri` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL,
  `req_data` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL,
  `res_data` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL,
  `headers` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

## Interface
#### URL
> [http://localhost:8080/v1/sendMessage](http://localhost:8080/v1/sendMessage)

#### Format
> JSON

#### HTTP Request
> POST

#### Request Body
|Field|Required|Type|Explain|
|:----- |:-------|:-----|----- |
|code |true |string|env TG_CHAN_`XXX` |
|message |true |string ||
|slient |false |bool |slient status|

#### Response Body

|Field|Type|Explain |
|:----- |:------|:----------------------------- |
|ok | bool |success status |

#### Example

> URL：[http://localhost:8080/v1/message](http://localhost:8080/v1/message)

```json
{
    "code": "ID1",
    "messsage": "test",
    "slient": true
}
```