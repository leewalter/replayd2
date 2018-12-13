# replayd2

A simple web service written in golang<br/>
1/ GET / will pull out the customer details<br/>
2/ POST or PUT / will set new customer details<br/>
3/ It will read replayd.ini and use the listener port to start up the golang built-in web server<br/>
4/ tested with -race flag<br/>
5/ perf tests with go-wrk<br/>
<br/>

still need to test with -cover test cases<br/>

bash script at
https://github.com/leewalter/replayd2/tree/master/bash

salt demo at
https://github.com/leewalter/replayd2/blob/master/salt/demo
