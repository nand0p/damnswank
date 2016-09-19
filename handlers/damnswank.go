package handlers

import (
    "fmt"
    "net/http"
    "math/rand"
    "strconv"
)

const basecolor = "f3151d"
const randomness = 10000

func DamnswankHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Print(" BASEHEX: ")
    fmt.Print(basecolor)
    colorint, err := strconv.ParseInt(basecolor, 16, 64)
    if err != nil { fmt.Println("strconv.ParseInt failed:", err) }
    fmt.Print(" BASEINT: ")
    fmt.Print(colorint)
    var random int64 = int64(rand.Intn(randomness) - randomness / 2)
    fmt.Print(" RANDOM: ")
    fmt.Print(random)
    var mutate int64 = colorint + random
    fmt.Print(" MUTATE: ")
    fmt.Print(mutate)
    var mutatehex string = strconv.FormatInt(mutate, 16)
    fmt.Print(" MUTATEHEX: ")
    fmt.Print(mutatehex)
    fmt.Fprintf(w, "<html><head><title> ---=== !!!!DAMN SWANK!!!! ===--- </title>")
    fmt.Fprintf(w, "<style type=text/css>a {text-decoration: none}</style>")
    fmt.Fprintf(w, "<link rel=icon href=https://s3.amazonaws.com/damnswank/favicon.ico type=image/x-icon>")
    fmt.Fprintf(w, "<script type=text/javascript>var _gaq=_gaq||[];_gaq.push(['_setAccount','UA-37443405-1']);_gaq.push(['_trackPageview']);")
    fmt.Fprintf(w, "(function(){var ga=document.createElement('script');ga.type='text/javascript';ga.async=true;")
    fmt.Fprintf(w, "ga.src=('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';")
    fmt.Fprintf(w, "var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);")
    fmt.Fprintf(w, "})();</script></head><body bgcolor=")
    fmt.Fprintf(w, mutatehex)
    fmt.Fprintf(w, " link=")
    fmt.Fprintf(w, basecolor)
    fmt.Fprintf(w, " alink=")
    fmt.Fprintf(w, basecolor)
    fmt.Fprintf(w, " vlink=")
    fmt.Fprintf(w, basecolor)
    fmt.Fprintf(w, "><p><br><p><b><h1>&nbsp;</h1></b><p><br><p><center><img src=https://s3.amazonaws.com/damnswank/damnswank.gif><font color=")
    fmt.Fprintf(w, mutatehex)
    fmt.Fprintf(w, ">")
    fmt.Fprintf(w, mutatehex)
    fmt.Fprintf(w, "</font><p><font color=")
    fmt.Fprintf(w, basecolor)
    fmt.Fprintf(w, "> &copy; 2001-2016 </font><a href=http://www.damnswank.com>DAMNSWANK</a><br><p></body></html>")
}
