## Indigo
A HTTP-based JSON-oriented hierarchical key/value storage

This code was created as a support for embedded systems self-configuration modules where we could, in real-time, adjust params, read later by any other system and, after fine-tunning it, store it on a json file, for later usage.

On my original scenario, I had multiple NodeMCU and arduino+ethernet shield-based systems, whose configuration changes all the time. On the same scenario, I had an UAV sending its telemetry and constantly requests its PID's coefficients and directions to follow. In both cases, various different embedded systems are requesting and setting its own configuration values via a HTTP endpoint and, periodically, a STORE verb is issued, to save configuration data into disk.

### HTTP-based data management interface

On [web.go](web.go), we describe and define the model for data insertion and retrieval using HTTP PATH and VERBs as control structure. The advantage of using such structures is, specially, extensibility. It's easy to append new functions and paths for the requests, what makes the applications modular and robust.

The base endpoints are:

|  VERB  |                  URL                  |                                                        INFO                                                        |
|:-------|:--------------------------------------|:-------------------------------------------------------------------------------------------------------------------|
| GET    | http://url:port/[table]/[key]         | To extract information stored at table.key. It returns null if there is no data;                                   |
| POST   | http://url:port/[table]/[key]/[value] | To insert an informarion and store it at table.key. It may be user to store a new value ou update an existing one; |
| DELETE | http://url:port/[table]/[key]         | To remove the stored data, if there is any;                                                                        |
| STORE  | http://url:port/                      | To store current in-memmory data into database file;                                                               |
| LOAD   | http://url:port/                      | To reload data from database file;                                                                                 |
| DUMP   | http://url:port/                      | To dump everything from in-memmory data into a json.                                                               |

To extend, is just a matter of add a new HTTP VERB and implement its response. 

### JSON-based hierarchical key/value data engine

On [data.go](data.go) we describe and define the data engine enveloped with the **DataBase** struct, that, once initialized, may be used to perform all operations over all data. These operations are:
- **DataBase.Init(`database name` as string)** To **start** a database, named as `database name` using `databasename.JSONdb` as  file for persistation in JSON format;
- **DataBase.Set(`table name` as string, `field` as string, `value` as string)** To **set** a data value, hierarchically, given a `table` and a `field`; 
- **DataBase.Get(`table name` as string, `field` as string)** To **get** a data value, hierarchically, given a `table` and a `field`;
- **DataBase.Del(`table name` as string, `field` as string)** To **remove** a data value, hierarchically, given a `table` and a `field`; 
- **DataBase.Dump()** To **dump** the whole database;
- **DataBase.Load()** To **load** the data from database file defined on **Init**;
- **DataBase.Store()** To **store** the data into the database file defined on **Init**;

Usage examples for [data.go](data.go) itself:
```
 d := new(DataBase)
 d.Init("test")
 d.Set("table1", "field1", "value1")
 d.Set("table2", "field2", "value2")
 d.Set("table1", "field3", "value3")
 fmt.Println(d.Get("table2", "field2"))
 d.Del("table1", "field1")
 fmt.Println(d.Dump())
 d.Store()
```

### Some usage examples

- For example, a temperature was sending:
   ```
     curl -XPOST http://10.0.1.6:8080/sensors/temp/31.2
   ```
- other sensor, a humidity one, was sending:
   ```
     curl -XPOST http://10.0.1.6:8080/sensors/humidity/63.8
   ```
- while other, a battery probe, was sending:
   ```
     curl -XPOST http://10.0.1.6:8080/sensors/batt/5.1
   ```
- and so on. Later, any system could store it into disk by sending a:
   ```
     curl -XSTORE http://10.0.1.6:8080/
   ```
- or, get the whole configuration by sending:
   ```
     curl -XDUMP http://10.0.1.6:8080/
   ```
   > which results in:
   ```json
   {
    "databasename": "indigo",
    "tables": {
     "sensors": {
      "batt": "5.1",
      "humidity": "63.8",
      "temp": "31.2"
     }
    }
   }
   ```

### TODO
- export valid SQL, not just JSON.
- test for more errors, not just ignore them

### Outro
