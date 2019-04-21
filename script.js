var keywords = ["MAXIMISE", "MAXIMIZE", "MUTE", "CHROME"]

function request(method,command){
  var success = false;
  fetch("http://localhost:8080/?command=" + encodeURI(command), {method: method})
  .then(
    function(response){
      if (response.status == 200){
        $('.email input').parent().addClass('success');
        console.log("switched");
      } else {
        $('.email input').parent().removeClass('success');
      }
    })
};


function launch(command){
  //linux addin commands
  var commandarr = command.split(" ");
  var keyword = commandarr[0].toUpperCase();

  //maximise
  if (keyword == "MAXIMISE" || keyword == "MAXIMIZE"){
    commandarr[0] = "wmctrl -a";
  }

  //mute
  if (keyword == "MUTE") {
    commandarr[0] = "amixer -q -D pulse sset Master toggle";
  }

  //chrome
  if (keywords == "CHROME"){
    commandarr[0] = "google-chrome";
  }


  command = commandarr.join(" ");
  console.log(command)
  request("POST", command);
  setTimeout(function(){$('.email input').parent().removeClass('success');}, 1000);
};



function getValue(id){
  return document.getElementById(id).value;
}

function search(){
  var t = getValue('searchbar');
  console.log(t);
  launch(t);
  $(".email input").val("");
}

function verifyCommand(command){
  if (request("GET", command) == true || keywords.includes(command.toUpperCase().split(" ")[0]) == true){
    console.log("success");
    return true;
  } else{
    return false;
  }
}

function answer(e) {
    if (e.keyCode == 13) {
      search();
     }
}

document.addEventListener('keydown', function(event) {
  document.getElementById("searchbar").focus();
}, true);
