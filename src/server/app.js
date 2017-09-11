$(document).ready(function(){
  $(".sign-item").height($(".sign-item").width());
  for (var i = 0; i < 360; i +=10){
    $(".color-picker").append("<div class='color' id=" + i + " style='background-color:hsl("+i+",100%,50%);float:left; position:relative; width:" + $(".color-picker").width()/36 +"; height:100%;'></div>");
  }
});

$( window ).resize(function() {
  $(".sign-item").height($(".sign-item").width());

});
