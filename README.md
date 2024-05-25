# Math-Skills

This is a project written in Go that performs statistics calculations such as Average, Median, Variance, and Standard Deviation

## Features

The purpose of this project is to calculate:

- Average

- Median

- Variance

- Standard Deviation

The program reads from a file and prints the result of each statistics mentioned above. 

## Prerequisites

- **Average**

It's the sum of all numbers in a set divided by how many numbers there are.

- **Median**

It's the middle number when the numbers are arranged in order (ascending or descending). If there's an even number of values, it's the average of the two middle numbers.

- **Variance**

It's a measure of how spread out the numbers in a set are from the average. It shows you how much the numbers deviate from the mean.

- **Standard Deviation**

It's the square root of the variance.

## Getting Started

To interact with the program on your local machine, you will need to have Go language installed on your machine. Use the link below to download and install Golang.

https://go.dev/doc/install

## Installation

1. Clone the repository:

   ```
   $ git clone https://github.com/johneliud/Math-Skills.git
   ```

2. Navigate into the go-reloaded directory:
   ```
   $ cd Math-Skills
   ```

## Usage

Run the program with the commands below:

```
$ go run . data.txt
```

### Example

To run the program, the user needs to populate data.txt file with more than one number. The program will then calculate statistics as shown below:

The contents in the data.txt can be as shown below:

```
1.   2
2.   4
3.   6
4.   8
5.   10
6.
```

Run the program:

```
go run . data.txt
```

If successful, the output will be displayed on the terminal as below:

```
Average: 6
Median: 6
Variance: 8
Standard Deviation: 3
```

**IMPORTANT NOTE:** The resulting values are rounded to the nearest number. Example **1.5** will be rounded to **2** while **1.4** will be rounded to **1**.

## Contact

Incase of further enquiries, please reach out via the email address johneliud4@gmail.com