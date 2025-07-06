# Defer
- A [defer](https://go.dev/tour/flowcontrol/12) statement defers the execution of a function until the surrounding function returns.
- The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

# Don't use defer in for-loop
- Don’t [defer](http://go-database-sql.org/retrieving.html) within a loop.
- A deferred statement doesn't get executed until the function exits, so a long-running function shouldn’t use it.
- If you do, you will slowly accumulate memory.
- If you are repeatedly querying and consuming result sets within a loop, you should explicitly call rows.Close() when you’re done with each result, and not use defer.