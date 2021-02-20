## Whats this ? 

Pawxi is yet another reverse proxy desinged with simplicity in mind. Born out of a certain users frustration at the complexity of certain proxies. 

## Cool how does it works ? 

Simply create a `config.toml` (yes pawxi avoids YAML) and pawxi will do the rest. 

Sample config 
```TOML
[ProxyConfig]
path = '/'
destination = 6000
entrypoint = 8080
```

## Stuff i'm working on 


- [ ] implement basic request forwarding 
  
- [ ] multi request forwarding 

- [ ] move to environment vars 

- [ ] parse TOML config

- [ ] compress request  

- [ ] Add TLS 

- [ ] graceful shutdown 

- [ ] service discovery 