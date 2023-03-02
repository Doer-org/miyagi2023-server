# mock API

## user API

#### ユーザー取得

idを指定してユーザーを取得する
(モックなのでidによらず基本同じ値が返ってきます)

```bash
curl --location --request GET 'localhost:8080/users/1'
```

response
```json
{
    "id": "0000",
    "name": "mahiro mock",
    "age": 22,
    "gender": "MEN",
    "birthday": "2000-04-29 00:00:00 +0000 UTC",
    "address": "hoge",
    "profile_img": "http://example.com",
    "prefecture": "TOYAMA",
    "created_at": "2023-03-02 13:46:40.957913336 +0000 UTC m=+0.020202792"
}
```

<br>

#### ユーザー登録

mockAPIではrequest bodyに何を入れても作成できます。
ただし、request body自体が空だとエラーがでます。


```bash
curl --location --request POST 'localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"hoge"
}'
```

response(success)
```json
{
    "id": "0000",
    "name": "mahiro mock",
    "age": 22,
    "gender": "MEN",
    "birthday": "2000-04-29 00:00:00 +0000 UTC",
    "address": "hoge",
    "profile_img": "http://example.com",
    "prefecture": "TOYAMA",
    "created_at": "2023-03-02 13:46:40.957913336 +0000 UTC m=+0.020202792"
}
```

response(error)
```json
{
    "message": "error: content length is 0"
}
```

<br>

## spot API

観光地に関するAPIです

#### 観光地取得

指定したidの観光地を取得します
```bash
curl --location --request GET 'localhost:8080/spots/1'
```

response
```json
{
    "id": "aa70470e-32e2-4552-8b56-37c72e11aa58",
    "name": "hoge mock",
    "detail": "hogehoge",
    "like": 23,
    "created_at": "2023-03-02 13:46:40.957908961 +0000 UTC m=+0.020198417"
}
```

<br>

#### 観光地登録

```bash
curl --location --request POST 'localhost:8080/spots' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"hoge"
}'
```

response
```json
{
    "id": "e50fc190-bf24-4c85-b2dc-65bac8fbb543",
    "name": "hoge mock",
    "detail": "hogehoge",
    "like": 23,
    "created_at": "2023-03-02 13:56:00.344650761 +0000 UTC m=+559.403534423"
}
```


<br>


#### 観光地一覧取得

```bash
curl --location --request GET 'localhost:8080/spots/list'
```

response
```json
{
    "spots": [
        {
            "id": "d9a4df51-da68-40a9-82db-7f16c16ce30f",
            "name": "hoge1 mock",
            "detail": "hogehoge1",
            "like": 123,
            "created_at": "2023-03-02 13:46:40.957909669 +0000 UTC m=+0.020199084"
        },
        {
            "id": "701ff9e0-2444-4207-bd97-6d96438dbc7c",
            "name": "hoge2 mock",
            "detail": "hogehoge2",
            "like": 223,
            "created_at": "2023-03-02 13:46:40.957910377 +0000 UTC m=+0.020199834"
        },
        {
            "id": "3424ff92-cf4c-45cf-a064-82343208cc9c",
            "name": "hoge3 mock",
            "detail": "hogehoge3",
            "like": 323,
            "created_at": "2023-03-02 13:46:40.957910961 +0000 UTC m=+0.020200417"
        }
    ]
}
```

<br>

## stamp card API

#### スタンプカード取得

```bash
curl --location --request GET 'localhost:8080/stamp_cards/1'
```

response
```json
{
    "id": "1bf86f8f-60d7-40f5-a5de-5c596c10140d",
    "name": "hoge mock",
    "created_at": "2023-03-02 13:46:40.957911502 +0000 UTC m=+0.020200959",
    "spots": [
        {
            "id": "d9a4df51-da68-40a9-82db-7f16c16ce30f",
            "name": "hoge1 mock",
            "detail": "hogehoge1",
            "like": 123,
            "created_at": "2023-03-02 13:46:40.957909669 +0000 UTC m=+0.020199084"
        },
        {
            "id": "701ff9e0-2444-4207-bd97-6d96438dbc7c",
            "name": "hoge2 mock",
            "detail": "hogehoge2",
            "like": 223,
            "created_at": "2023-03-02 13:46:40.957910377 +0000 UTC m=+0.020199834"
        },
        {
            "id": "3424ff92-cf4c-45cf-a064-82343208cc9c",
            "name": "hoge3 mock",
            "detail": "hogehoge3",
            "like": 323,
            "created_at": "2023-03-02 13:46:40.957910961 +0000 UTC m=+0.020200417"
        }
    ]
}
```

<br>


#### スタンプカード登録

```bash
curl --location --request POST 'localhost:8080/stamp_cards/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"hoge"
}'
```

