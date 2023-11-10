# Cat Breed

The application fetches cat breed information from the Cat Fact API, processes and organizes the data, and writes the result to a JSON file (`data/out.json`).

## Run the application

```bash 
    go run cmd/app/main.go
```

## Structure

- `main.go`: The main application file.
- `internal/model/model.go`: Contains the data models for cat breed and response.
- `data/out.json`: The output file where the processed data is written.