# Retrieving Result Sets in SQL
- Don’t [defer](http://go-database-sql.org/retrieving.html) within a loop. 
- A deferred statement doesn’t get executed until the function exits, so a long-running function shouldn’t use it. 
- If you do, you will slowly accumulate memory. 
- If you are repeatedly querying and consuming result sets within a loop, you should explicitly call rows.Close() when you’re done with each result, and not use defer.