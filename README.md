**uuid-gen**
Golang/Go based uuid generator for a distributed system with requirements as below
  1. Unsigned 64 bits integer (non alphanumeric)
  2. uuid values increases with the time
      uuid at timestamp t1 < uuid t2,given timestaps t1 < t2
  3. Generation capactiy : 1000 uuids/miliseconds

**Usage**

The suggestion is to run the uuid-gen as a sidecar container along with the main container in the cloud cluster.
The generator supports the restful API for getting the uuid at the endpoint : /apis/v1/uuids

**Testing**

curl/postman utilities to generate GET http/https request with endpoint /apis/v1/uuids

**CLI**

CLI support TBD
