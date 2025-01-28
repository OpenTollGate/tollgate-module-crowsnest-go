# Tollgate Module - crowsnest (go)

This Tollgate module is responsible to look if there is LAN in sight, meaning it will look out for upstream tollgates to connect to. It collects datapoints like the other tollgate's price, pubkey, signal strength, etc. This data is passed on as a Nostr event on the local crowsnest ([crowsnest Module](https://github.com/OpenTollGate/tollgate-module-crowsnest-go))

The on-device crowsnest makes no decisions on this data because connecting to another tollgate is a financial decision that is taken by the [Merchant module](https://github.com/OpenTollGate/tollgate-module-merchant-go).

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
