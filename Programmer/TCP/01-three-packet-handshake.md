
```
• Client     SYN=1          seq=`x`     
• Server     SYN=1    ACK=1 ack=`x+1`           seq=`y`  
• Client                    seq=x+1       ACK=1 ack=`y+1`      
```