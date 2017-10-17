
function updateColors(){
  $.get( "LEDP.txt", function( data ) {
    data = data.split("\n");
    $(".project-title").css("background-color", "rgb(" + data[0] + "," +data[1]+ ","+data[2] +")");
  });
}

function initDocument(){
  var mobile = (/iphone|ipod|android|blackberry|mini|windows\sce|palm/i.test(navigator.userAgent.toLowerCase()));

  updateColors();
  $(".color-picker").empty();
  for (var i = 0; i < 360; i +=40){
    $(".color-picker").append("<div class='color' id=" + i + " style='background-color:hsl("+i+",100%,50%);float:left; position:relative; width:" + $(".color-picker").width()/9 +"; height:100%;'></div>");
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

}

$(document).ready(function(){
  initDocument();
  var i = setInterval(updateColors, 1000);
});

$( window ).resize(function() {
  initDocument();

});
