## Whats this ? 

Pawxi is yet another reverse proxy desinged with simplicity in mind. Born out of a certain users frustration at the complexity of  setting up certain proxies. 


## Features :sparkles:

 - GZIP compression :zap:

 - Live Reload  🔃

 - TLS termination ( coming soon! ) :fire:


## Usage 

Using the binary.

Create a pawxi.toml file 

```toml
[proxy]
usegzip = "True"
binds = "8080"
routes = [
    {path="/",destination="http://localhost:6000/"},
    {path="/home",destination="http://localhost:5000"},
    {path="/app",destination="http://localhost:4000"},
]
```


## Demo 

![demo](demo.gif)




