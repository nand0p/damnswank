<html><head> 
<title> 
---=== !!!!DAMN SWANK!!!! ===---
</title>
<script type="text/javascript">

  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-37443405-1']);
  _gaq.push(['_trackPageview']);

  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();

</script>
</head>


<body bgcolor=9933CC text=pink vlink=blue>
<center><font face=Helvetica><p><br><p><b><h1>
&nbsp; 
</h1></b><p><br><p>
<a href=http://www.damnswank.com/wordpress>
<img src=http://www.damnswank.com/damnswank.gif>
<p>&copy; 2001-2015 DAMNSWANK<br><p></a>

<?php
        $COUNT_FILE = "counter.txt";
        $IMG_DIR_URL = "http://www.damnswank.com/digits/";
        $NB_DIGITS = 8;
        if (file_exists($COUNT_FILE)) {
                $fp = fopen("$COUNT_FILE", "r+");
                flock($fp, 1);
                $count = fgets($fp, 4096);
                $count += 1; 
                fseek($fp,0);
                fputs($fp, $count);
                flock($fp, 3);
                fclose($fp);
        }
        else {
                echo "Can't find file, check '\$file' var...<BR>";
                exit;
        }
        chop($count);
        $nb_digits = max(strlen($count), $NB_DIGITS);
        $count = substr("0000000000".$count, -$nb_digits);
        $digits = preg_split("//", $count);
        for($i = 0; $i <= $nb_digits; $i++) {
                if ($digits[$i] != "") {
                        $html_result .=  "<IMG SRC=\"https://s3.amazonaws.com/damnswank/counter/[$i].gif\">";
                }
        } 
        echo $html_result; 
?>
</body>
</html>

