# Interview Assignment - Coop Logistics API

**Implement Server side part that can handle client requests
with specific business logic.**

## Acceptance Criteria

- All three parts must be done.
- API must be capable of handling all requests sent by this client without dropping them.

## Assignment Appendix

Coop Logistics Engine simulates several cargo units that
deliver goods between different warehouses in the world.

It sends `gRPC` requests to `127.0.0.1:50051` (can be redefined),
look at service in
[docker-compose](../docker-compose.yml) "interview_backend_coop_logistics_client"

## Assignment

The result of this task should be an API system that fulfils the criteria
described in three parts below.

*Don't hesitate to ask questions for clarification if you have them.*

**Good luck and we hope you will have fun!**
___

### Part One: Implement Backend API

Implement an API sever that provides the following service over gRPC:
[proto-file](../api/v1/logistics.proto).

- The solution should output a log message to STDOUT
once per second with the number of received messages per second.

### Part Two: Store Delivery Units

The cargo units run between the different warehouses in the world.

Your task is to store these delivery paths,
data comes into the API you implemented in *Part One*.
Write at least one unit test that verifies that your solution works.


The solution should:
- Output a log message to STDOUT once per second
  with the number of received messages per second.
- Output summary
  - Total delivery units.
  - Warehouses that has been supplied (units reached destination).
  - How many delivery units have warehouses.

### Part Three: Export Delivery Paths

Given the delivery paths stored in *Part Two*.

Implement an API endpoint that can export data about warehouses and it's
suppliers.
