# Etcd Study Guide


# CLI Commands

## How do you get cluster info?

<details><summary>show</summary>
<p>

```
    etcdctl ls  /_etc/machines --recursive
```

### List all known hosts

```
    etcdctl get /_etc/machines/<token>        
```

### Details of a host

```
    etcdctl get /_etc/config
```

</p>
</details>

## How do you access the key space?

<details><summary>show</summary>
<p>

### Get top-level keys

```
etcdctl ls

```

### Get full tree

```
etcdctl ls / --recursive

```

### Get key details

```

etcdctl get <key path>
```

### Get key value only

```
etcdctl get <key path> --rev=<number> 

```

### Get key older revisions 

```

etcdctl -o extended get <key path> 
```

### Get key and etadata

```
etcdctl get <key path> -w=json 

```

### Output in json w/ metadata

```
etcdctl get <key path> -w=json 

```

</p>
</details>
    
    
## How do you execute batch queries?

### Get all keys key1, key2, key3, ..., key10

```
    etcdctl get key1 key10 

```

### Get all keys matching ^key

```
    etcdctl get --prefix key                    

```
### Get max 10 keys matching ^key

```
    etcdctl get --prefix key --limit=10         

```

### Creating a path

* Removes only empty paths

```
    etcdctl mkdir /newpath
    etcdctl rmdir /newpath     

```

</p>
</details>


## How can you manipulate keys?

<details><summary>show</summary>
<p>


### Create key

```
    etcdctl mk     /path/newkey some-data

```

### Create or update key

```
    etcdctl set    /path/newkey some-data 

```

### Update key

```
    etcdctl update /path/key new-data           
    etcdctl put    /path/key new-data
    etcdctl rm     /path/key
    etcdctl rm     /path --recursive

```

</p>
</details>

## How can you make data and paths expire?

<details><summary>show</summary>
<p>

* by passing --ttl when creating paths

### Path with expiration

```
    etcdctl mkdir     /path --ttl 120     

```

### Reset path expiration

```
    etcdctl updatedir /path --ttl 120     

```

</p>
</details>
    
## Monitoring paths

<details><summary>show</summary>
<p>

```

    etcdctl watch /path
    etcdctl watch --recursive /path

```
    
### Trigger command on event

```
    etcdctl watch --recursive /path -- printf "Path /path was changed.\n"

```

</p>
</details>


## How can you compact revisions

<details><summary>show</summary>
<p>

### Drop all revisions older than <number>

```
    etcdctl compact <number>     

```

</p>
</details>

    
## How can you use curl to get keys?

<details><summary>show</summary>
<p>


### Sample curl

```

    curl -L http://127.0.0.1:4001/v2/keys/

```

</p>
</details>

## What are some examples of endpoints?

<details><summary>show</summary>
<p>



```
    /version
    /v2/stats/self         # Node info
    /v2/stats/store        # Statisitics ops/s
    /v2/stats/leader       # Cluster master/slave details
    
    /v2/keys
    /v2/keys/?recursive=true
```

</p>
</details>

## What port is the admin API on?

<details><summary>show</summary>
<p>

Separately from the port 4001 cluster API there is also an admin API for configuration changes on default port 7001

```
    /v2/admin/config       # GET returns settings, XPUT changes settings
    /v2/admin/machines     # Cluster details

```

</p>
</details>