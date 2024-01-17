## Setup and Run

### Prerequisites
Make sure to have Protoc installed to generate proto files before initializing the project. See detail https://grpc.io/docs/protoc-installation/

### 1. Init Project

Before running the API and services, make sure to set up the project by performing the following steps:

```bash
# Setup Go modules tidy dependencies and generate proto file
make init
```

### 2. Run API Gateway

To start the API Gateway, use the following command:

```bash
make api
```

### 3. Run gRPC Article Service

To launch the gRPC Article Service, execute the following command:

```bash
make article-service
```

Now, the API Gateway and the gRPC Article Service should be up and running.

## Specification API

## 1. Create a New Article

### Endpoint
`POST /article/`

### Request
```json
{
  "title": "Your Title Here (minimum 20 characters)",
  "content": "Your Content Here (minimum 200 characters)",
  "category": "Your Category (minimum 3 characters)",
  "status": "publish/draft/trash"
}
```

### Response
```json
{}
```

### Description
Creates a new article with the provided title, content, category, and status.

---

## 2. Retrieve Articles with Pagination

### Endpoint
`GET /article/<limit>/<offset>`

### Request
- `limit`: Number of articles to retrieve (integer)
- `offset`: Starting point for retrieving articles (integer)

### Response
```json
{
  "data": [
    {
      "id": ""
      "title": "",
      "content": "",
      "category": "",
      "status": ""
    }
  ],
  "total": 0
}
```

### Description
Retrieves articles from the database with pagination. The `limit` parameter determines the number of articles, and the `offset` parameter defines the starting point.

---

## 3. Retrieve a Specific Article

### Endpoint
`GET /article/<id>`

### Response
```json
{
  "id": "",
  "title": "",
  "content": "",
  "category": "",
  "status": ""
}
```

### Description
Retrieves the article with the specified ID.

---

## 4. Update an Article

### Endpoint
`PUT /article/<id>`

### Request
```json
{
  "title": "Updated Title (minimum 20 characters)",
  "content": "Updated Content (minimum 200 characters)",
  "category": "Updated Category (minimum 3 characters)",
  "status": "publish/draft/trash"
}
```

### Response
```json
{}
```

### Description
Updates the data of the article with the specified ID.

---

## 5. Delete an Article

### Endpoint
`POST/DELETE /article/<id>`


### Description
Deletes the article with the specified ID.

### Article Data Specifications
- **Title**: Required, minimum 20 characters
- **Content**: Required, minimum 200 characters
- **Category**: Required, minimum 3 characters
- **Status**: Required, choose between "publish", "draft", or "trash"
