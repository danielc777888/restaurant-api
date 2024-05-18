# CURL

## Create Dish
```
curl http://localhost:8080/api/v1/dishes \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "Pasta","description": "Nice","price": 49, "restaurantID": 1}'
```

## Get Dish
```
curl http://localhost:8080/api/v1/dishes/1
```

## List Dish
```
curl http://localhost:8080/api/v1/dishes
```

## Update Dish
```
curl http://localhost:8080/api/v1/dishes \
    --include \
    --header "Content-Type: application/json" \
    --request "PATCH" \
    --data '{"id": 1, "name": "Pasta","description": "Nice11","price": 49, "restaurantID": 1}'
```

## Delete Dish
```
curl http://localhost:8080/api/v1/dishes/4 \
    --include \
    --header "Content-Type: application/json" \
    --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg2MDE2MjIsInN1YiI6Mn0.TmNqSNVjodz8IJ_KudXzKqnL37JW7WEIIfFEFGEGawA" \
    --request "DELETE"
```

## Create Rating
```
curl http://localhost:8080/api/v1/ratings \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"description": "Great service", "dishID": 1}'
```

## Register User
```
curl http://localhost:8080/api/v1/users/register \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "Gimli","emailAddress": "gimli@erabo.com","password": "superpassword"}'
```

## Login User
```
curl http://localhost:8080/api/v1/users/login \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"emailAddress": "gimli@erabo.com","password": "superpassword"}'
```


## List Restaurants
```
curl http://localhost:8080/api/v1/restaurants
```