# replayd2

A simple web service written in golang
1/ GET / will pull out the customer details
2/ POST or PUT / will set new customer details
3/ It will read replayd.ini and use the listener port to start up the golang built-in web server
4/ tested with -race flag
5/ perf tests with go-wrk

still need to test with -cover test cases
still need to try more deployment tool, e.g. salt