response
```json
{
    "id": "1bf86f8f-60d7-40f5-a5de-5c596c10140d",
    "name": "hoge mock",
    "created_at": "2023-03-02 13:46:40.957911502 +0000 UTC m=+0.020200959",
    "spots": [
        {
            "id": "d9a4df51-da68-40a9-82db-7f16c16ce30f",
            "name": "hoge1 mock",
            "detail": "hogehoge1",
            "like": 123,
            "created_at": "2023-03-02 13:46:40.957909669 +0000 UTC m=+0.020199084"
        },
        {
            "id": "701ff9e0-2444-4207-bd97-6d96438dbc7c",
            "name": "hoge2 mock",
            "detail": "hogehoge2",
            "like": 223,
            "created_at": "2023-03-02 13:46:40.957910377 +0000 UTC m=+0.020199834"
        },
        {
            "id": "3424ff92-cf4c-45cf-a064-82343208cc9c",
            "name": "hoge3 mock",
            "detail": "hogehoge3",
            "like": 323,
            "created_at": "2023-03-02 13:46:40.957910961 +0000 UTC m=+0.020200417"
        }
    ]
}
```

<br>



#### スタンプカード一覧取得

```bash
curl --location --request GET 'localhost:8080/stamp_cards/list'
```

response
```json
{
    "stamp_cards": [
        {
            "id": "c0f8064d-9116-46e3-b8a8-afd4498e8f78",
            "name": "hoge1 mock",
            "created_at": "2023-03-02 13:46:40.957912211 +0000 UTC m=+0.020201626",
            "spots": [
                {
                    "id": "d9a4df51-da68-40a9-82db-7f16c16ce30f",
                    "name": "hoge1 mock",
                    "detail": "hogehoge1",
                    "like": 123,
                    "created_at": "2023-03-02 13:46:40.957909669 +0000 UTC m=+0.020199084"
                },
                {
                    "id": "701ff9e0-2444-4207-bd97-6d96438dbc7c",
                    "name": "hoge2 mock",
                    "detail": "hogehoge2",
                    "like": 223,
                    "created_at": "2023-03-02 13:46:40.957910377 +0000 UTC m=+0.020199834"
                },
                {
                    "id": "3424ff92-cf4c-45cf-a064-82343208cc9c",
                    "name": "hoge3 mock",
                    "detail": "hogehoge3",
                    "like": 323,
                    "created_at": "2023-03-02 13:46:40.957910961 +0000 UTC m=+0.020200417"
                }
            ]
        },
        {
            "id": "6391bfb2-81cc-480c-a586-af87e8a56f2c",
            "name": "hoge2 mock",
            "created_at": "2023-03-02 13:46:40.957912752 +0000 UTC m=+0.020202209",
            "spots": [
                {
                    "id": "d9a4df51-da68-40a9-82db-7f16c16ce30f",
                    "name": "hoge1 mock",
                    "detail": "hogehoge1",
                    "like": 123,
                    "created_at": "2023-03-02 13:46:40.957909669 +0000 UTC m=+0.020199084"
                },
                {
                    "id": "701ff9e0-2444-4207-bd97-6d96438dbc7c",
                    "name": "hoge2 mock",
                    "detail": "hogehoge2",
                    "like": 223,
                    "created_at": "2023-03-02 13:46:40.957910377 +0000 UTC m=+0.020199834"
                },
                {
                    "id": "3424ff92-cf4c-45cf-a064-82343208cc9c",
                    "name": "hoge3 mock",
                    "detail": "hogehoge3",
                    "like": 323,
                    "created_at": "2023-03-02 13:46:40.957910961 +0000 UTC m=+0.020200417"
                }
            ]
        }
    ]
}
```


<br>

## stamp log API

#### スタンプログ保存

```bash
curl --location --request POST 'localhost:8080/stamp_logs' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"hoge"
}'
```

response
```json
{
    "id": "35e308d0-7086-43a2-9bbf-cecae9f7d3d1",
    "spot": {
        "id": "aa70470e-32e2-4552-8b56-37c72e11aa58",
        "name": "hoge mock",
        "detail": "hogehoge",
        "like": 23,
        "created_at": "2023-03-02 13:46:40.957908961 +0000 UTC m=+0.020198417"
    },
    "user": {
        "id": "0000",
        "name": "mahiro mock",
        "age": 22,
        "gender": "MEN",
        "birthday": "2000-04-29 00:00:00 +0000 UTC",
        "address": "hoge",
        "profile_img": "http://example.com",
        "prefecture": "TOYAMA",
        "created_at": "2023-03-02 13:46:40.957913336 +0000 UTC m=+0.020202792"
    },
    "stamp_card": {
        "id": "1bf86f8f-60d7-40f5-a5de-5c596c10140d",
        "name": "hoge mock",
        "created_at": "2023-03-02 13:46:40.957911502 +0000 UTC m=+0.020200959",
        "spots": [
            {
                "id": "d9a4df51-da68-40a9-82db-7f16c16ce30f",
                "name": "hoge1 mock",
                "detail": "hogehoge1",
                "like": 123,
                "created_at": "2023-03-02 13:46:40.957909669 +0000 UTC m=+0.020199084"
            },
            {
                "id": "701ff9e0-2444-4207-bd97-6d96438dbc7c",
                "name": "hoge2 mock",
                "detail": "hogehoge2",
                "like": 223,
                "created_at": "2023-03-02 13:46:40.957910377 +0000 UTC m=+0.020199834"
            },
            {
                "id": "3424ff92-cf4c-45cf-a064-82343208cc9c",
                "name": "hoge3 mock",
                "detail": "hogehoge3",
                "like": 323,
                "created_at": "2023-03-02 13:46:40.957910961 +0000 UTC m=+0.020200417"
            }
        ]
    },
    "created_at": "2023-03-02 13:46:40.957917336 +0000 UTC m=+0.020206792"
}
```

