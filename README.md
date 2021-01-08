# Temperature Sensor

## Notes
This is intended to run on a [Raspberry Pi](https://www.raspberrypi.org/) with an attached DHT-22 Temperature/Humidity sensor on GPIO 4.

## Build
```shell 
env GOOS=linux GOARCH=arm go build 
```

## Run on Raspberry Pi
1. Move the compiled binary to the Raspberry Pi
1. On the Raspberry Pi change to the directory where you placed the binary.
1. ```shell
    sudo ./TemperatureSensor
    ```

## Comments
That's it! Try it out if you would like. And shoutout to [MichaelS11](https://github.com/MichaelS11). His code made things so much easier since it is impossible with [periph.io](https://periph.io/) to do microsecond listening on the GPIO port.

