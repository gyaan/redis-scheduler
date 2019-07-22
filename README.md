# redis-scheduler

In the software development we interact with many platform and legacy systems. Some of the old system are not capable to handle huge server requests, so we have to create a robust system using new tech and system design which handles all the request but still we need data from the old system and that should sync. 
    So let's say we have one million records and we want to update them regular bases. Lets see what can be the problem.
We can’t update them in single process we need to chunking and then update the records.
We need to rotate the records as well in the circular way.
We need to add the new records as well for update.
If we want to update them in a certain periods then we need to increase the parallel workers if number of records are increasing regularly.


So let’s take example of CRM system. Where we have transactions of user and we want to display those transactions to a new system. For this we need to get transactions regularly bases from CRM. 

We would have we implemented like below if have a mysql table for user.

Create a flag in the user tables. Which can have the following values 
Updated
Updating
Pending
                
So  by default lets its value is pending. When we choose the user for update status will be updating (so that no conflicts in parallel processing), after update it will be updated and when all users are updated we have to make them again pending.

To implement above thing we have to use three database queries and database lock while choosing the users for updates. Which will be a bit slow and difficult to scale it for millions of users(records).

Here comes Redis. We all know redis is in memory cache, fast and easy to implement but Redis has certain feature which fits for such situations.  Let’s try to implement an update records system using Redis. 

Redis has list data type and can be implemented like queue and we can perform certain operation in it like push, pop, lpush (left push), rpush(right push), lpop(left pop), rpop (right pop) etc etc.

Lets create a list called circular_list_for_update which holds all the records for update and another list called current_list_for_update which holds the records which we want to update in specific time (let's say every hours, half hours). Second list is nothing just chunk of records for update process.

//code for creating circular list for update and chunking the records

//code for processing the current list records


Records update process with redis circular queue is fast and less error prone. Fits for worker environment and easily scalable.
