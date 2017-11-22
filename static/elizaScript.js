/* $("#user-input-form").submit(function(event){
          event.preventDefault();
          $.get('/user-input', { value: $('#user-input').val() } )
            .done(function (data) {
              $(".containerM").append($("div"));
              //$('#output-area').val(data);
            })
        });*/
        
        /*$(document).ready(function(){
          $("#submitBtn").click(function(){
            $(".containerM").append($("div"));
            $( "div" ).text( value );
          });

        });*/
       function loadDoc() {
          var xhttp = new XMLHttpRequest();
          var userInput, userOutput;
          
          xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
              // Adapted from: https://stackoverflow.com/questions/14094697/javascript-how-to-create-new-div-dynamically-change-it-move-it-modify-it-in
              
              var iDiv = document.getElementById('message-container')
              document.getElementsByTagName('body')[0].appendChild(iDiv);
              var divInput = document.createElement('div');
              var divOutput = document.createElement('div');
              divInput.textContent = "User: " + document.getElementById('user-input').value;
              
              divInput.id = 'input-areaEliza';
              divInput.setAttribute("role", "alert")
              divInput.setAttribute("class", "alert alert-danger");
              iDiv.appendChild(divInput);

              console.log(divInput);
              divOutput.id = 'output-areaEliza';
              divOutput.setAttribute("role", "alert")
              divOutput.setAttribute("class", "alert alert-info");
              iDiv.appendChild(divOutput);

              console.log(divOutput);
              divOutput.textContent = xhttp.responseText;
              window.scrollTo(0, document.body.scrollHeight);
            }
          };
          xhttp.open("POST", "user-input");
          xhttp.setRequestHeader("Cache-Control", "no-cache");
          xhttp.setRequestHeader("value", document.getElementById("user-input").value);
          xhttp.send();
        }

        $("#user-input").keydown((evt)=>{
          if (evt.keyCode == 13){
              loadDoc();
          }});

        window.scrollTo(0, document.body.scrollHeight);