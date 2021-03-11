## Whats this ? 

Pawxi is yet another reverse proxy desinged with simplicity in mind. Born out of a certain users frustration at the complexity of certain proxies. 

## Cool how does it work ? 

Simply create a `pawxi.toml` (yes pawxi avoids YAML) and pawxi will do the rest. 

Sample config 
```TOML
[proxy]
domain = "local.dev"
usegzip = "true"

routes = [
    {path="/",destination="http://localhost:6000/"},
    {path="/home",destination="http://localhost:5000"},
    {path="/app",destination="http://localhost:4000"},
]

```


## Stuff i'm working on 


- [x] implement basic request forwarding 
  
- [x] multi request forwarding 

- [ ] visualize config 

- [x] parse TOML config

- [x] compress request  

- [ ] Add TLS 

