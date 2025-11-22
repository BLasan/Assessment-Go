# Assessment-Go

A Go web service that provides a hello-world endpoint with validation.

## Endpoint

**GET /hello-world**

Query Parameters:
- `name` (required): A name starting only with English letters (A-Z, a-z)

## Response Rules

1. If the `name` parameter is empty or starts with invalid characters (non-English letters), the service returns:
   - HTTP Status: `400 Bad Request`
   - Body: `{"error": "Invalid Input"}`

2. If the `name` parameter starts with letters A-M or a-m (case-insensitive), the service returns:
   - HTTP Status: `200 OK`
   - Body: `{"message": "Hello {name}"}`

3. If the `name` parameter starts with letters N-Z n-z (case-insensitive), the service returns:
   - HTTP Status: `400 Bad Request`
   - Body: `{"error": "Invalid Input"}`

4. If the `name` parameter has leading/trailing spaces, then those will be trimmed and process. If the response will be changed according to the validations given.

## Assumptions

1. There can be names containing numbers, unicode characters in the middle or at the end, but not as the starting character.

## Running the Service

### Build
```bash
go build -o hello-world-server
```

### Run
```bash
./hello-world-server
```

The server will start on port 8080.

### Test
```bash
go test -v
```

## Example Requests

```bash
# Valid request (returns success)
curl "http://localhost:8080/hello-world?name=Alice"
# Response: {"message":"Hello Alice"}

# Invalid request (starts with N-Z)
curl "http://localhost:8080/hello-world?name=Nancy"
# Response: {"error":"Invalid Input"}

# Invalid request (contains numbers)
curl "http://localhost:8080/hello-world?name=John123"
# Response: {"error":"Invalid Input"}

# Invalid request (empty name)
curl "http://localhost:8080/hello-world?name="
# Response: {"error":"Invalid Input"}
```