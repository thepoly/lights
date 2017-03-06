<?php
error_reporting(0);
if($_POST['colors'] != ''){
	$colors = explode(',',$_POST['colors']);//expand color variables to an array
	$textfile = "LED.txt"; // Declares the name and location of the .txt file

	$fileLocation = "$textfile";
	$fh = fopen($fileLocation, 'w   ') or die("Something went wrong!"); // Opens up the .txt file for writing and replaces any previous content
	fwrite($fh, $colors[0]."\n");
	fwrite($fh, $colors[1]."\n");
	fwrite($fh, $colors[2]."\n");
	fclose($fh);

	header("HTTP/1.0 200 OK");
}
?>
