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

> URL：[http://localhost:8080/v1/sendMessage](http://localhost:8080/v1/sendMessage)

```json
{
    "code": "ID1",
    "messsage": "test",
    "slient": true
}
```