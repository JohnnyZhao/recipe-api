# Recipe API Service

The Recipe API Service is an HTTP API service that provides a set of endpoints to perform CRUD operations on recipe records. The service is built using Go programming language and uses the Gin web framework and GORM ORM library.

## Start Service for Testing

To start testing the service, make sure you have Docker and docker-compose installed. Then, start the MySQL container and API service by running the following command:
```bash
docker-compose up -d --build
```

Once the services are up and running, the API service will be available at http://0.0.0.0:8080. Additionally, the MySQL instance will be accessible at the host 0.0.0.0 with port 3306, database name is recipe, and the root user's password will be example


## Usage

The Recipe API Service provides the following endpoints:

### POST /recipes
Create a new recipe record.

Request(All fields are required):
```json
{
    "title": "Spaghetti Bolognese",
    "making_time": "45 mins",
    "serves": "2 people",
    "ingredients": "spaghetti, ground beef, tomato sauce",
    "cost": 10
}
```

Response:
```json
{
    "message": "Recipe successfully created!",
    "recipe": {
        "id": 1,
        "title": "Spaghetti Bolognese",
        "making_time": "45 mins",
        "serves": "2 people",
        "ingredients": "spaghetti, ground beef, tomato sauce",
        "cost": 10,
        "created_at": "2022-04-05T12:00:00Z",
        "updated_at": "2022-04-05T12:00:00Z"
    }
}
```
Error Response:

If an error occurs while create a recipe, the API returns an HTTP error response with an error message in the response body.

The error response has the following format:
```json
{
    "message": "Error message",
    "required": "Field1, Field2, ..."
}
```

The `message` field contains the error message, while the `required` field contains a comma-separated list of required fields that were missing from the request payload.

### GET /recipes
Get a list of all recipe records.

Response:
```json
{
    "recipes": [
        {
            "id": 1,
            "title": "Spaghetti Bolognese",
            "making_time": "45 mins",
            "serves": "2 people",
            "ingredients": "spaghetti, ground beef, tomato sauce",
            "cost": 10,
            "created_at": "2022-04-05T12:00:00Z",
            "updated_at": "2022-04-05T12:00:00Z"
        },
        {
            "id": 2,
            "title": "Fish and Chips",
            "making_time": "30 mins",
            "serves": "2 people",
            "ingredients": "fish, potatoes",
            "cost": 8,
            "created_at": "2022-04-05T12:00:00Z",
            "updated_at": "2022-04-05T12:00:00Z"
        }
    ]
}

```

### GET /recipes/:id
Get a specific recipe record by ID.

Response:
```json
{
  "message": "Recipe details by id",
  "recipes": [
    {
      "id": 1,
      "title": "Spaghetti Bolognese",
      "making_time": "45 mins",
      "serves": "2 people",
      "ingredients": "spaghetti, ground beef, tomato sauce",
      "cost": 10,
      "created_at": "2022-04-05T12:00:00Z",
      "updated_at": "2022-04-05T12:00:00Z"
    }
  ]
}
```

### PATCH /recipes/:id

Update a specific recipe record by ID.

Request(provide at least one field below):
```json
{
  "title": "Spaghetti Bolognese",
  "making_time": "45 mins",
  "serves": "2 people",
  "ingredients": "spaghetti, ground beef, tomato sauce",
  "cost": 10
}
```

Response:
```json
{
    "message": "Recipe successfully updated!",
    "recipe": {
        "id": 1,
        "title": "Spaghetti Carbonara",
        "making_time": "45 mins",
        "serves": "2 people",
        "ingredients": "spaghetti, ground beef, tomato sauce",
        "cost": 12,
        "created_at": "2022-04-05T12:00:00Z",
        "updated_at": "2022-04-05T13:00:00Z"
    }
}
```

### DELETE /recipes/:id

Delete a specific recipe record by ID.

Response:
```json
{
    "message": "Recipe successfully removed!"
}
```


### General Error Response

If an error occurs while processing a request, the Recipe API Service returns an HTTP error response with an error message in the response body.

The error response has the following formats, you could find the error message useful:

1. Internal Server Error:

```json
{
"message": "Internal Server Error"
}
```

2. Bad Status Error:

```json
{
"message": "invalid payload"
}
```

3. Recipe Not Found Error:

```json
{
  "message": "No recipe found"
}
```