# CURL

## Create Dish
```
curl http://localhost:8080/api/v1/dishes \
    --include \
    --header "Content-Type: application/json" \
    --header "RestaurantID: e814691f-b53e-45c4-8253-e2f2a7f5ff35" \
    --request "POST" \
    --data '{"name": "Pasta","description": "Nice","price": 49, "restaurantID": "e814691f-b53e-45c4-8253-e2f2a7f5ff35"}'
```

## Get Dish
```
curl http://localhost:8080/api/v1/dishes/d6262803-770a-4dd1-b680-e0a932d6e636 \
    --header "Content-Type: application/json" \
    --header "RestaurantID: e814691f-b53e-45c4-8253-e2f2a7f5ff35"
```

## List Dish
```
curl http://localhost:8080/api/v1/dishes\
    --header "Content-Type: application/json" \
    --header "RestaurantID: e814691f-b53e-45c4-8253-e2f2a7f5ff35"
```

## Update Dish
```
curl http://localhost:8080/api/v1/dishes \
    --include \
    --header "Content-Type: application/json" \
    --header "RestaurantID: e814691f-b53e-45c4-8253-e2f2a7f5ff35" \
    --request "PATCH" \
    --data '{"id": 1, "name": "Pasta","description": "Nice11","price": 49, "restaurantID": "e814691f-b53e-45c4-8253-e2f2a7f5ff35"}'
```

## Delete Dish
```
curl http://localhost:8080/api/v1/dishes/d2b72550-bf06-4ff9-ace3-09a1b83a8218 \
    --include \
    --header "Content-Type: application/json" \
    --header "RestaurantID: e814691f-b53e-45c4-8253-e2f2a7f5ff35" \
    --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg3MTM1MTIsInN1YiI6IjRhMGVkNTQ1LTY2M2QtNGI2OS1iNzA5LTg4MmIxMGI1ZGZjMiJ9.XVdCBb5aONRuzyRhnuR-EfT9Brx7-sn9binGsq_fE7s" \
    --request "DELETE"
```

## Create Rating
```
curl http://localhost:8080/api/v1/ratings \
    --include \
    --header "Content-Type: application/json" \
    --header "RestaurantID: e814691f-b53e-45c4-8253-e2f2a7f5ff35" \
    --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg3NDIyODcsInN1YiI6IjhhYmMyMGI4LTYzYzEtNGJhZi04MGIzLTM1MmM4ODAzYmNiNSJ9.oqQq32ehtDAPkPgwdv0JyLvuSpu5UB5lYJUOsaDlRZE" \
    --request "POST" \
    --data '{"description": "Food was wonderful!", "dishID": "44185708-b33e-4781-918b-baba7878348f"}'
```

## Register User
```
curl http://localhost:8080/api/v1/users/register \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "Smaug","emailAddress": "smaug@erabor.com","password": "superpassword"}'
```

## Login User
```
curl http://localhost:8080/api/v1/users/login \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"emailAddress": "gimli@erabor.com","password": "superpassword"}'

curl http://localhost:8080/api/v1/users/login \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"emailAddress": "keely@erabor.com","password": "password"}'

curl http://localhost:8080/api/v1/users/login \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"emailAddress": "smaug@erabor.com","password": "password"}'
```


## List Restaurants
```
curl http://localhost:8080/api/v1/restaurants
```

### Gemini PRO
```
curl 'https://generativelanguage.googleapis.com/v1/models/gemini-pro:generateContent?key=AIzaSyABKNVcyrO5EqwckFP7TZB4OcNNaTmykek' \
    --include \
    --header 'Content-Type: application/json' \
    --request "POST" \
    --data '{ "contents" : [ { "parts": [ {"text": "Write a story about a magic backpack"} ] } ] }'
```