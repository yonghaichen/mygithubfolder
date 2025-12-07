

- Coils (1 bit): readable and writable, states: on/off
- Discrete Inputs (1 bit): readable, states: on/off
- Input Registers (16 bit): readable, essentially measurements and statuses
- Holding Registers (16 bit): readable and writable, essentially configuration values





Function codes:  

```
* Read discrete inputs (0x02)                      
* Read           coils (0x01)     
* Write single   coil  (0x05)         
* Write multiple coils (0x0f)         


* Read           input   registers (0x04)   
* Read           holding registers (0x03)    
  Read multiple  holding registers  
* Write single   holding register  (0x06)    
* Write multiple holding registers (0x10)       
```