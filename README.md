To add IS to the list of IS, send POST request with the following content:

```
{"name": "inno", "owner": "fio", "admin": "admin", "contacts": "contact", "infsys_id": 1}
```

To add info about IS, navigate to the desired IS id (eg. `localhost:9000/1/info`) and send POST request with the following content

```
{"id": 1, "name": "name", "owner": "owner", "vms": "vms", "cpu": 2, "ram": 128, "hdd": "128mb", "softwareused": "softwareused", "adminname": "adminname", "adminemail": "adminemail", "admintg": "admintg", "resourceassignment": "resourceassignment", "status": true}
```
