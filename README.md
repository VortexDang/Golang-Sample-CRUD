
# Udacity Golang Final Project

This is the project - which can be also called a CRUD helps you to interact with a sample mock database.


## Requirements

```bash
    Go >= 1.19.3
    Go Extention
```
## Installation

Install Golang

```bash
    https://go.dev/doc/install
```
    
## Deployment

To deploy this project run

```bash
  go run main.go
```


## API Reference
#### Get all customers

```http
  GET /customers
```

#### Get customer

```http
  GET /customers/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of customer to fetch |

#### Get all items

```http
  GET /customers
```

#### Add new customer to list

```http
  POST /customers
```

#### Update customer to list

```http
  PUT /customers/{id}
```

#### Delete customer

```http
  DELETE /customers/{id}
```


## Authors

- [@VortexDang](https://github.com/VortexDang)


## Support

For support, email ben.tran@jungtalents.com or join our Slack.

