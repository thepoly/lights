var colorPicker = Vue.component('color-picker', {
  template: `
  <div class="picker" v-bind:style="bgc">
    <div class="side-container">
    <div class="main-content">
    <div class="project-title">
    <p class = "main-header">PolyLights</p>
    </div>
    </div>

    </div>
    <div style="float:center;position:absolute;top:50%;left:50%;transform:translate(-50%,-50%);height:100px;width:200px;margin: 0 auto;">
    <input class="color_field" v-model="message" @input="update" placeholder="#FFFFFF">
    <img class="important" v-bind:style="important" src="/images/important.png"></img>
    </div>
  </div>
  `,
  data() {
		return {
			message: "",
      bgc: {
        backgroundColor: "#F00FFF",

      },
      important: {
        position: "absolute",
        top:5000,
        left:200,
      },
      changing: false,
		}
	},
  created: function(){
    var thing = this;
    setInterval(function(){
      if(!thing.changing){
        thing.updateColor(true);
      }
    }, 300);

  },
  methods: {
    updateColor: function(change){
      var thing = this;
      $.get( "LEDP.txt", function( data ) {
        data = data.split("\n");
        if(change)
        {
          thing.bgc.backgroundColor = "rgb(" + data[0] + "," +data[1]+ ","+data[2] +")";
        }
        console.log(thing.bgc.backgroundColor);
      });
    },
    postColor: function(){
      var newColor = ($(".picker").css("background-color").substring(4));
      newColor = newColor.substring(0, newColor.length-1)
      newColor = newColor.replace(/ /g,'')
      //$("#" + this.id);
      $.post( "submit?word=" + newColor + ",0", function( data ) {
        this.updateColor;
      });
    },

    update: function(e){
      this.important.top = 5000;
      var thing = this
      if(this.message != ""){
        this.changing = true;
        this.bgc.backgroundColor=this.message;
        setTimeout(function(){
          thing.postColor();
          setTimeout(function(){
            thing.changing = false;
          }, 1000);
        }, 400);

      }
      if(this.message == "s"){
        this.important.top = 300;
      }
      if(this.message == "sh"){
        this.important.top = 275;
      }
      if(this.message == "shi"){
        this.important.top = 250;
      }
      if(this.message == "shir"){
        this.important.top = 225;
      }
      if(this.message == "shirl"){
        this.important.top = 200;
      }
      if(this.message == "shirle"){
        this.important.top = 175;
      }
      if(this.message == "shirley"){
        this.important.top = 150;
      }

    }
  }
})

Vue.component('app-body', {
  template: `

  <div>
    <div class="header"><img class="title-logo" src="images/logo_m.png"/></div>

    <color-picker/>
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
