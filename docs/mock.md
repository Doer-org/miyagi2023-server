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
    "img_url": "http://example.com",
    "prefecture": "TOYAMA",
    "created_at": "2023-03-03 03:18:21.503441044 +0000 UTC m=+0.015616293"
}
```

<br>

#### ユーザー登録

mockAPIではrequest bodyに何を入れても作成できます。
ただし、request body自体が空だとエラーがでます。


```bash
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"hoge",
    "name":"hoge",
    "age":22,
    "gender":"MEN",
    "birthday":"2000-04-29",
    "address":"hoge",
    "img_url":"http://example.com",
    "prefecture":"TOYAMA"
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
    "img_url": "http://example.com",
    "prefecture": "TOYAMA",
    "created_at": "2023-03-03 03:59:30.576256839 +0000 UTC m=+0.016292584"
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
    "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
    "name": "hoge mock",
    "detail": "hogehoge",
    "like": 23,
    "img_url": "http://example.com",
    "address": "hoge",
    "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
}
```

<br>

#### 観光地登録

```bash
curl --location --request POST 'localhost:8080/spots' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"hoge",
    "detail":"hoge",
    "address":"hoge",
    "img_url":"http://example.com"
}'
```

response
```json
{
    "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
    "name": "hoge mock",
    "detail": "hogehoge",
    "like": 23,
    "img_url": "http://example.com",
    "address": "hoge",
    "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
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
            "id": "d5cb979f-acff-4ffa-9170-9b6ccadd9464",
            "name": "hoge1 mock",
            "detail": "hogehoge1",
            "like": 123,
            "img_url": "http://example.com",
            "address": "hoge",
            "created_at": "2023-03-03 03:59:30.576252298 +0000 UTC m=+0.016288042"
        },
        {
            "id": "52090406-261a-4531-87b9-d1619520438f",
            "name": "hoge2 mock",
            "detail": "hogehoge2",
            "like": 223,
            "img_url": "http://example.com",
            "address": "hoge",
            "created_at": "2023-03-03 03:59:30.576252881 +0000 UTC m=+0.016288626"
        },
        {
            "id": "be2d1c94-24cd-439c-8639-e0b9a58d21c8",
            "name": "hoge3 mock",
            "detail": "hogehoge3",
            "like": 323,
            "img_url": "http://example.com",
            "address": "hoge",
            "created_at": "2023-03-03 03:59:30.576253423 +0000 UTC m=+0.016289167"
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
    "spot_id":"hoge",
    "user_id":"hoge"
}
'
```

response
```json
{
    "id": "970e9421-8bd2-439b-85d6-0cbc853ad167",
    "spot": {
        "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
        "name": "hoge mock",
        "detail": "hogehoge",
        "like": 23,
        "img_url": "http://example.com",
        "address": "hoge",
        "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
    },
    "user": {
        "id": "0000",
        "name": "mahiro mock",
        "age": 22,
        "gender": "MEN",
        "birthday": "2000-04-29 00:00:00 +0000 UTC",
        "address": "hoge",
        "img_url": "http://example.com",
        "prefecture": "TOYAMA",
        "created_at": "2023-03-03 03:59:30.576256839 +0000 UTC m=+0.016292584"
    },
    "coupon": {
        "id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
        "name": "hoge mock",
        "expiration_date": 100,
        "discount_rate": 10,
        "created_at": "2023-03-03 03:59:30.576250089 +0000 UTC m=+0.016285792",
        "spot": {
            "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
            "name": "hoge mock",
            "detail": "hogehoge",
            "like": 23,
            "img_url": "http://example.com",
            "address": "hoge",
            "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
        }
    },
    "created_at": "2023-03-03 03:59:30.576259381 +0000 UTC m=+0.016295084"
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
            "id": "15295f26-aeea-4338-a76a-e254f6913571",
            "spot_id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
            "user_id": "0000",
            "coupon_id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
            "created_at": "2023-03-03 03:59:30.576259923 +0000 UTC m=+0.016295626"
        },
        {
            "id": "6db3cba9-28ba-4909-bf0e-78c09138780c",
            "spot_id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
            "user_id": "0000",
            "coupon_id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
            "created_at": "2023-03-03 03:59:30.576260423 +0000 UTC m=+0.016296167"
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
    "name":"hoge",
    "expiration_date":10,
    "discount_rate":10,
    "spot_id":"hoge"
}
'
```

```json
{
    "id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
    "name": "hoge mock",
    "expiration_date": 100,
    "discount_rate": 10,
    "created_at": "2023-03-03 03:59:30.576250089 +0000 UTC m=+0.016285792",
    "spot": {
        "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
        "name": "hoge mock",
        "detail": "hogehoge",
        "like": 23,
        "img_url": "http://example.com",
        "address": "hoge",
        "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
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
            "id": "581d6d96-435c-4259-b17f-89972b8efc0f",
            "name": "hoge1 mock",
            "expiration_date": 100,
            "discount_rate": 10,
            "created_at": "2023-03-03 03:59:30.576250673 +0000 UTC m=+0.016286376",
            "spot": {
                "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
                "name": "hoge mock",
                "detail": "hogehoge",
                "like": 23,
                "img_url": "http://example.com",
                "address": "hoge",
                "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
            }
        },
        {
            "id": "f4d53b4c-08c4-43da-96f4-6f94ee3fb846",
            "name": "hoge2 mock",
            "expiration_date": 200,
            "discount_rate": 20,
            "created_at": "2023-03-03 03:59:30.576251256 +0000 UTC m=+0.016286959",
            "spot": {
                "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
                "name": "hoge mock",
                "detail": "hogehoge",
                "like": 23,
                "img_url": "http://example.com",
                "address": "hoge",
                "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
            }
        },
        {
            "id": "74dcb2ad-cb83-4113-b539-dd4eac471402",
            "name": "hoge3 mock",
            "expiration_date": 300,
            "discount_rate": 30,
            "created_at": "2023-03-03 03:59:30.576251756 +0000 UTC m=+0.016287501",
            "spot": {
                "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
                "name": "hoge mock",
                "detail": "hogehoge",
                "like": 23,
                "img_url": "http://example.com",
                "address": "hoge",
                "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
            }
        }
    ]
}
```

<br>

## coupon status API

#### クーポンのステータスを取得

```bash
curl --location --request GET 'localhost:8080/coupon_statuses/1'
```

response
```json
{
    "id": "772a0451-d240-4821-b3db-51bc9faa8569",
    "used_flg": false,
    "created_at": "2023-03-03 03:59:30.576257464 +0000 UTC m=+0.016293209",
    "updated_at": "2023-03-03 03:59:30.576257506 +0000 UTC m=+0.016293251",
    "coupon": {
        "id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
        "name": "hoge mock",
        "expiration_date": 100,
        "discount_rate": 10,
        "created_at": "2023-03-03 03:59:30.576250089 +0000 UTC m=+0.016285792",
        "spot": {
            "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
            "name": "hoge mock",
            "detail": "hogehoge",
            "like": 23,
            "img_url": "http://example.com",
            "address": "hoge",
            "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
        }
    },
    "user": {
        "id": "0000",
        "name": "mahiro mock",
        "age": 22,
        "gender": "MEN",
        "birthday": "2000-04-29 00:00:00 +0000 UTC",
        "address": "hoge",
        "img_url": "http://example.com",
        "prefecture": "TOYAMA",
        "created_at": "2023-03-03 03:59:30.576256839 +0000 UTC m=+0.016292584"
    }
}
```

<br>

#### クーポンのステータスを登録する

```bash
curl --location --request POST 'localhost:8080/coupon_statuses' \
--header 'Content-Type: application/json' \
--data-raw '{
    "coupon_id":"hoge",
    "user_id":"hoge"
}
'
```

response
```json
{
    "id": "772a0451-d240-4821-b3db-51bc9faa8569",
    "used_flg": false,
    "created_at": "2023-03-03 03:59:30.576257464 +0000 UTC m=+0.016293209",
    "updated_at": "2023-03-03 03:59:30.576257506 +0000 UTC m=+0.016293251",
    "coupon": {
        "id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
        "name": "hoge mock",
        "expiration_date": 100,
        "discount_rate": 10,
        "created_at": "2023-03-03 03:59:30.576250089 +0000 UTC m=+0.016285792",
        "spot": {
            "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
            "name": "hoge mock",
            "detail": "hogehoge",
            "like": 23,
            "img_url": "http://example.com",
            "address": "hoge",
            "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
        }
    },
    "user": {
        "id": "0000",
        "name": "mahiro mock",
        "age": 22,
        "gender": "MEN",
        "birthday": "2000-04-29 00:00:00 +0000 UTC",
        "address": "hoge",
        "img_url": "http://example.com",
        "prefecture": "TOYAMA",
        "created_at": "2023-03-03 03:59:30.576256839 +0000 UTC m=+0.016292584"
    }
}
```

<br>



#### 指定したクーポンのステータスを使用済みに更新する

以下の場合、
クーポンステータスのidが1でuserのidが1のクーポンのステータスが更新される

```bash
curl --location --request PUT 'localhost:8080/coupon_statuses/1/users/1/use'
```

response
```json
{
    "id": "772a0451-d240-4821-b3db-51bc9faa8569",
    "used_flg": false,
    "created_at": "2023-03-03 03:59:30.576257464 +0000 UTC m=+0.016293209",
    "updated_at": "2023-03-03 03:59:30.576257506 +0000 UTC m=+0.016293251",
    "coupon": {
        "id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
        "name": "hoge mock",
        "expiration_date": 100,
        "discount_rate": 10,
        "created_at": "2023-03-03 03:59:30.576250089 +0000 UTC m=+0.016285792",
        "spot": {
            "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
            "name": "hoge mock",
            "detail": "hogehoge",
            "like": 23,
            "img_url": "http://example.com",
            "address": "hoge",
            "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
        }
    },
    "user_id": "0000"
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
            "id": "3fbafa9b-a454-43cb-85dd-c8052dbf7a88",
            "used_flg": false,
            "created_at": "2023-03-03 03:59:30.576258048 +0000 UTC m=+0.016293792",
            "updated_at": "2023-03-03 03:59:30.576258089 +0000 UTC m=+0.016293834",
            "coupon": {
                "id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
                "name": "hoge mock",
                "expiration_date": 100,
                "discount_rate": 10,
                "created_at": "2023-03-03 03:59:30.576250089 +0000 UTC m=+0.016285792",
                "spot": {
                    "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
                    "name": "hoge mock",
                    "detail": "hogehoge",
                    "like": 23,
                    "img_url": "http://example.com",
                    "address": "hoge",
                    "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
                }
            },
            "user_id": "0000"
        },
        {
            "id": "21ecb9ab-59c1-46f8-b4ae-871ca415c7d5",
            "used_flg": true,
            "created_at": "2023-03-03 03:59:30.576258756 +0000 UTC m=+0.016294459",
            "updated_at": "2023-03-03 03:59:30.576258798 +0000 UTC m=+0.016294501",
            "coupon": {
                "id": "d5fec20c-fccf-42c1-bc16-aec8351426be",
                "name": "hoge mock",
                "expiration_date": 100,
                "discount_rate": 10,
                "created_at": "2023-03-03 03:59:30.576250089 +0000 UTC m=+0.016285792",
                "spot": {
                    "id": "195fb4fe-9700-49ce-80a0-111b91a1e1e4",
                    "name": "hoge mock",
                    "detail": "hogehoge",
                    "like": 23,
                    "img_url": "http://example.com",
                    "address": "hoge",
                    "created_at": "2023-03-03 03:59:30.576249381 +0000 UTC m=+0.016285084"
                }
            },
            "user_id": "0000"
        }
    ]
}
```

<br>