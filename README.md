# READ me

My name is Kevin Delassus and this is my 3rd Year Data Representation and Querying Eliza Project. The instructions were to create a chatbot web application in GO. Users must be able to access the program through their web browsers.
The chatbot should be guided the Eliza Program. If we wished we could use any technique to enhance your chatbot, such as machine learning methods.

## Running Eliza
 To run the Eliza web app you first need to make sure that you have GO on your PC. If you dont go to the link below and download and install GO.
 
    https://golang.org/
    

To complete the next step GIT is also required on you PC. If you don't have GIT installed go to the link below and download and install GIT

    https://git-scm.com/

Once both has been installed you are ready for Eliza!
Open your CMD again and navigate to where you would like to clone the Eliza repository. To clone the repository enter:
```sh
    $ git clone https://github.com/kevind992/ChatBot_Web_Application.git
```     
Once the repository has been cloned, navigate into the Eliza folder.
To build the GO file enter:
```sh
   $ go build Chatbot.go
```    
To run the program, enter:    
```sh
    $ Chatbot.exe
```
You not need to open your browser and in the address bar type:

    localhost:8080

The Eliza program should now be running on the screen.

To use simply enter what you which to say to Eliza, press send and she should reply back to you. Enjoy her wise wisdom! 

### Who and What is Eliza?

Eliza is a computer program that emulates a Rogerian psychotherapist. When the original Eliza first appeared in the 60's, some people actually mistook her for human. The illusion of intelligence works best, however, if you limit your conversation to talking about yourself and your life.

### Layout of Filing

The Eliza folder contains 2 Go files, one which runs the server and the second which contains all the code for Eliza herself.
Inside the static folder contains a HTML file and a JavaScript file.
Inside the index.html contains all the code for the layout.
Inside the JavaScript file contains all the code for sending the data to and from the Go file. It also creates new dynamic divs to display the input and response to the user.
Finally, the rest of the files included are the .gitignore, license and the READme which you are currently reading.

### Software's Used

To create the user interface I used bootstrap
    https://getbootstrap.com/
Bootstrap is an open source toolkit for developing with HTML, CSS, and JS. Quickly prototype your ideas or build your entire app with our Sass variables and mixins, responsive grid system, extensive prebuilt components, and powerful plugins built on jQuery. 

Go programming language was used for the server and Eliza herself.
https://golang.org/
Go is a programming language created at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing, memory safety features and CSP-style concurrent programming features added. The compiler and other language tools originally developed by Google are all free and open source.

JavaScript was used for running in the background. It would take and output the user input and output.
https://www.javascript.com/

### Areas for Improvement

In the future I would love to expand Eliza's intelligence. I would like to make her more real like. (Not so blunt!)
I would also like to improve my JavaScript by adding jQuery.
I attempted this but failed to get it working before the project deadline so reverted back to plain JavaScript.

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
