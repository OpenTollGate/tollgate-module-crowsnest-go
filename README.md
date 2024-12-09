# Tollgate Module - relay (go)

This Tollgate module is responsible to look if there is LAN in sight, meaning it will look out for upstream tollgates to connect to. It collects datapoints like the other tollgate's price, pubkey, signal strength, etc. This data is passed on as a Nostr event on the local relay ([Relay Module](https://github.com/OpenTollGate/tollgate-module-relay-go))

The on-device crowsnest makes no decisions on this data because connecting to another tollgate is a financial decision that is taken by the [Merchant module](https://github.com/OpenTollGate/tollgate-module-merchant-go).

# Compile for ATH79 (GL-AR300 NOR)

```bash
cd ./src
env GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o relay -trimpath -ldflags="-s -w"

# Hint: copy to connected router 
scp relay root@119.201.26.1:/tmp/relay
```