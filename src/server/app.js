function updateColors(){
  $.get( "LEDL.txt", function( data ) {
    data = data.split("\n");
    console.log(data);
    $(".sign-item").css("background-color", "rgb(" + data[0] + "," +data[1]+ ","+data[2] +")");
  });
}

$(document).ready(function(){
  $(".sign-item").height($(".sign-item").width());

  updateColors();
  $(".color-picker").empty();
  for (var i = 0; i < 360; i +=20){
    $(".color-picker").append("<div class='color' id=" + i + " style='background-color:hsl("+i+",100%,50%);float:left; position:relative; width:" + $(".color-picker").width()/18 +"; height:100%;'></div>");
  }

  $(".color-picker").show();
  for (var i = 0; i < 1000; i ++){
    $(".color-picker").animate({'height': '8rem'}, 100);
  }

  $(".color").click(function(){
    var newColor = ($("#" + this.id).css("background-color").substring(4));
    newColor = newColor.substring(0, newColor.length-1)
    newColor = newColor.replace(/ /g,'')
    console.log(newColor);
    //$("#" + this.id);
    $.post( "submit?word=" + newColor + ",0", function( data ) {
      console.log(data);
      updateColors();
    });

  });

});

$( window ).resize(function() {
  $(".sign-item").height($(".sign-item").width());

  updateColors();
  $(".color-picker").empty();

  for (var i = 0; i < 360; i +=20){
    $(".color-picker").append("<div class='color' id=" + i + " style='background-color:hsl("+i+",100%,50%);float:left; position:relative; width:" + $(".color-picker").width()/18 +"; height:100%;'></div>");
  }

  $(".color-picker").show();
  for (var i = 0; i < 1000; i ++){
    $(".color-picker").animate({'height': '8rem'}, 100);
  }

  $(".color").click(function(){
    var newColor = ($("#" + this.id).css("background-color").substring(4));
    newColor = newColor.substring(0, newColor.length-1)
    console.log(newColor);
    //$("#" + this.id);
    $.post( "submit?word=" + newColor + ",0", function( data ) {
      console.log(data);
      updateColors();
    });

  });

});
