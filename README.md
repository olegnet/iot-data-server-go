### iot-data-server

  Server companion for [temperature sensors](https://github.com/olegnet/arduino-temp-sensors) project

  Go with Fiber web engine and Postgres database (port of [Rust version](https://github.com/olegnet/iot-data-server))

#### Database

```shell
sudo -u postgres createuser iot-data -S -R -P
```

```postgresql
create table temperature_sensors
(
    sensor_id int not null,
    temperature real not null,
    time timestamp not null,
    constraint temperature_sensors_pk
        primary key (sensor_id, time)
);

grant insert, select on temperature_sensors to "iot-data";

```
