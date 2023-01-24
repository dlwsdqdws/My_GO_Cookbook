# GORM

- [GORM](#gorm)
  - [Installation](#installation)
  - [Connect to Database](#connect-to-database)
  - [CRUD](#crud)
    - [Create](#create)
    - [Read](#read)
    - [Update](#update)
    - [Delete](#delete)
  - [Transaction](#transaction)
  - [Hook](#hook)
  - [Plugins](#plugins)

## Installation

```go
go get -u gorm.io/gorm
// take mysql as an example
go get -u gorm.io/driver/mysql
```

## Connect to Database

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

## CRUD

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

### Create

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

### Read

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
db.Model(&product).Where("name = ?", "jinzhu").Update("Price", 200)
```

- Multiple

```go
db.Model(&Product{ID : 111}).Updates(Product{Name : "hello", Age : 20})
```

- Can use `map` or `Select` update zero values.

```go
db.Model(&product).Updates(map[string]interface{}{"Price": 200, "activated": false})
```

- Selected Column

```go
// only update price even though multiple columns in the map
db.Model(&Product{ID : 111}).Select("Price").Updates(map[string]interface{}{"Price": 200, "activated": false})
```

- SQL

```go
db.Model(&Product{ID : 111}).Updates("age", gorm.Expr("age * ? + ?", 2, 100))
```

### Delete

- Hard Delete

```go
db.Delete(&p)
// DELETE from products where id = 10;

// where could be used
db.Where("name = ?", "jinzhu").Delete(&p)
// DELETE from products where id = 10 AND name = "jinzhu";

db.Delete(&User{}, 10)
// DELETE FROM products WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM products WHERE id = 10;

db.Delete(&products, []int{1,2,3})
// DELETE FROM products WHERE id IN (1,2,3);

db.Where("product LIKE ?", "%jinzhu%").Delete(&Product{})
// DELETE from products where product LIKE "%jinzhu%";

db.Delete(&Product{}, "product LIKE ?", "%jinzhu%")
// DELETE from products where product LIKE "%jinzhu%";
```

- Soft Delete

```go
// use gorm.DeletedAt
// when call Delete(), the data will not be deleted physically
// but label DeletedAt as current time
// when calling Find(), soft-deleted data will be ignored
type User struct {
  ID int64
  Name string
  Age int64
  Deleted gorm.DeletedAt
}

db.Where("age = ?", 20).Delete(&User{})
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// Use Unscoped can find soft-deleted data or realize hard delete
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
```

## Transaction

- Transaction is a sequence of database operations that are either all executed or none of them are executed. A transaction consists of all database operations performed between the start of the transaction and the end of the transaction.
- If CUD is not needed, better to disable transaction to improve performance. Use `PrepareStmt` caches prepared statements can improve the speed of subsequent calls.

```go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
  PrepareStmt: true
})
```

- Gorm Transaction provides `Begin()`, `Commit` and `Rollback()`.

```go
// use tx instead of db to start a transaction
tx := db.Begin()
​
// db operations
tx.Create(...)
​
// ...
​
// rollback if error occurs
if err := tx.Create(...).Error; err != nil{
  tx.Rollback()
  return
}
​
// otherwise submit this transaction
tx.Commit()
```

- Gorm also provides a Transaction method.

```go
if err := db.Transaction(func(tx *gorm.DB) error {
  if err := tx.Create(&User{Name: "name"}).Error; err != nil {
    // rollback automatically
    return err
  }

  if err := tx.Create(&User{Name: "name1"}).Error; err != nil {
    return err
  }

  return nil
}); err != nil{
  return
}
```

## Hook

- Hook are functions called before or after CRUD operations.

```go
type User struct {
    ID      int64
    Name    string `gorm:"default:galeone"`
    Age     string `gorm:"default:18"`
}
​
type Email struct {
    ID      int64
    Name    string
    Email   string
}
​
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    if u.Age < 0 {
        return errors.New("can't save invalid data")
    }
    return nil;
}
​
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
    return tx.Create(&Email{ID: u.ID, Email: i.Name + "@***.com"}).Error
}
```

- Hook will be called automatically when calling CRUD. If it returns an error, GORM will stop subsequent operations and rollback the transaction.

## Plugins

|            plugins            |                 links                 |
| :---------------------------: | :-----------------------------------: |
|        Code Generation        |      www.github.com/go-gorm/gen       |
| Optimizer/Index/Comment Hints |     www.github.com/go-gorm/hints      |
|        Sharding Tables        |    www.github.com/go-gorm/sharding    |
|        Optimistic Lock        | www.github.com/go-gorm/optimisticlock |
|     Read/Write Splitting      |   www.github.com/go-gorm/dbresolver   |
|         OpenTelemetry         | www.github.com/go-gorm/opentelemetry  |