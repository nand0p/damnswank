package handlers

import (
    "fmt"
    "net/http"
)

func DamnswankHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><head><title> ---=== !!!!DAMN SWANK!!!! ===--- </title>")
    fmt.Fprintf(w, "<link rel=icon href=https://s3.amazonaws.com/damnswank/favicon.ico type=image/x-icon />")
    fmt.Fprintf(w, "<script type=text/javascript>var _gaq=_gaq||[];_gaq.push(['_setAccount','UA-37443405-1']);_gaq.push(['_trackPageview']);")
    fmt.Fprintf(w, "(function(){var ga=document.createElement('script');ga.type='text/javascript';ga.async=true;")
    fmt.Fprintf(w, "ga.src=('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';")
    fmt.Fprintf(w, "var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);")
    fmt.Fprintf(w, "})();</script></head><body bgcolor=DD3322 text=pink vlink=blue>") 
    fmt.Fprintf(w, "<center><font face=Helvetica><p><br><p><b><h1> &nbsp; </h1></b><p><br><p>")
    fmt.Fprintf(w, "<a href=http://www.damnswank.com><img src=https://s3.amazonaws.com/damnswank/damnswank.gif />")
    fmt.Fprintf(w, "<p>&copy; 2001-2016 DAMNSWANK<br><p></a>9933CC</body></html>")
}
