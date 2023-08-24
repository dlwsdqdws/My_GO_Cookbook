# Database

- [Database](#database)
  - [Data Storage](#data-storage)
    - [Storage System](#storage-system)
    - [Database Classification](#database-classification)
      - [Relational Database](#relational-database)
      - [Non-relational Database](#non-relational-database)
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
- RAID0+1: combine RAID0 and RAID1, with strong fault tolerance and good write bandwidth.

### Database Classification

#### Relational Database

#### Non-relational Database

## RDBMS

## Redis
