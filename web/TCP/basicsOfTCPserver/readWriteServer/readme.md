# TCP servers
 
 ..* Creating a server is a three step process
    1. listen to some port,
    2. accept or reject(close) it
    3. start reading and writing from it

 > It's like how you use your cellphone. You a chime (notification) you then decide whether to accept that notification and do something or just ignore. Once, you accept it (if its a text you write back to it or if its a call you answer back(audio writing))   

..* Few points to remember :

 An open connection blocks.
 The reader is reading from the open connection.
 The writer writes to this open connection
 But, 
 How does the reader know when it should stop reading?

 find out in this program.
 Run the main.go and go to http://localhost:8080/