<html>

<head>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
<script>
var quotes = [
  "Poly Lights"

];
var q = quotes.length - 1;
var whichquote=Math.round(Math.random()*(q));
var nah = false;;
$('div').on('click', function(){
    alert($(this).attr("id"));
});

function wurd(val,name){
  if(!nah){
    $(val).html($(val).children());
    $(name).fadeIn(500);
  }
};

function changeColor(name,color,letter,dnc){
  nah = true;
  $(name).css('background-color',color);
  $(name).append(letter);
  $(dnc).hide();
  setTimeout(function(){
    nah = false;
}, 1000);
}

$(document).ready(function(){
	$("body").hide();
	$("#demo").html(quotes[whichquote]);
	$("body").fadeIn(1000);
  $(".do-not-clear").hide();
  $(".do-not-clear-2").hide();
  $(".do-not-clear-3").hide();
  $(".do-not-clear-4").hide();
  $.ajax({
   type: "POST",
     url: 'index.html',
     data: {'word=':"aWord"},
     success: function(data) {
       alert(data);

     }
   });
});


</script>
	<title>Poly Lights</title>
  <link rel="stylesheet" type="text/css" href="hover.css">
	<link rel="stylesheet" type="text/css" href="http://fonts.googleapis.com/css?family=Open+Sans">
	<style>
  .title{
		margin: 0 auto;
		width:50%;

		padding-top:0 ;
		color: #2c3e50;
  }
  .color{
    width: 20%;
    height: 40%;
    float: left;
  }
  .fourSelect{
    width:50%;
    height:50%;
    position: relative;
    float: left;

    background-color: black;
  }
	.lbox{
		background-color: blue;
		width: 800px;
		height: 200px;
		margin: 0 auto;
		text-align: center;
		font-size: 100px;
		color:#2C3E50;
	}
	.boxOne {
			margin-left: 0px;
	    width: 200px;
	    height: 200px;
	    background: #4183D7;
	    position: relative;
	    top: 0px;
	    left: 0px;
	}

	.boxTwo {
	    width: 200px;
	    height: 200px;
			background: #22A7F0;
	    position: relative;
	    left: 200px;
	    top: -200px;
	}
	.boxThree {
			width: 200px;
			height: 200px;
			background: #59ABE3;
			position: relative;
			left: 400px;
			top: -400px;
	}
	.boxFour {
			width: 200px;
			height: 200px;
			background: #81CFE0;
			position: relative;
			left: 600px;
			top: -600px;
	}
	.left{
		width:50%;
		height: 100%;
		background-color: #c0392b;
		float:left;
		border-radius: 50%;
	}
	.right{
		width:50%;
		height: 100%;
		float: left;

	}
	body{
    background-color: #ecf0f1;
    background-image: url("canvas.png");
    margin: 0;
    padding: 0;
	  font-family: Open Sans;
		font-style: italic;
		font-variant: normal;
    color:#2C3E50;
	}
	p{
		font-size:100px;
    color:#2C3E50;
	}
	</style>
</head>
<body>

  <div class = "title">
	   <center><p id = "demo"></p></center>
  </div>

		<div class = "lbox">
			<div onclick="wurd('.boxOne','.do-not-clear')" class = "boxOne">P
        <div class = "do-not-clear">
        <div class = "fourSelect" onclick="changeColor('.boxOne','#3498db','P','.do-not-clear') "style = "background-color: #3498db;"></div>
        <div class = "fourSelect" onclick="changeColor('.boxOne','#e74c3c','P','.do-not-clear'); changeColor('.boxTwo','#e74c3c','')" style = "background-color: #e74c3c;"></div>
        <div class = "fourSelect" onclick="changeColor('.boxOne','#2ecc71','P','.do-not-clear')" style = "background-color: #2ecc71;"></div>
        <div class = "fourSelect" onclick="changeColor('.boxOne','#ecf0f1','P','.do-not-clear')" style = "background-color: #ecf0f1;"></div>
      </div>
			</div>
				<div onclick="wurd('.boxTwo','.do-not-clear-2')" class = "boxTwo">O
          <div class = "do-not-clear-2">
          <div class = "fourSelect" onclick="changeColor('.boxTwo','#3498db','O','.do-not-clear-2') "style = "background-color: #3498db;"></div>
          <div class = "fourSelect" onclick="changeColor('.boxTwo','#e74c3c','O','.do-not-clear-2');changeColor('.boxOne','#e74c3c','')" style = "background-color: #e74c3c;"></div>
          <div class = "fourSelect" onclick="changeColor('.boxTwo','#2ecc71','O','.do-not-clear-2')" style = "background-color: #2ecc71;"></div>
          <div class = "fourSelect" onclick="changeColor('.boxTwo','#ecf0f1','O','.do-not-clear-2')" style = "background-color: #ecf0f1;"></div>
        </div>
				</div>
					<div onclick="wurd('.boxThree','.do-not-clear-3')" class = "boxThree">L
            <div class = "do-not-clear-3">
            <div class = "fourSelect" onclick="changeColor('.boxThree','#3498db','L','.do-not-clear-3') "style = "background-color: #3498db;"></div>
            <div class = "fourSelect" onclick="changeColor('.boxThree','#e74c3c','L','.do-not-clear-3')" style = "background-color: #e74c3c;"></div>
            <div class = "fourSelect" onclick="changeColor('.boxThree','#2ecc71','L','.do-not-clear-3')" style = "background-color: #2ecc71;"></div>
            <div class = "fourSelect" onclick="changeColor('.boxThree','#ecf0f1','L','.do-not-clear-3')" style = "background-color: #ecf0f1;"></div>
          </div>
					</div>
						<div onclick="wurd('.boxFour','.do-not-clear-4')" class = "boxFour">Y
              <div class = "do-not-clear-4">
              <div class = "fourSelect" onclick="changeColor('.boxFour','#3498db','Y','.do-not-clear-4') "style = "background-color: #3498db;"></div>
              <div class = "fourSelect" onclick="changeColor('.boxFour','#e74c3c','Y','.do-not-clear-4')" style = "background-color: #e74c3c;"></div>
              <div class = "fourSelect" onclick="changeColor('.boxFour','#2ecc71','Y','.do-not-clear-4')" style = "background-color: #2ecc71;"></div>
              <div class = "fourSelect" onclick="changeColor('.boxFour','#ecf0f1','Y','.do-not-clear-4')" style = "background-color: #ecf0f1;"></div>
            </div>
						</div>
		</div>
</body>
</html>
