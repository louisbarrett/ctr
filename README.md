
#content-type header rewrite

This is a content-type proxy that rewrites data from the a url to the desired content-type.  

# Usage:  

To use this proxy, send a GET request to the /api/{content-type}/ endpoint with a query parameter of "url".  

For example:  

/api/text/  

?url=https://www.google.com  

This will return the content-type of the Google homepage as "text/html".


# Installation:  

To install this proxy, simply run:  

go install github.com/louisbarrett/ctr@latest

# Configuration:  

The only configuration required is to set the webServerPort variable to the desired port.  

By default, the webServerPort is set to ":8080". 

