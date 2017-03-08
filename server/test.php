<?php
$colors = explode(',',$_POST['word']);//expand color variables to an array

$wurd = $colors[3];

if ($wurd != ''){

}
$textfile = "LED" . $wurd . ".txt"; // Declares the name and location of the .txt file
echo $textfile;

$fileLocation = "$textfile";
$fh = fopen($fileLocation, 'w   ') or die("Something went wrong!"); // Opens up the .txt file for writing and replaces any previous content
fwrite($fh, $colors[0]."\n");
fwrite($fh, $colors[1]."\n");
fwrite($fh, $colors[2]."\n");
fclose($fh);

header("HTTP/1.0 200 OK");
?>
