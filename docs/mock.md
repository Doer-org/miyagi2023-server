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

## coupon API

#### クーポン取得

```bash
curl --location --request GET 'localhost:8080/coupons/1'
```

response
```json
{
    "id": "e91ec300-be69-497c-ab6e-14ca2549fd3b",
    "name": "hoge mock",
    "expiration_date": 100,
    "discount_rate": 10,
    "created_at": "2023-03-02 22:13:27.341234965 +0000 UTC m=+0.019154584",
    "spot": {
        "id": "49cd0066-6373-4abd-b778-ffd3dbb13a80",
        "name": "hoge mock",
        "detail": "hogehoge",
        "like": 23,
        "created_at": "2023-03-02 22:13:27.34123409 +0000 UTC m=+0.019153709"
    }
}
```

<br>

#### クーポン登録

```bash
curl --location --request POST 'localhost:8080/coupons' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"hoge"
}'
```

```json
{
    "id": "e91ec300-be69-497c-ab6e-14ca2549fd3b",
    "name": "hoge mock",
    "expiration_date": 100,
    "discount_rate": 10,
    "created_at": "2023-03-02 22:13:27.341234965 +0000 UTC m=+0.019154584",
    "spot": {
        "id": "49cd0066-6373-4abd-b778-ffd3dbb13a80",
        "name": "hoge mock",
        "detail": "hogehoge",
        "like": 23,
        "created_at": "2023-03-02 22:13:27.34123409 +0000 UTC m=+0.019153709"
    }
}
```

#### クーポン一覧取得


```bash
curl --location --request GET 'localhost:8080/coupons/list'
```

response
```json
{
    "coupons": [
        {
            "id": "d5ee4cc5-b51c-406c-a49e-0cb313ef5ef8",
            "name": "hoge1 mock",
            "expiration_date": 100,
            "discount_rate": 10,
            "created_at": "2023-03-02 22:13:27.341235757 +0000 UTC m=+0.019155376",
            "spot": {
                "id": "49cd0066-6373-4abd-b778-ffd3dbb13a80",
                "name": "hoge mock",
                "detail": "hogehoge",
                "like": 23,
                "created_at": "2023-03-02 22:13:27.34123409 +0000 UTC m=+0.019153709"
            }
        },
        {
            "id": "643f5511-219d-4240-923e-161a600a1ee1",
            "name": "hoge2 mock",
            "expiration_date": 200,
            "discount_rate": 20,
            "created_at": "2023-03-02 22:13:27.341236507 +0000 UTC m=+0.019156126",
            "spot": {
                "id": "49cd0066-6373-4abd-b778-ffd3dbb13a80",
                "name": "hoge mock",
                "detail": "hogehoge",
                "like": 23,
                "created_at": "2023-03-02 22:13:27.34123409 +0000 UTC m=+0.019153709"
            }
        },
        {
            "id": "98233980-bd65-4ea2-893d-424fdd44a180",
            "name": "hoge3 mock",
            "expiration_date": 300,
            "discount_rate": 30,
            "created_at": "2023-03-02 22:13:27.341237215 +0000 UTC m=+0.019156792",
            "spot": {
                "id": "49cd0066-6373-4abd-b778-ffd3dbb13a80",
                "name": "hoge mock",
                "detail": "hogehoge",
                "like": 23,
                "created_at": "2023-03-02 22:13:27.34123409 +0000 UTC m=+0.019153709"
            }
        }
    ]
}
```

<br>

## coupon status API

#### 発行済みクーポンを取得

```bash
curl --location --request GET 'localhost:8080/coupon_statuses/1'
```

response
```json
{
    "id": "ad0a854d-0153-434f-b5a1-ac2f3e0db3e8",
    "used_flg": false,
    "created_at": "2023-03-02 22:17:43.782622501 +0000 UTC m=+0.000896417",
    "updated_at": "2023-03-02 22:17:43.782622584 +0000 UTC m=+0.000896459",
    "coupon": {
        "id": "a02008e9-7224-4126-a50b-30daadd29c41",
        "name": "hoge mock",
        "expiration_date": 100,
        "discount_rate": 10,
        "created_at": "2023-03-02 22:17:43.782616334 +0000 UTC m=+0.000890251",
        "spot": {
            "id": "07658978-8d0e-40c7-b78a-08f4cefc2c5a",
            "name": "hoge mock",
            "detail": "hogehoge",
            "like": 23,
            "created_at": "2023-03-02 22:17:43.782615584 +0000 UTC m=+0.000889459"
        }
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
        "created_at": "2023-03-02 22:17:43.782621959 +0000 UTC m=+0.000895834"
    }
}
```

<br>

#### クーポンを発行する

```bash
curl --location --request POST 'localhost:8080/coupon_statuses' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"hoge"
}'
```

