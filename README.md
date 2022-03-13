# Climate-GO

### Introduction

Climate-GO is a "CLI-Mate" tool that polls weather data for
your area.

### Development

This command was used to create this boilerplate:
```
cobra init --pkg-name github.com/joshua-hansen/climate-go --author "Joshua Hansen joshuahansen@outlook.com" --license MIT --config ./.cobra.yaml
```

1. Pull the repository
2. `$ go build`
3. `$ ./climate-go`

### Configuration
The default configuration path is `$HOME/.climate.yaml`.

So far this is what that configuration file looks like:
```
{
  debugapp: false,
  units: "imperial",
  climateapiurl: "api.openweathermap.org/data/2.5/weather",
  climateapikey: "<openweathermap-api-key>"
}
```
P.S. I haven't figured out how to write configs using viper yet.
___
### External Resources

Using the [spf13/cobra](https://github.com/spf13/cobra/) tool to create
commands.

Using the [spf13/viper](https://github.com/spf13/viper/) tool to read/write
configuration files.

Using a weather api from [openweathermap](https://openweathermap.org/api).
