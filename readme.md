# goos-sidecar
    goos 边车，插件模式
    
### 基本设想
    1、基于caddy、coreDNS的插件系统进行进一步开发
    2、dns发现：优先使用coreDNS的缓存系统
    3、关于配置获取，基于dns系统，优先适配spring clould nacos config api
    4、关于服务注册，基于dns系统，优先适配spring clould nacos discovery api
    
### core
    proxy 实现核心sidecar代理
    coreDNS 实现服务注册，使用 DNS Filter
    
    
### goos sidecar plugins
    1、config-spring-nacos:支持nacos的服务发现api
    
    

spring application