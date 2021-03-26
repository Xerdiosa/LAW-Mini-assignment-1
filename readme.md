# Mini assignment 1

0. pastikan dependensi terpenuhi dengan:

``` bash
go get
```

1. jalankan server dengan:

``` bash
go run main.go
```

2. dapatkan token dengan:  

``` curl
curl --location --request POST 'localhost:8080/oauth/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'client_id=87647232352' \
--data-urlencode 'client_secret=108358321' \
--data-urlencode 'username=akbar' \
--data-urlencode 'password=password'
```

3. coba tembak resource dengan:

``` curl
curl --location --request POST 'localhost:8080/oauth/resource' \
--header 'Authorization: Bearer {token yang tadi di dapat}'
```

4. uji coba dengan:

``` curl
bombardier --method=POST \
-H "Authorization: Bearer {token yang tadi di dapat}" \
-H "Content-Type: application/x-www-form-urlencoded" \
-n 10000 localhost:8080/oauth/resource
```
