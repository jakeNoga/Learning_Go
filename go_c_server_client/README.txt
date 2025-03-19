Ok this is more complex then it really needs to be for the simple example but this essentially will prove that it
will work for a larger project.


What I need is a compiled c library that is called by a c program that makes a call to my go server. My go server 
I want to run c and go code. I want interactions between the client and the server to be in go. How I am thinking 
of accomplishing this is first creating a go library that will go on top of my c library that essentially will 
handle translating the c struct to a grpc call and transferring the response back to c. 

My go server will take the calls in the form of grpc's. I want 1 function to be in go and another to be in c.