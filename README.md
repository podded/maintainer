# Bouncer

![Image of Gopher Worker](assets/maintainer.svg)

Maintainer is designed to interface directly with mongodb and redis to find and correct any errors from the ingest process.
It will attempt to fix any problem it can, and report those that it cannot 

## Installation

Installation is as easy as cloning and running make (assuming you have go installed).

```shell script
git clone https://github.com/podded/maintainer
cd maintainer
make
```


## Usage

Running maintainer without any arguments
```shell script
./bin/maintainer
```
will attempt to run all fixes. Alternatively you can run any fix by name individually.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)