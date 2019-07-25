# Python API Demo Client for Neutral Basket


## Pantsbuild
We use Pantsbuild because it has good support for Python and GRPC
out of the box. Gradlew requires custom PyPi repositories.

You don't need to download anything to use Pantsbuild. Simply run
`./pants test ::` will set up Pantsbuild for you.

## Python
The code is written for Python 2.7. But we use 3.6.x with Pants 1.16.0.
We recommend pyenv for managing multiple Python versions: `https://github.com/pyenv/pyenv`.

## Clean installation Steps
```
install pyenv
pyenv install 3.6.7
cd clients/python
pyenv local 3.6.7
pyenv shell 3.6.7
rm -rf .cache .pants.d $HOME/.cache/pants
./pants run src/python/neutral/apps:demo -- --env test
```


## Organization
The code is split into three main components: the library
(under `src/python/neutral/libs`), the driver (`src/python/neutral/apps/demo`)
and the resources (under `resources/{configs, contracts}`).

A copy of the compiled NeutralValidator contract is placed under
`resources/contracts`.

## Backend environments
Different backend environments are hard-coded in the configs resources.
Three environments are provided (and more will be added as we move toward
production launch).

`local`: Against the local sandbox (on the same host).

`test`: Against the test sandbox with Ganache network.

`ropsten`: Against the ropsten sandbox.

To choose one, simply use the command-line option `--env`.

## How to run
```
./pants binary src/python/neutral/apps:demo --        \
      [--config-file-override CONFIG_FILE_OVERRIDE] \
      [--env ENV] [-P]  \
      [-e EXECUTE_EVERY_N_QUOTES] \
      [--log-level LOGLEVEL] \
      [-k SRC_PRIVATE_KEY] \
      [--dst-address DST_ADDRESS] \
      [-S]

optional arguments:
  -h, --help            show this help message and exit
  --config-file-override CONFIG_FILE_OVERRIDE
                        override the pre-compiled config in the resources
  --env ENV             service environment to run against
  -P, --is-private-account
  -e EXECUTE_EVERY_N_QUOTES, --execute EXECUTE_EVERY_N_QUOTES
                        obtain an execution authorization on every N quotes
  --log-level LOGLEVEL  set the logger level
  -k SRC_PRIVATE_KEY, --src-private-key SRC_PRIVATE_KEY
                        Private key for src address
  --dst-address DST_ADDRESS
                        Ethereum address for the destination of the
                        transaction; default is src address
  -S, --settle-the-commitment
                        settle the commitment to the Neutral smart contract
```

Examples:

Against test sandbox using default key and public account:
```
./pants run src/python/neutral/apps:demo -- -e 1 -S --env test

```
Against test sandbox using default key and private account:
```
./pants run src/python/neutral/apps:demo -- -e 1 -S -P --env test
```

Against ropsten sandbox using a user-supplied key and private account:
```
./pants run src/python/neutral/apps:demo -- -e 1 --env ropsten -k account-key -P
```

** Note that the PEX is not supported because the resource zip files are not supported.
So the pex file produced by `./pants binary src/python/neutral/apps:demo` does not
work. **
