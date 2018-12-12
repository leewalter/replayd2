#!/bin/bash
# Walter Lee  	2018-12-12   bash shell to do some curl GET/POST with functions and error checking

# if debug, then uncomment below -x, 
#set -x 

#tells script to exit as soon as any line in the script fails
set -e 

# Define GET function, $1 is url string

getURL() {

#GET
# both below are global vars

        output=`curl -s $1`
        status=$?
      
}

# Define POST function, $1 is url string
postURL() {

#POST
        output=`curl -s --location --request POST $1 --data "Walter testing Postman Echo post body."`
        status=$?
}        

# 1st set with a POST
# do POST to postman echo test
postURL "https://postman-echo.com/post"
printf "\nWalter Lee just did a POST test\n\n"

if [ $status -eq 0 ]
then
  printf  "Post Postman Echo output is \n$output\n"
else
  printf  "curl POST failed !error code is $status\n"
fi 

#reset status = 0 and output = "" because they are global vars.
status=0
output=""

#2nd check results with a GET at Postman Echo 
getURL "https://postman-echo.com/get?foo1=bar1&foo2=bar2"
printf "\njust did a GET test @Postman Echo\n\n"

if [ $status -eq 0 ]
then
  printf "Get Postman Echo is $output \n"
else
  printf "curl GET failed ! error code is $status\n"
fi 


