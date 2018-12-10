# replayd2

A simple web service written in golang\n
1/ GET / will pull out the customer details\n
2/ POST or PUT / will set new customer details\n
3/ It will read replayd.ini and use the listener port to start up the golang built-in web server\n
4/ tested with -race flag\n
5/ perf tests with go-wrk\n
\n
still need to test with -cover test cases\n
still need to try more deployment tool, e.g. salt\n