function populatePicker(){
  $(".cpicker").empty();

  for (var i = 0; i < 360; i +=90){
    $(".cpicker").append("<div class='color' id=" + i + " style='background-color:hsl("+i+",100%,50%);float:right; position:relative; width:" + ($(".cpicker").width())/4 +"; height:100%;'></div>");
  }
}

function updatePageColor(){
  $.get( "LEDP.txt", function( data ) {
    data = data.split("\n");
    $(".header").css("background-color", "rgb(" + data[0] + "," +data[1]+ ","+data[2] +")");
  });
}

var colorPicker = Vue.component('color-picker', {
  template: `
    <div>
      <div v-on:click='updateColor' class="cpicker" style="padding-left:0;height:25%;width:100%;border-radius:5px; border-color:black; border-style: solid;float:left;"></div>
    </div>`
    ,
  methods:{
    colorChange: function(e){
      this.populate(e)
    },
    updateBg: function(){
      updatePageColor();
    },
    updateColor: function(e){
      var newColor = ($(e.target).css("background-color").substring(4));
      newColor = newColor.substring(0, newColor.length-1)
      newColor = newColor.replace(/ /g,'')
      var comp = this;
      $.post( "submit?word=" + newColor + ",0", function( data ) {
        comp.updateBg();
      });
    }
  },
  mounted(){
    populatePicker();
    updatePageColor();

    setInterval(function(){
      updatePageColor();
    },1000);
  }

})

$( window ).resize(function() {
  console.log("asdf");
  populatePicker();
});

Vue.component('app-body', {
  template: `

  <div>
    <div class="header"><img class="title-logo" src="images/logo_m.png"/></div>
    <div class ="project-title"><p class="main-header">PolyLights</p></div>
    <div style="margin: 0 auto; width:75%;top:50px;position:relative;"><color-picker/></div>
    <div class="sub-text">Change the color of <a class="link" href="https://poly.rpi.edu">The Rensselaer Polytechnic's</a> Lights in the Rensselaer Student Union.
    </div>
`
})


var app = new Vue({
  el: '#app',
  data: {

  },

  methods:{
  }
})
