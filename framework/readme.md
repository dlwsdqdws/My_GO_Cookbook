# Framework

- [Framework](#framework)
  - [ORM - Gorm](#orm---gorm)
    - [Installation](#installation)
    - [Connect to Database](#connect-to-database)
    - [CRUD](#crud)
      - [Create](#create)
      - [Read](#read)
    - [Update](#update)
    - [delete](#delete)
  - [RPC - Kitex](#rpc---kitex)
  - [HTTP - Hertz](#http---hertz)


## ORM - Gorm

### Installation

```go
go get -u gorm.io/gorm
// take mysql as an example
go get -u gorm.io/driver/mysql
```

### Connect to Database

- Gorm can support MySQL, PostgreSQL, SQlite, SQL Server. Take SQLServer as an example.

```go
// import driver
import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
```

- DSN(data source name).

```go
"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```



- Note that always check err before operating CRUD. Here `panic` is suggested if the database cannot be connected.

### CRUD

- GORM uses `ID` as the primary key, the snake-case of the structure name as the table name, the snake-case field name as the column name, and uses the `CreatedAt`, `UpdatedAt` fields to track creation and update time. So, `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` will be automatically created and work as their name.

```go
type Product struct {
    ID      uint    `gorm:"primarykey"`
    Code    string  `gorm:"column: code"`
    Price   uint    `gorm:"column: user_id"`

    // can set default values
    Name    string `gorm:"default:galeone"`
    Age     int64  `gorm:"default:18"`

    // the belows can be created automatically
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt sql.NullTime `gorm:"index"`
}
```

#### Create

- One piece of data.

```go
p := &Product({ Code : "042", Price: 100})
res := db.Create(p)
if res.Error != nil{
    // error handler
}
```

- Multiple pieces of data

```go
// Create a list of struct
products := []*Product{{Code : "041"}, {Code : "042"}, {Code : "043"}}
res := db.Create(products)
if res.Error != nil{
    // error handler
}
```

- Is it no need to set values for `ID`, `CreatedAt`, etc.
- Use `clause.Onconfict` to handle conflict. We cannot use `when` after `Create()`.

```go
p := &Product({ Code : "042", ID: 1})
// here we do nothing when conflict happens
db.Clauses(clause.Onconfict{DoNothing : true}).Create(&p)
```

#### Read

- `First` method returns the first data that meets the specified criteria, `ErrRecodeNotFound` if no such data.

```go
u := &Prodyct{}
db.First(u)
```

- `Find` method returns multiple data meets `where` criteria, nothing if no such data.

```go
p := &Product{}
// Get first matched record
res := db.Where("name = ?", "jinzhu").First(p)
// SELECT * FROM products WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

products := make([]*Product, 0)
// Get all matched records
res = db.Where("name <> ?", "jinzhu").Find(&products)
// SELECT * FROM products WHERE name <> 'jinzhu';

// Other Inquiries 
// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&products)
// SELECT * FROM products WHERE name IN ('jinzhu','jinzhu 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&products)
// SELECT * FROM products WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&products)
// SELECT * FROM products WHERE name = 'jinzhu' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&products)
// SELECT * FROM products WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&products)
// SELECT * FROM products WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
```

- When using a struct as an inquiry, the zero values(eg, 0, false) will not be used. If zero values is needed, we can use `map` as a inquery.

```go
db.Where(map[string]interface{}{"name": "jinzhu", "age": 0}).Find(&products)
// SELECT * FROM products WHERE name = "jinzhu" AND age = 0;
```

- Slice can used as a inquery.

```go
db.Where([]int64{20, 21, 22}).Find(&products)
// SELECT * FROM products WHERE id IN (20, 21, 22);
```

### Update

- Single

```go
db.Model(&product).Update("Price", 200)
```

### delete

## RPC - Kitex

## HTTP - Hertz