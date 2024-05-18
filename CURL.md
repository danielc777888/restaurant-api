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
curl http://localhost:8080/api/v1/dishes/d6262803-770a-4dd1-b680-e0a932d6e636 \
    --include \
    --header "Content-Type: application/json" \
    --header "RestaurantID: e814691f-b53e-45c4-8253-e2f2a7f5ff35" \
    --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg2NDc2NDIsInN1YiI6IjIyM2FhY2Y3LTM1ZjgtNDFkZS05ZDM2LTFiNjAxNmU5NDUzNCJ9.yGZkvaSe6StgGNTy6fS6XtCHMPBM-f5DCUBscpsrvCU" \
    --request "DELETE"
```

## Create Rating
```
curl http://localhost:8080/api/v1/ratings \
    --include \
    --header "Content-Type: application/json" \
    --header "RestaurantID: e814691f-b53e-45c4-8253-e2f2a7f5ff35" \
    --request "POST" \
    --data '{"description": "Great service", "dishID": "d6262803-770a-4dd1-b680-e0a932d6e636"}'
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
    --data '{"emailAddress": "smaug@erabor.com","password": "superpassword"}'
```


## List Restaurants
```
curl http://localhost:8080/api/v1/restaurants
```