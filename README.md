# Receipt Processor

This is a Go-based webservice implementing the Receipt Processor API as specified. It processes receipts via a POST endpoint, assigns a unique ID, and calculates points for a receipt via a GET endpoint based on predefined rules. Data is stored in memory, as persistence is not required.

## Features

- **API Endpoints**:
  - `POST /receipts/process`: Accepts a JSON receipt and returns a unique ID.
  - `GET /receipts/{id}/points`: Returns the points awarded for the specified receipt ID.
- **Validation**: Strict input validation per the API schema.
- **Concurrency**: Thread-safe in-memory storage using a mutex.
- **Testing**: Unit tests for points calculation, covering provided examples.
- **Lightweight**: Minimal dependencies (only `github.com/google/uuid`).

## Prerequisites

- **Go** (version 1.18 or later recommended)
- **Optional**: Docker, if you prefer to run the service in a container

## Running the Application

### Option 1: Using Go Directly

1. Clone the repository:

   ```sh
   git clone https://github.com/lokeshguduru/receipt-processor.git
   cd receipt-processor
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Run the server:

   ```sh
   go run .
   ```

## Usage

### Process a Receipt

**PowerShell (or Terminal):**

```sh
curl -X POST -H "Content-Type: application/json" -d ^
"{\"retailer\": \"Target\", \"purchaseDate\": \"2022-01-01\", \"purchaseTime\": \"13:01\", \"items\": [{\"shortDescription\": \"Item\", \"price\": \"1.00\"}], \"total\": \"1.00\"}" ^
http://localhost:8080/receipts/process
```

**Response:**

```json
{"id": "some-uuid"}
```

### Get Points

Replace `<id>` with the ID from the POST response:

```sh
curl http://localhost:8080/receipts/<id>/points
```

**Response:**

```json
{"points": 87}
```

### Using a JSON File

1. Save as `receipt.json`:

   ```json
   {
     "retailer": "Target",
     "purchaseDate": "2022-01-01",
     "purchaseTime": "13:01",
     "items": [
       {
         "shortDescription": "Item",
         "price": "1.00"
       }
     ],
     "total": "1.00"
   }
   ```

2. Send using curl:

   ```sh
   curl -X POST -H "Content-Type: application/json" -d @receipt.json http://localhost:8080/receipts/process
   ```

### Testing

Run unit tests to verify points calculation:

   ```sh
   go test -v
   ```


## Notes

- Data is not persisted (in-memory only).
- Errors:
  - `400` for invalid receipts
  - `404` for unknown IDs
- The service uses the standard library plus github.com/google/uuid for ID generation.
