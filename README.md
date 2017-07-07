# Big Brother

Delightful remote status collector.

## Usage

### Agent
The agent needs to run on every machine you want to observe.
The agent sends the collected information to any interested observers at predefined intervals.

You can set the interval and network configurations in `agent/config/config.yaml`

### Master
The master gathers the information from all the agents and displays it in a UI or makes it available to query using an API.

Master settings can be found in `master/config/config.yaml`

## License
Big Brother is released under the MIT license. See LICENSE for details.

## Contact
Follow me on twitter [@mcostea](https://twitter.com/mcostea)

