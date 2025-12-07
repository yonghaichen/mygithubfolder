 
```
Client:  FIN=1         seq=x  
Server:          ACK=1 ack=x+1          seq=y   

Server:  FIN=1   ACK=1 ack=x+1          seq=z   
Client:                seq=x+1    ACK=1 ack=z+1   
```
