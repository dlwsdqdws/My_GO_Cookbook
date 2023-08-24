# Database

- [Database](#database)
  - [Data Storage](#data-storage)
    - [Storage System](#storage-system)
    - [Database Classification](#database-classification)
      - [Relational Database](#relational-database)
      - [Non-relational Database](#non-relational-database)
    - [Database VS Traditional Storage System](#database-vs-traditional-storage-system)
  - [Products](#products)
    - [Storage](#storage)
      - [Standalone Storage](#standalone-storage)
    - [Database](#database-1)
  - [RDBMS](#rdbms)
  - [Redis](#redis)

## Data Storage

### Storage System

1. The storage system is a software that provides Read/Write & control interfaces and can safely and effectively persist data.

<p align="center"><img src="../static/img/database/storage/storage.png" alt="RPC Process" width="500"/></p>

- To meet the Read/Write limitations of the hardware, cache is essential.
- Copy is CPU intensive and should be used as little as possible.

2. Redundant Array of Independent Disks (RAID)

- RAID0: Simple combination of multiple disks improves the bandwidth, but there is no fault-tolerant design.
- RAID1: One disk corresponds to one additional mirror disk, the real space utilization rate is 50%, but the fault tolerance is strong.
- RAID0+1: Combine RAID0 and RAID1, with strong fault tolerance and good write bandwidth.

### Database Classification

#### Relational Database

1. Data Structure: Relational databases use tables to store data, where tables consist of rows and columns, with each column representing an attribute. Data is stored in a structured manner, with data types and relationships defined within pre-defined schemas.
2. Data Consistency: Relational databases emphasize data consistency and integrity. Data follows strict transactions to ensure integrity during data modifications.
3. Query Language: Relational databases utilize SQL (Structured Query Language) for data querying and manipulation. SQL is a powerful standardized query language used for various operations, including data retrieval, insertion, updating, and deletion.
4. Use Cases: They are suitable for scenarios that demand data integrity and consistency, such as finance, e-commerce, management systems, etc.

#### Non-relational Database

1. Data Structure: Non-relational databases adopt more flexible data models, such as key-value pairs, documents, column families, graphs, etc. Data can be semi-structured or unstructured.
2. Data Consistency: Non-relational databases often relax the requirement for strict data consistency, emphasizing availability and partition tolerance. Some provide weak or eventual consistency guarantees.
3. Query Language: Non-relational databases have varying query languages based on the database type, lacking a unified standard. Some offer SQL-like querying languages, while others use APIs.
4. Use Cases: They are suitable for big data, high concurrency, distributed environments, and scenarios requiring flexible data models and scalability, such as social networks, real-time analytics, log processing, etc.

### Database VS Traditional Storage System

Database is a kind of storage system but it has many advantages over traditional storage system.

1. Structured data management.

<p align="center"><img src="../static/img/database/storage/structured.png" alt="RPC Process" width="500"/></p>

2. Transaction Capabilities: ACID.

- Atomicity: Atomicity refers to a transaction being the smallest indivisible unit of work. All operations within a transaction are either executed entirely or not at all, avoiding situations where some operations succeed while others fail. If an error occurs during transaction execution, the system rolls back the transaction, restoring the database to its state before the transaction started.
- Consistency: Consistency ensures that the state of the database remains consistent before and after a transaction. During transaction execution, the integrity constraints of the database cannot be violated, meaning that data must satisfy predefined rules such as uniqueness and foreign key constraints.
- Isolation: Isolation means that when multiple transactions are running concurrently, each transaction's operations are isolated from those of other transactions. This ensures that the modifications made by one transaction are not visible to other transactions until the first transaction is committed. Isolation prevents data interference and conflicts between different transactions.
- Durability: Durability ensures that once a transaction is committed, the modifications made by it are permanently saved in the database, even in the event of system crashes or failures. The database system persists the modifications of committed transactions to storage media, ensuring the persistence of data.

3. Complex Query Capability: Domain-Specific Language (DSL)

<p align="center"><img src="../static/img/database/storage/query.png" alt="RPC Process" width="500"/></p>

## Products

### Storage

#### Standalone Storage

1. File System 

- Index Node: a data structure used to store metadata about a file or directory. Each file or directory in the file system has a corresponding inode entry, which contains various information about the file but does not include the actual content of the file. 
- Directory Entry: a data structure refers to an entry within a directory that associates a file or directory name with its corresponding inode number. It serves as a mapping between the human-readable name of a file or directory and the internal data structure (inode) that contains metadata about that file or directory.

2. Key-Value

- Method: put(key, value) & get(key)
- Data Structure: LSM-Tree (Sacrifice read performance in pursuit of write performance)
  
<p align="center"><img src="../static/img/database/products/lsmtree.png" alt="RPC Process" width="500"/></p>

### Database

## RDBMS

## Redis
