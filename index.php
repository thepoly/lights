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
<html>
<head>
	<script src="//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js"></script>
	
<script src='spectrum.js'></script>
<link rel='stylesheet' href='spectrum.css' />
	</head>
<body>
	<script>

	function getRGB(color) {
	    var result;
	    // Look for rgb(num,num,num)
	    if (result = /rgb\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*\)/.exec(color)) return [parseInt(result[1]), parseInt(result[2]), parseInt(result[3])];
	    // Look for rgb(num%,num%,num%)
	    if (result = /rgb\(\s*([0-9]+(?:\.[0-9]+)?)\%\s*,\s*([0-9]+(?:\.[0-9]+)?)\%\s*,\s*([0-9]+(?:\.[0-9]+)?)\%\s*\)/.exec(color)) return [parseFloat(result[1]) * 2.55, parseFloat(result[2]) * 2.55, parseFloat(result[3]) * 2.55];
	    // Look for #a0b1c2
	    if (result = /#([a-fA-F0-9]{2})([a-fA-F0-9]{2})([a-fA-F0-9]{2})/.exec(color)) return [parseInt(result[1], 16), parseInt(result[2], 16), parseInt(result[3], 16)];
	    // Look for #fff
	    if (result = /#([a-fA-F0-9])([a-fA-F0-9])([a-fA-F0-9])/.exec(color)) return [parseInt(result[1] + result[1], 16), parseInt(result[2] + result[2], 16), parseInt(result[3] + result[3], 16)];
}
	</script>
	
	<center>Select A color for the lights-> <input type='text' id="minipicker"/></center>
<script>
$("#minipicker").spectrum({
    color: "#ECC",
    showInput: true,
    className: "full-spectrum",
    showInitial: true,
    showPalette: true,
    showSelectionPalette: true,
    maxPaletteSize: 10,
    preferredFormat: "hex",
    localStorageKey: "spectrum.demo",
    move: function (color) {
        
    },
    show: function () {
    
    },
    beforeShow: function () {
    
    },
    hide: function () {
    
    },
    change: function() {
       $.ajax({
				type: "POST",
			  	url: 'index.php',
			  	data: {'colors':getRGB($('#minipicker').val()).join(',')},
			  	success: function(data) {
			    	//alert("changed, but you wont know unless technophilia is broadcasting live!");
			  	}
			}); 
    },
    palette: [
        ["rgb(0, 0, 0)", "rgb(255, 0, 0)", "rgb(0, 255, 0)",
        "rgb(0, 0, 255)"]
    ]
});
</script>
	</body>
</html>
