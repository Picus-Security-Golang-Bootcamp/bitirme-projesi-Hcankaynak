# bitirme-projesi

# Best Online Shopping Site (BOSS)

## Project Structure

- docs: store swagger yml file.
- internal: except api directory, all main functionalities are here. Including entities, handlers.
- pkg: base operations. read config, connect db, jwt implementation, middleware.

## EndPoints

#### User Endpoints

### User

```http
POST /users/signUp
POST /users/login
```

### Category

- First one listing all categories
- Second one add new category (singular)

```http
POST /category/list
POST /category
```

### Product

- First creating a new product

```http
POST /products
```

### Basket

- First getting basket for user
- Second one adding item to basket

```http
GET /basket
POST /basket
```