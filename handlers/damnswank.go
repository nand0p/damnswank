package handlers

import (
    "fmt"
    "net/http"
    "math/rand"
    "strconv"
)

const basecolor = "f3151d"
const randomness = 10000
const refresh = "4"

func DamnswankHandler(w http.ResponseWriter, r *http.Request) {
    colorint, err := strconv.ParseInt(basecolor, 16, 64)
    if err != nil { fmt.Println("strconv.ParseInt failed:", err) }
    var random int64 = int64(rand.Intn(randomness) - randomness / 2)
    var mutate int64 = colorint + random
    var mutatehex string = strconv.FormatInt(mutate, 16)
    fmt.Print(" BASEHEX: " + basecolor + " BASEINT: " + strconv.Itoa(int(colorint)) + " RANDOM: " + strconv.Itoa(int(random)) +
              " MUTATE: " + strconv.Itoa(int(mutate)) + " MUTATEHEX: " + mutatehex + " :: ")
    fmt.Fprintf(w, "<html><head><title> ---=== !!!!DAMN SWANK!!!! ===--- </title>" +
                   "<style type=text/css>a:link, a:visited, a:active {text-decoration: none; color: " + basecolor + ";} " +
                   "a:hover {color: black}</style><meta http-equiv=refresh content=" + refresh + ">" +
                   "<link rel=\"shortcut icon\" href=https://s3.amazonaws.com/damnswank/favicon.ico type=image/x-icon>" +
                   "<script type=text/javascript>var _gaq=_gaq||[];_gaq.push(['_setAccount','UA-37443405-1']);_gaq.push(['_trackPageview']);" +
                   "(function(){var ga=document.createElement('script');ga.type='text/javascript';ga.async=true;" +
                   "ga.src=('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';" +
                   "var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);})();</script></head>" +
                   "<body bgcolor=" + mutatehex + "><p><br><p><b><h1>&nbsp;</h1></b><p><br><p><center>" +
                   "<img src=https://s3.amazonaws.com/damnswank/damnswank.gif><font color=" + mutatehex + ">" +
                   mutatehex + "</font><p><font color=" + basecolor + ">&copy; 2001-2016&nbsp;</font>" +
                   "<a href=http://www.damnswank.com>DAMNSWANK</a><br><p></body></html>")
}
