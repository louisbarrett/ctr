
# Content-Type Rewrite Proxy

This is a content-type rewrite proxy that rewrites data from the a url to the desired content-type.  

# Usage:  
```
Usage of ./ctr:
  -content-type string
        Default content type to use if not specified in the url (default "application/json")
  -port string
        Port to listen on (default ":8080")
 ```

To use this proxy, send a GET request to the /api/{content-type}/ endpoint with a query parameter of "url".  

For example:  

/api/application-pdf/  

?url=https://github.com/angea/pocorgtfo/raw/master/contents/articles/05-05.pdf  

This will return the content-type of the Google homepage as "text/html".


# Installation:  

To install this proxy, simply run:  

go install github.com/louisbarrett/ctr@latest

# Configuration:  

The only configuration required is to set the webServerPort variable to the desired port.  

By default, the webServerPort is set to ":8080". 

By default, the content-type is set to application/json when no url path is used. 
This behavior can be overridden using the `-content-type` parameter
