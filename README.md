To add IS to the list of IS, send POST request within the route `localhost:9000` with the following content:

```
{"name": "inno", "owner": "fio", "admin": "admin", "contacts": "contact"}
```

It will assign an ID to the entry and corresponding fields that you specify. 

To add info about certain IS, navigate to the desired IS id (e.g. `localhost:9000/1/info`) and send POST request with the following content

```
{"name": "name", "owner": "owner", "vms": "vms", "cpu": 2, "ram": 128, "hdd": "128mb", "softwareused": "softwareused", "adminname": "adminname", "adminemail": "adminemail", "admintg": "admintg", "resourceassignment": "resourceassignment", "status": true}
```

To update the info about informational system in the list with all systems navigate to the desired IS id (e.g. `localhost:9000/1`) and send PUT request with the feild you want to update:

```
{"name": "infsys"}
```

To list all ISs send GET request to the route `localhost:9000`

To get full info about IS send GET request to the route with desired IS (e.g. `localhost:9000/1/info`)

To delete IS send DELETE request to the desired route (e.g. `localhost:9000/1`). This will delete corresponding entries from `listinfsys` and `infsys` tables.
