**uuidgen**
Golang/Go based uuid generator for a distributed system with requirements as below
  1. Unsigned 64 bits integer (non alphanumeric)
  2. uuid values increases with the time
      uuid at timestamp t1 < uuid t2,given timestaps t1 < t2
  3. Generation capactiy : 1000 uuids/miliseconds

**Usage**

Install the prorgram as below,
**go install github.com/deepakkumarbhagat/uuidgen/cmd/uuidgen@latest**

The program **uuidgen** thus installed, is a rest server supporting only a GET method at the endpoint "/apis/v1/uuids" and port 8080.
The server acheives generating capacity of 1000 uuids/milisecs by implementing fan-out/worker pattern with go routines.

![image](https://github.com/user-attachments/assets/90db21da-bf3f-4e7e-8962-68b9a6e7e10c)

**Testing**

Tool to send rest request to the server , it's is advisable to try simulating burst of generation request 
to the server and verify if the server absorbs the shock correctly.

![image](https://github.com/user-attachments/assets/d7c029eb-498f-4ca5-b46a-1473aef34495)


**CLI**
CLI support TBD