<br>

#### スタンプログ一覧取得

```bash
curl --location --request GET 'localhost:8080/stamp_logs/list'
```

response
```json
{
    "stamp_logs": [
        {
            "id": "be9bd4d0-3dfe-4d12-ac69-71d832671ac9",
            "spot": {
                "id": "aa70470e-32e2-4552-8b56-37c72e11aa58",
                "name": "hoge mock",
                "detail": "hogehoge",
                "like": 23,
                "created_at": "2023-03-02 13:46:40.957908961 +0000 UTC m=+0.020198417"
            },
            "user": {
                "id": "0000",
                "name": "mahiro mock",
                "age": 22,
                "gender": "MEN",
                "birthday": "2000-04-29 00:00:00 +0000 UTC",
                "address": "hoge",
                "profile_img": "http://example.com",
                "prefecture": "TOYAMA",
                "created_at": "2023-03-02 13:46:40.957913336 +0000 UTC m=+0.020202792"
            },
            "stamp_card": {
                "id": "1bf86f8f-60d7-40f5-a5de-5c596c10140d",
                "name": "hoge mock",
                "created_at": "2023-03-02 13:46:40.957911502 +0000 UTC m=+0.020200959",
                "spots": [
                    {
                        "id": "d9a4df51-da68-40a9-82db-7f16c16ce30f",
                        "name": "hoge1 mock",
                        "detail": "hogehoge1",
                        "like": 123,
                        "created_at": "2023-03-02 13:46:40.957909669 +0000 UTC m=+0.020199084"
                    },
                    {
                        "id": "701ff9e0-2444-4207-bd97-6d96438dbc7c",
                        "name": "hoge2 mock",
                        "detail": "hogehoge2",
                        "like": 223,
                        "created_at": "2023-03-02 13:46:40.957910377 +0000 UTC m=+0.020199834"
                    },
                    {
                        "id": "3424ff92-cf4c-45cf-a064-82343208cc9c",
                        "name": "hoge3 mock",
                        "detail": "hogehoge3",
                        "like": 323,
                        "created_at": "2023-03-02 13:46:40.957910961 +0000 UTC m=+0.020200417"
                    }
                ]
            },
            "created_at": "2023-03-02 13:46:40.957917877 +0000 UTC m=+0.020207334"
        },
        {
            "id": "37630b1b-6973-4a31-9bb8-b3bde129d670",
            "spot": {
                "id": "aa70470e-32e2-4552-8b56-37c72e11aa58",
                "name": "hoge mock",
                "detail": "hogehoge",
                "like": 23,
                "created_at": "2023-03-02 13:46:40.957908961 +0000 UTC m=+0.020198417"
            },
            "user": {
                "id": "0000",
                "name": "mahiro mock",
                "age": 22,
                "gender": "MEN",
                "birthday": "2000-04-29 00:00:00 +0000 UTC",
                "address": "hoge",
                "profile_img": "http://example.com",
                "prefecture": "TOYAMA",
                "created_at": "2023-03-02 13:46:40.957913336 +0000 UTC m=+0.020202792"
            },
            "stamp_card": {
                "id": "1bf86f8f-60d7-40f5-a5de-5c596c10140d",
                "name": "hoge mock",
                "created_at": "2023-03-02 13:46:40.957911502 +0000 UTC m=+0.020200959",
                "spots": [
                    {
                        "id": "d9a4df51-da68-40a9-82db-7f16c16ce30f",
                        "name": "hoge1 mock",
                        "detail": "hogehoge1",
                        "like": 123,
                        "created_at": "2023-03-02 13:46:40.957909669 +0000 UTC m=+0.020199084"
                    },
                    {
                        "id": "701ff9e0-2444-4207-bd97-6d96438dbc7c",
                        "name": "hoge2 mock",
                        "detail": "hogehoge2",
                        "like": 223,
                        "created_at": "2023-03-02 13:46:40.957910377 +0000 UTC m=+0.020199834"
                    },
                    {
                        "id": "3424ff92-cf4c-45cf-a064-82343208cc9c",
                        "name": "hoge3 mock",
                        "detail": "hogehoge3",
                        "like": 323,
                        "created_at": "2023-03-02 13:46:40.957910961 +0000 UTC m=+0.020200417"
                    }
                ]
            },
            "created_at": "2023-03-02 13:46:40.957918919 +0000 UTC m=+0.020208334"
        }
    ]
}
```

<br>