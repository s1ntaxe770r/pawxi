## Whats this ? 

Pawxi is yet another reverse proxy desinged with simplicity in mind. Born out of a certain users frustration at the complexity of certain proxies. 

## Cool how does it work ? 

Simply create a `pawxi.toml` (yes pawxi avoids YAML) and pawxi will do the rest. 

Sample config 
```TOML
[proxy]
path = '/'
destination = "http://localhost:6000"
entrypoint = 8080
```


## Stuff i'm working on 


- [x] implement basic request forwarding 
  
- [ ] multi request forwarding 

- [ ] move to environment vars 

- [x] parse TOML config

- [x] compress request  

- [ ] Add TLS 
