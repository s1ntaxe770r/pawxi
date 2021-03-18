                                  ![paw](https://img.icons8.com/cute-clipart/64/000000/dog-footprint.png)


## Whats this ? 

Pawxi is yet another reverse proxy designed with simplicity in mind. Born out of a certain users frustration at the complexity of  setting up certain proxies. 


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


 usegzip  ->  enable or disable gzip 

 Binds  -> what port should the proxy run on 

routes -> paths you want to proxy to


## Demo 

![demo](demo.gif)








