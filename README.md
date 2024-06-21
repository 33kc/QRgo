# QRgo

## Overview

QRgo is a simple Go application that allows users to enter a website URL and generates a corresponding QR code. This QR code can be scanned to quickly access the website.

## features

- Input website URL and generate QR code
- Save the generated QR code as an image file

## Prerequisites

- Go 1.16 or later

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/33kc/QRgo.git -b main qrgo
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:
    ```sh
    qrgo -url "https://website.com"
    ```

2. Enter the website URL when prompted. (with https://)

3. The generated QR code will be saved as an image file in the current directory

## Example

```sh
qrgo -url "https://33kc.github.io"
```
