# Tollgate Module - crowsnest (go)
![](https://cdn.satellite.earth/5105c47ab72005a749d38d79fce223fd49a81ad02ce69ac32409830d8acecf5b.png)

This tollgate module is responsible for interaction with the router's radio. It sets the network information and broadcasts things like pricing. It's also responsible for looking out for upstream tollgates to connect to, and collecting datapoints like the other tollgate's price, pubkey, signal strength, etc. 

The on-device crowsnest makes no decisions on this data because connecting to another tollgate is a financial decision, which is a responsibility of the [Merchant module](https://github.com/OpenTollGate/tollgate-module-merchant-go).

# Compile for ATH79 (GL-AR300 NOR)

```bash
cd ./src
env GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o crowsnest -trimpath -ldflags="-s -w"

# Hint: copy to connected router 
scp crowsnest root@192.168.1.1:/root/crowsnest
```

# Compile for GL-MT3000

## Build

```bash
cd ./src
env GOOS=linux GOARCH=arm64 go build -o crowsnest -trimpath -ldflags="-s -w"
scp -O crowsnest root@192.168.1.1:/root/crowsnest # X.X == Router IP
```

packages: 

- iw (at least for debugging)

wifi scanning https://superuser.com/questions/958889/nl80211-not-foundhow-resolve-this-error-in-openwrt
