// Author: Kevin Delassus
//This file contains all the JavaScript code for my Eliza program

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
          
          xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {

              // Adapted from: https://stackoverflow.com/questions/14094697/javascript-how-to-create-new-div-dynamically-change-it-move-it-modify-it-in
              //Creating a message-container object
              var messageDiv = document.getElementById('message-container')
              //Making it the parent
              document.getElementsByTagName('body')[0].appendChild(messageDiv);
              //Creating two new div elements
              var divInput = document.createElement('div');
              var divOutput = document.createElement('div');
              //Assigning the users input into divInput div
              divInput.textContent = "User: " + document.getElementById('user-input').value;
              
              //Adding attributes and adding it to the parent div
              divInput.id = 'input-areaEliza';
              divInput.setAttribute("role", "alert");
              //Assigning bootstrap to the div
              divInput.setAttribute("class", "alert alert-danger");
              messageDiv.appendChild(divInput);

              //Adding attributes and adding it to the parent div 
              divOutput.id = 'output-areaEliza';
              divOutput.setAttribute("role", "alert");
              //Assigning bootstrap to the div
              divOutput.setAttribute("class", "alert alert-info");
              messageDiv.appendChild(divOutput);

              //Assigning the output div with the Eliza responce which was coded with GO
              divOutput.textContent = xhttp.responseText;

              //Once the user as entered and submitted a responce the textbox is cleared
              //Adapted from: https://stackoverflow.com/questions/1852061/clear-textbox-after-form-submit
              document.getElementById("user-input").value = "";
              //Keeps the screen the the current message. Stops user from having to scroll
              //Adapated from: https://www.w3schools.com/jsref/prop_element_scrolltop.asp
              window.scrollTo(0, document.body.scrollHeight);
            }
          };
          //Sending the user input to the go code
          xhttp.open("POST", "user-input");
          xhttp.setRequestHeader("Cache-Control", "no-cache");
          xhttp.setRequestHeader("value", document.getElementById("user-input").value);
          xhttp.send();
        }

        //jQuery for the user to be able to press enter instead of having to use the mouse.
        $("#user-input").keydown((evt)=>{
          if (evt.keyCode == 13){
              loadDoc();
          }});