response
```json
{
    "id": "3867b44b-c61d-4a17-859e-5703fd195901",
    "used_flg": false,
    "created_at": "2023-03-02 22:27:44.090949584 +0000 UTC m=+0.023213085",
    "updated_at": "2023-03-02 22:27:44.090949667 +0000 UTC m=+0.023213126",
    "coupon": {
        "id": "1c9bebab-e480-41cf-825b-c4e2ead35520",
        "name": "hoge mock",
        "expiration_date": 100,
        "discount_rate": 10,
        "created_at": "2023-03-02 22:27:44.090941834 +0000 UTC m=+0.023205335",
        "spot": {
            "id": "f5d22baa-2745-4ffd-a3bf-22283c5d370b",
            "name": "hoge mock",
            "detail": "hogehoge",
            "like": 23,
            "created_at": "2023-03-02 22:27:44.090940959 +0000 UTC m=+0.023204460"
        }
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
        "created_at": "2023-03-02 22:27:44.090949084 +0000 UTC m=+0.023212543"
    }
}
```

<br>



#### 発行されたクーポンを使用済みに更新する

以下の場合、
クーポンステータスのidが1でuserのidが1のクーポンのステータスが更新される

```bash
curl --location --request PUT 'localhost:8080/coupon_statuses/1/users/1/use'
```

response
```json
{
    "id": "21ed194e-cf48-4849-bbd1-a9a5ce84ff3f",
    "used_flg": false,
    "created_at": "2023-03-02 22:33:53.283142213 +0000 UTC m=+0.012966918",
    "updated_at": "2023-03-02 22:33:53.283142297 +0000 UTC m=+0.012966960",
    "coupon": {
        "id": "dc5cfff6-3fa3-41c2-9c2e-95cdc43211c1",
        "name": "hoge mock",
        "expiration_date": 100,
        "discount_rate": 10,
        "created_at": "2023-03-02 22:33:53.283134422 +0000 UTC m=+0.012959085",
        "spot": {
            "id": "2e74b170-1d05-4e9c-8776-69ee60a39260",
            "name": "hoge mock",
            "detail": "hogehoge",
            "like": 23,
            "created_at": "2023-03-02 22:33:53.283133547 +0000 UTC m=+0.012958251"
        }
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
        "created_at": "2023-03-02 22:33:53.28314163 +0000 UTC m=+0.012966335"
    }
}
```


<br>

#### 特定のユーザーのクーポンのステータスの一覧を取得

```bash
curl --location --request GET 'localhost:8080/coupon_statuses/list/users/1'
```

(memo: userの情報が何度も入って冗長なので対応したい)

response
```json
{
    "coupon_statuses": [
        {
            "id": "e6ae4c65-1f6c-4bab-a57b-4712b9042b02",
            "used_flg": false,
            "created_at": "2023-03-02 22:20:29.118659175 +0000 UTC m=+0.021894085",
            "updated_at": "2023-03-02 22:20:29.118659216 +0000 UTC m=+0.021894168",
            "coupon": {
                "id": "0aefb3f1-c5eb-4fa6-aff7-81af3ed42c57",
                "name": "hoge mock",
                "expiration_date": 100,
                "discount_rate": 10,
                "created_at": "2023-03-02 22:20:29.118652133 +0000 UTC m=+0.021887043",
                "spot": {
                    "id": "a5cbfbf6-d2e2-4bad-b299-6959a851928d",
                    "name": "hoge mock",
                    "detail": "hogehoge",
                    "like": 23,
                    "created_at": "2023-03-02 22:20:29.118651466 +0000 UTC m=+0.021886418"
                }
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
                "created_at": "2023-03-02 22:20:29.11865805 +0000 UTC m=+0.021892960"
            }
        },
        {
            "id": "881a2816-bd49-4c21-bf37-264f7cfd046c",
            "used_flg": true,
            "created_at": "2023-03-02 22:20:29.118659758 +0000 UTC m=+0.021894668",
            "updated_at": "2023-03-02 22:20:29.1186598 +0000 UTC m=+0.021894710",
            "coupon": {
                "id": "0aefb3f1-c5eb-4fa6-aff7-81af3ed42c57",
                "name": "hoge mock",
                "expiration_date": 100,
                "discount_rate": 10,
                "created_at": "2023-03-02 22:20:29.118652133 +0000 UTC m=+0.021887043",
                "spot": {
                    "id": "a5cbfbf6-d2e2-4bad-b299-6959a851928d",
                    "name": "hoge mock",
                    "detail": "hogehoge",
                    "like": 23,
                    "created_at": "2023-03-02 22:20:29.118651466 +0000 UTC m=+0.021886418"
                }
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
                "created_at": "2023-03-02 22:20:29.11865805 +0000 UTC m=+0.021892960"
            }
        }
    ]
}
```

<br>