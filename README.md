# Mini-Lalamove
My Name: Ho Yat Man Peter
Email: peter539@live.hk

This application is implemented in Go with the use of contiainer.

It is my first time to build a command line application. I take it as an opportunity to learn the Go language and give a try to the container development workflow to show that I am a fast learner as well as a competitive candidate for this position.



## Getting Started

### Prerequisites



* Docker is installed in the environment

```sh
docker -v
Docker version 20.10.17, build 100c701
```


### Installation

* Run the `setup.sh` script to build an executable and run in the container environment.

```
./setup.sh
```


<!-- USAGE EXAMPLES -->
## Usage

### Create order
* Create an order by the following command: `llm create_order [from] [to]`

##### Example:

```
/ # llm create_order "MK" "SSP"
1
/ # llm create_order "CWB" "MK"
2
```

### List orders
* List all the available (non-taken) orders with the following command: `llm list_orders`

##### Example:

```
/ # llm list_orders
1,MK,SSP
2,CWB,MK
```

### Take order
* Take an order by ID with the following command: `llm take_order [id]`

##### Example:

```
/ # llm take_order 1
```


